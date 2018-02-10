package main

import (
	"fmt"
	"ftpsync/config"
	"os"
	"sync"
)

func main() {
	var cfgpath = "config.json"
	if len(os.Args) == 2 {
		cfgpath = os.Args[1]
	}

	fmt.Printf("Config from %s \n", cfgpath)
	config.Load(cfgpath)
	conf := config.Get()
	var wg sync.WaitGroup

	for _, profile := range conf.Profiles {
		wg.Add(1)
		go Sync(profile, &wg)
	}

	wg.Wait()
}
