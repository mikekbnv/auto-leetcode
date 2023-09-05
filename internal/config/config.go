package config

import (
	"fmt"
	"log"
	"os"
	"strings"
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

func InitLeetcodeConfig() (error) {
	if LeetcodeConfig == nil {
		lock.Lock()
		defer lock.Unlock()
		if LeetcodeConfig == nil {
			//fmt.Println("LeetcodeConfig is nil, loading config from file")
			config, err := loadConfigFromFile()
			if err != nil {
				log.Fatalf("Error loading config from file: %v", err)
				return err
			}
			LeetcodeConfig = config
			//fmt.Println("Loaded config from file successfully")
		} else {
			return fmt.Errorf("leetcodeConfig already initialized")
		}
	} else {
		return fmt.Errorf("leetcodeConfig already initialized")
	}

	return nil
}

func loadConfigFromFile() (*Config, error) {
	yamlFile, err := os.ReadFile("config.yaml")
	if err != nil {
		return nil, err
	}

	var config Config

	if err := yaml.Unmarshal(yamlFile, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func SetConfigField(field string, value interface{}) {
	lock.Lock()
	defer lock.Unlock()
	switch field {
	case "csrf_token":
		LeetcodeConfig.CSRFToken = value.(string)
	case "jwt_token":
		LeetcodeConfig.JWTToken = value.(string)
	case "contest_id":
		LeetcodeConfig.ContestID = value.(string)
	case "delay":
		LeetcodeConfig.Delay = value.(int)
	}
}

func ParseCookies(cookieStr string) map[string]string {
	cookies := make(map[string]string)

	cookiePairs := strings.Split(cookieStr, ";")

	for _, cookiePair := range cookiePairs {
		parts := strings.SplitN(strings.TrimSpace(cookiePair), "=", 2)
		if len(parts) == 2 {
			key := parts[0]
			value := parts[1]
			cookies[key] = value
		}
	}

	return cookies
}
	



