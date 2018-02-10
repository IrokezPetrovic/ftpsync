package main

import (
	"testing"

	"github.com/secsy/goftp"
)

func TestNotExistFolder(t *testing.T) {
	var ftpConfig goftp.Config
	ftpConfig.User = "Administrator@digisky.lan"
	ftpConfig.Password = "Copoakbo123"

	ftp, err := goftp.DialConfig(ftpConfig, "monitor.digisky.ru:2124")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	s, err := ftp.Stat("/Pool0/production/NOTEXISITDIR")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	if s == nil {
		t.Log("s=nil")
		t.FailNow()
	}
}
