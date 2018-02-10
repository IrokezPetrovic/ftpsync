package main

import (
	"fmt"
	"ftpsync/config"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/secsy/goftp"
)

func Sync(task config.Task, wg *sync.WaitGroup) {
	defer wg.Done()
	var ftpConf goftp.Config
	ftpConf.User = task.Destination.Username
	ftpConf.Password = task.Destination.Password
	ftpClient, err := goftp.DialConfig(ftpConf, "monitor.digisky.ru:2124")

	if err != nil {
		return
	}
	syncDir(task.Source, string(filepath.Separator), ftpClient, task.Destination.Path)
}

func syncDir(base, path string, ftp *goftp.Client, ftpbase string) {
	fmt.Printf("Sync %s \n", path)
	files, err := ioutil.ReadDir(filepath.Join(base, path))
	if err != nil {
		return
	}

	ftpFilesMap := make(map[string]os.FileInfo)
	locFilesMap := make(map[string]os.FileInfo)

	ftpfiles, err := ftp.ReadDir(ftpath(ftpbase + "/" + path))

	for _, ftpfile := range ftpfiles {
		ftpFilesMap[ftpfile.Name()] = ftpfile
	}

	for _, file := range files {
		ftpFile := ftpFilesMap[file.Name()]
		if file.IsDir() {
			if ftpFile != nil && !ftpFile.IsDir() {
				ftp.Delete(ftpath(ftpbase + "/" + path + "/" + file.Name()))
				ftpFile = nil
			}

			if ftpFile == nil {
				fmt.Printf("Make dir %s \n", ftpath(ftpbase+"/"+path+"/"+file.Name()))
				ftp.Mkdir(ftpath(ftpbase + "/" + path + "/" + file.Name()))
				syncDir(base, filepath.Join(path, file.Name()), ftp, ftpbase)
			}

		} else {

			if ftpFile != nil && ftpFile.IsDir() {
				//ftp.Rmdir(ftpath(ftpbase + "/" + path + "/" + file.Name()))
				rmdir(ftpath(ftpbase+"/"+path+"/"+file.Name()), ftp)
				ftpFile = nil
			}

			if ftpFile == nil ||
				ftpFile.ModTime().Unix() < file.ModTime().Unix() {
				uploadToFtp(base, filepath.Join(path, file.Name()), ftp, ftpbase)
			}
		}
		locFilesMap[file.Name()] = file
	}

	for _, ftpFile := range ftpfiles {
		locFile := locFilesMap[ftpFile.Name()]
		if locFile == nil {
			if ftpFile.IsDir() {
				fmt.Printf("Remove directory %s \n", ftpath(ftpbase+"/"+path+"/"+ftpFile.Name()))
				rmdir(ftpath(ftpbase+"/"+path+"/"+ftpFile.Name()), ftp)

			} else {
				fmt.Printf("Remove file %s \n", ftpath(ftpbase+"/"+path+"/"+ftpFile.Name()))
				ftp.Delete(ftpath(ftpbase + "/" + path + "/" + ftpFile.Name()))
			}
		}
	}
}

func uploadToFtp(base, src string, ftp *goftp.Client, ftpbase string) {
	fmt.Printf("Upload %s to %s \n", src, ftpath(ftpbase+"/"+src))
	file, _ := os.Open(filepath.Join(base, src))
	ftp.Store(ftpath(ftpbase+"/"+src), file)
	file.Close()

}

var sep string = string(filepath.Separator)

func ftpath(p string) string {
	return strings.Replace(p, sep, "/", -1)
}

func rmdir(path string, ftp *goftp.Client) {
	files, _ := ftp.ReadDir(path)
	for _, file := range files {
		if file.IsDir() {
			rmdir(path+"/"+file.Name(), ftp)
		} else {
			ftp.Delete(path + "/" + file.Name())
		}
	}
	ftp.Rmdir(path)
}
