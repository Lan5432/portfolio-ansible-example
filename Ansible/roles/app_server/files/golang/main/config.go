package main

import (
	"fmt"
	"os"

	"github.com/tkanos/gonfig"
)

const CONFIG_FILE = "config/config.json"
const DEFAULT_CONFIG_FILE = "config/defaultConfig.json"

type Config struct {
	Port int
	TemplateFolder string
}

func GetConfig() Config {
	config := Config{}
	var configErr error
	_, fileErr := os.Stat(CONFIG_FILE)
	
	if fileErr == nil {
		configErr = gonfig.GetConf(CONFIG_FILE, &config)  
	} else {
		configErr = gonfig.GetConf(DEFAULT_CONFIG_FILE, &config)  
	}

	if configErr != nil {
		fmt.Println(configErr)
	}
	
	return config
}