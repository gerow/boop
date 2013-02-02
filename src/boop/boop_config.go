package boop

import (
  "io/ioutil"
  "encoding/json"
)

var default_config_file_locations = [...]string { "/usr/etc/boop/config.json", "./config.json" }

type Command struct {
  path string
  command string
  only_allow_ips []string
  limit_rate int
}

type Config struct {
  port int
  only_allow_ips []string
  commands []Command
}

func LoadConfig() (config Config) {
  for _, v := range default_config_file_locations {
    ok := false
    if config, ok = LoadConfigFromFile(v); ok {
      return
    }
  }
  return
}

func LoadConfigFromFile(location string) (config Config, ok bool) {
  file, e := ioutil.ReadFile(location)
  if e != nil {
    return config, false
  }

  json.Unmarshal(file, &config)

  return config, true
}
