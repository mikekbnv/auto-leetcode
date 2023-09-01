package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sync"

	"gopkg.in/yaml.v2"
)

var (
	lock           = &sync.Mutex{}
	LeetcodeConfig *Config
)

type Config struct {
	CSRFToken string `yaml:"csrf_token"`
	JWTToken  string `yaml:"jwt_token"`
	ContestID string `yaml:"contest_id"`
	Delay     int    `yaml:"delay"`
}

func loadConfigFromFile() (*Config, error) {
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return nil, err
	}

	var config Config

	if err := yaml.Unmarshal(yamlFile, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func CreateLeetcodeConfig() *Config {

	if LeetcodeConfig == nil {
		lock.Lock()
		defer lock.Unlock()
		if LeetcodeConfig == nil {
			config, err := loadConfigFromFile()
			if err != nil {
				log.Fatalf("Error loading config from file: %v", err)
			}
			LeetcodeConfig = config
		} else {
			fmt.Println("LeetcodeConfig already initialized")
		}
	} else {
		fmt.Println("LeetcodeConfig already initialized")
	}

	return LeetcodeConfig
}
