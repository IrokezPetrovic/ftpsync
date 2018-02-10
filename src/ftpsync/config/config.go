package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	conf Config
)

func Load(cfgpath string) {
	bytes, err := ioutil.ReadFile(cfgpath)
	if err != nil {
		panic("Can`t load config")
	}

	err = json.Unmarshal(bytes, &conf)
	if err != nil {
		fmt.Println(err)
		panic("Error parse config")
	}
}

func Get() *Config {
	return &conf
}
