package boop

import (
  "testing"
  "os"
)

func TestLoadConfigFromFile(t *testing.T) {
  const filename = "test.config.json"
  const port = 1980
  const only_allow_ips_len = 2
  const only_allow_ips_root_0 = "192.168.1.1"
  const only_allow_ips_root_1 = "192.169.1.1"

  const commands_len = 2

  const commands_0_path = "/here/is/path"
  const commands_0_command = "touch hello"
  const commands_0_only_allow_ips_len = 2
  const commands_0_only_allow_ips_0 = "192.188.1.5"
  const commands_0_only_allow_ips_1 = "192.162.1.55"
  const commands_0_limit_rate = 120

  const commands_1_path = "/here/is/another/path"
  const commands_1_command = "touch hello"
  const commands_1_only_allow_ips_len = 2
  const commands_1_only_allow_ips_0 = "192.188.1.5"
  const commands_1_only_allow_ips_1 = "192.162.1.55"
  const commands_1_limit_rate = 120

  cwd,_ := os.Getwd()
  t.Logf("cwd is %s", cwd)

  config,ok := LoadConfigFromFile(filename)

  if !ok {
    t.Errorf("LoadConfigFromFile(%s) returned ok = %t", filename, ok)
  }

  if config.port != port {
    t.Errorf("config.port = %d, want %d", config.port, port)
  }

  if len(config.only_allow_ips) != only_allow_ips_len {
    t.Errorf("len(config.only_allow_ips) = %d, want %d", len(config.only_allow_ips), only_allow_ips_len)
  }

  if config.only_allow_ips[0] != only_allow_ips_root_0 {
    t.Errorf("config.only_allow_ips[0] = %d, want %d", config.only_allow_ips[0], only_allow_ips_root_0)
  }

  if config.only_allow_ips[1] != only_allow_ips_root_1 {
    t.Errorf("config.only_allow_ips[1] = %d, want %d", config.only_allow_ips[1], only_allow_ips_root_0)
  }

  if len(config.commands) != commands_len {
    t.Errorf("len(config.commands) = %d, want %d", len(config.commands), commands_len)
  }

  if config.commands[0].path != commands_0_path {
    t.Errorf("config.commands[0].path = %s, want %s", config.commands[0].path, commands_0_path)
  }

  if config.commands[0].command != commands_0_command {
    t.Errorf("config.commands[0].command = %s, want %s", config.commands[0].command, commands_0_command)
  }

  if len(config.commands[0].only_allow_ips) != commands_0_only_allow_ips_len {
    t.Errorf("len(config.commands[0].only_allow_ips) = %d, want %d", len(config.commands[0].only_allow_ips), commands_0_only_allow_ips_len)
  }

  if config.commands[0].only_allow_ips[0] != commands_0_only_allow_ips_0 {
    t.Errorf("commands[0].only_allow_ips[0] = %s, want %s", config.commands[0].only_allow_ips[0], commands_0_only_allow_ips_0)
  }

  if config.commands[0].only_allow_ips[1] != commands_0_only_allow_ips_1 {
    t.Errorf("commands[0].only_allow_ips[1] = %s, want %s", config.commands[0].only_allow_ips[1], commands_0_only_allow_ips_1)
  }

  if config.commands[0].limit_rate != commands_0_limit_rate {
    t.Errorf("commads[0].limit_rate = %d, want %d", config.commands[0].limit_rate, commands_0_limit_rate)
  }

  if config.commands[0].path != commands_0_path {
    t.Errorf("config.commands[0].path = %s, want %s", config.commands[0].path, commands_0_path)
  }


  if config.commands[1].command != commands_1_command {
    t.Errorf("config.commands[1].command = %s, want %s", config.commands[1].command, commands_1_command)
  }

  if len(config.commands[1].only_allow_ips) != commands_1_only_allow_ips_len {
    t.Errorf("len(config.commands[1].only_allow_ips) = %d, want %d", len(config.commands[1].only_allow_ips), commands_1_only_allow_ips_len)
  }

  if config.commands[1].only_allow_ips[0] != commands_1_only_allow_ips_0 {
    t.Errorf("commands[1].only_allow_ips[0] = %s, want %s", config.commands[1].only_allow_ips[0], commands_1_only_allow_ips_0)
  }

  if config.commands[1].only_allow_ips[1] != commands_1_only_allow_ips_1 {
    t.Errorf("commands[1].only_allow_ips[1] = %s, want %s", config.commands[1].only_allow_ips[1], commands_1_only_allow_ips_1)
  }

  if config.commands[1].limit_rate != commands_1_limit_rate {
    t.Errorf("commads[1].limit_rate = %d, want %d", config.commands[1].limit_rate, commands_1_limit_rate)
  }
}
