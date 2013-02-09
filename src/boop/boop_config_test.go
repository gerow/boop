package boop

import (
  "testing"
  "os"
  "reflect"
)

func TestLoadConfigFromFileWithDefaults(t *testing.T) {
  const filename = "test.empty.config.json"

  var expected_config Config

  expected_config.Port = 9180

  config, err := LoadConfigFromFile(filename)

  if err != nil {
    t.Errorf("Got error from LoadConfigFromFile:", err)
  }

  if !reflect.DeepEqual(expected_config, *config) {
    t.Errorf("Loaded config does not match expected. Expected %v, got %v", expected_config, *config)
  }

}

func TestLoadConfigFromFile(t *testing.T) {
  const filename = "test.config.json"
  const Port = 9180
  const OnlyAllowIps_len = 2
  const OnlyAllowIps_root_0 = "192.168.1.1"
  const OnlyAllowIps_root_1 = "192.169.1.1"

  const Commands_len = 2

  const Commands_0_Path = "/here/is/path"
  const Commands_0_Command = "touch hello"
  const Commands_0_OnlyAllowIps_len = 2
  const Commands_0_OnlyAllowIps_0 = "192.188.1.5"
  const Commands_0_OnlyAllowIps_1 = "192.162.1.55"
  const Commands_0_LimitRate = 120

  const Commands_1_Path = "/here/is/another/path"
  const Commands_1_Command = "touch hello"
  const Commands_1_OnlyAllowIps_len = 2
  const Commands_1_OnlyAllowIps_0 = "192.188.1.5"
  const Commands_1_OnlyAllowIps_1 = "192.162.1.55"
  const Commands_1_LimitRate = 120

  cwd,_ := os.Getwd()
  t.Logf("cwd is %s", cwd)

  config,err := LoadConfigFromFile(filename)

  if err != nil {
    t.Errorf("Got error from LoadConfigFromFile:", err)
  }

  t.Logf("config looks like %v", config)

  if config.Port != Port {
    t.Errorf("config.Port = %d, want %d", config.Port, Port)
  }

  if len(config.OnlyAllowIps) != OnlyAllowIps_len {
    t.Errorf("len(config.OnlyAllowIps) = %d, want %d", len(config.OnlyAllowIps), OnlyAllowIps_len)
  }

  if config.OnlyAllowIps[0] != OnlyAllowIps_root_0 {
    t.Errorf("config.OnlyAllowIps[0] = %d, want %d", config.OnlyAllowIps[0], OnlyAllowIps_root_0)
  }

  if config.OnlyAllowIps[1] != OnlyAllowIps_root_1 {
    t.Errorf("config.OnlyAllowIps[1] = %d, want %d", config.OnlyAllowIps[1], OnlyAllowIps_root_0)
  }

  if len(config.Commands) != Commands_len {
    t.Errorf("len(config.Commands) = %d, want %d", len(config.Commands), Commands_len)
  }

  if config.Commands[0].Path != Commands_0_Path {
    t.Errorf("config.Commands[0].Path = %s, want %s", config.Commands[0].Path, Commands_0_Path)
  }

  if config.Commands[0].Command != Commands_0_Command {
    t.Errorf("config.Commands[0].Command = %s, want %s", config.Commands[0].Command, Commands_0_Command)
  }

  if len(config.Commands[0].OnlyAllowIps) != Commands_0_OnlyAllowIps_len {
    t.Errorf("len(config.Commands[0].OnlyAllowIps) = %d, want %d", len(config.Commands[0].OnlyAllowIps), Commands_0_OnlyAllowIps_len)
  }

  if config.Commands[0].OnlyAllowIps[0] != Commands_0_OnlyAllowIps_0 {
    t.Errorf("Commands[0].OnlyAllowIps[0] = %s, want %s", config.Commands[0].OnlyAllowIps[0], Commands_0_OnlyAllowIps_0)
  }

  if config.Commands[0].OnlyAllowIps[1] != Commands_0_OnlyAllowIps_1 {
    t.Errorf("Commands[0].OnlyAllowIps[1] = %s, want %s", config.Commands[0].OnlyAllowIps[1], Commands_0_OnlyAllowIps_1)
  }

  if config.Commands[0].LimitRate != Commands_0_LimitRate {
    t.Errorf("commads[0].LimitRate = %d, want %d", config.Commands[0].LimitRate, Commands_0_LimitRate)
  }

  if config.Commands[0].Path != Commands_0_Path {
    t.Errorf("config.Commands[0].Path = %s, want %s", config.Commands[0].Path, Commands_0_Path)
  }


  if config.Commands[1].Command != Commands_1_Command {
    t.Errorf("config.Commands[1].Command = %s, want %s", config.Commands[1].Command, Commands_1_Command)
  }

  if len(config.Commands[1].OnlyAllowIps) != Commands_1_OnlyAllowIps_len {
    t.Errorf("len(config.Commands[1].OnlyAllowIps) = %d, want %d", len(config.Commands[1].OnlyAllowIps), Commands_1_OnlyAllowIps_len)
  }

  if config.Commands[1].OnlyAllowIps[0] != Commands_1_OnlyAllowIps_0 {
    t.Errorf("Commands[1].OnlyAllowIps[0] = %s, want %s", config.Commands[1].OnlyAllowIps[0], Commands_1_OnlyAllowIps_0)
  }

  if config.Commands[1].OnlyAllowIps[1] != Commands_1_OnlyAllowIps_1 {
    t.Errorf("Commands[1].OnlyAllowIps[1] = %s, want %s", config.Commands[1].OnlyAllowIps[1], Commands_1_OnlyAllowIps_1)
  }

  if config.Commands[1].LimitRate != Commands_1_LimitRate {
    t.Errorf("commads[1].LimitRate = %d, want %d", config.Commands[1].LimitRate, Commands_1_LimitRate)
  }
}
