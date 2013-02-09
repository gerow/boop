package boop

import (
  "io/ioutil"
  "encoding/json"
  "fmt"
)

var default_config_file_locations = [...]string { "/usr/etc/boop/config.json", "./config.json" }

type Command struct {
  Path string `json:"path"`
  Command string `json:"command"`
  OnlyAllowIps []string `json:"only_allow_ips"`
  LimitRate int `json:"limit_rate"`
}

type Config struct {
  Port int `json:"port"`
  OnlyAllowIps []string `json:"only_allow_ips"`
  Commands []Command `json:"commands"`
}

func LoadConfig() (config *Config) {
  for _, v := range default_config_file_locations {
    ok := false
    if config, ok = LoadConfigFromFile(v); ok {
      return
    }
  }
  return
}

func LoadConfigFromFile(location string) (config *Config, ok bool) {
  var conf Config

  file, e := ioutil.ReadFile(location)
  if e != nil {
    return config, false
  }

  err := json.Unmarshal(file, &conf)
  if err != nil {
    fmt.Println("error:",err)
  }

  return &conf, true
}
