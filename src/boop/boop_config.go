package boop

import (
  "io/ioutil"
  "encoding/json"
)

var default_config_file_locations = [...]string { "./config.json", "/usr/etc/boop/config.json" }
const default_port = 9180
const post_action = "POST"

type Command struct {
  Path string `json:"path"`
  Command string `json:"command"`
  OnlyAllowIps []string `json:"only_allow_ips"`
  LimitRate int `json:"limit_rate"`
  AllowActions []string `json:"allow_actions"`
}

type Config struct {
  Port int `json:"port"`
  OnlyAllowIps []string `json:"only_allow_ips"`
  Commands []Command `json:"commands"`
}

func LoadConfig() (*Config, error) {
  var err error
  var config *Config

  for _, v := range default_config_file_locations {
    if config, err = LoadConfigFromFile(v); err == nil {
      return config, nil
    }
  }
  return nil, err
}

func LoadConfigFromFile(location string) (*Config, error) {
  file_data, err := ioutil.ReadFile(location)
  if err != nil {
    return nil, err
  }

  return LoadConfigFromBytes(file_data)
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

func fillInDefaultValues(config *Config) (*Config) {
  if (config.Port == 0) {
    config.Port = default_port;
  }

  // Make sure we allow only POST by
  // default
  for _, v := range config.Commands {
    if len(v.AllowActions) == 0 {
      v.AllowActions = append(v.AllowActions, post_action)
    }
  }

  return config
}


