package boop

import (
	"encoding/json"
	"io/ioutil"
        "sync"
        "time"
)

var Locations = [...]string{"./config.json", "/etc/boop/config.json"}

const DefaultPort = 9180

type Command struct {
        sync.Mutex            `json:"_"`
	Path         string   `json:"path"`
	Command      string   `json:"command"`
	OnlyAllowIps []string `json:"onlyAllowIps"`
	LimitRate    int64      `json:"limitRate"`
        LastTimeRun  time.Time      `json:"_"`
}

type Config struct {
	Port         int       `json:"port"`
	OnlyAllowIps []string  `json:"onlyAllowIps"`
	Commands     []Command `json:"commands"`
}

func LoadConfig() (*Config, error) {
	var err error
	var config *Config

	for _, v := range Locations {
		if config, err = LoadConfigFromFile(v); err == nil {
			return config, nil
		}
	}
	return nil, err
}

func LoadConfigFromFile(location string) (*Config, error) {
	fileData, err := ioutil.ReadFile(location)
	if err != nil {
		return nil, err
	}

	return LoadConfigFromBytes(fileData)
}

func LoadConfigFromBytes(data []byte) (*Config, error) {
	var config Config
	err := json.Unmarshal(data, &config)
	if err != nil {
		return &config, err
	}

	fillInDefaultValues(&config)
	return &config, nil
}

func fillInDefaultValues(config *Config) *Config {
	if config.Port == 0 {
		config.Port = DefaultPort
	}
	return config
}
