package boop

import (
  "testing"
)

func TestLoadConfigFromFileWithDefaults(t *testing.T) {
  const filename = "test.empty.config.json"

  var expected_config Config

  expected_config.Port = 9180

  config, err := LoadConfigFromFile(filename)

  if err != nil {
    t.Errorf("Got error from LoadConfigFromFile:", err)
  }

  if !reflect.DeepEqual(expected_config, *config, t) {
    t.Errorf("Loaded config does not match expected. Expected %v, got %v", expected_config, *config)
  }

}

func TestLoadConfigFromFile(t *testing.T) {
  const filename = "test.config.json"

  var exp Config

  exp.Port = 9180
  exp.OnlyAllowIps = []string{ "192.168.1.1", "192.169.1.1" }
  exp.Commands = []Command{
      Command{
        "GET /here/is/get/path",
        "touch hello",
        []string{ "192.188.1.5", "192.162.1.55" },
        120 },
      Command{
        "POST /here/is/post/path",
        "touch hello2",
        []string{ "192.188.2.5", "192.162.2.55" },
        320 },
      Command{
        "DELETE /here/is/delete/path",
        "sleep 200",
        []string{ },
        0 } }

  config,err := LoadConfigFromFile(filename)

  if err != nil {
    t.Errorf("Got error from LoadConfigFromFile:", err)
  }

  if !reflect.DeepEqual(exp, *config, t) {
    t.Errorf("Loaded config does not match expected. Expected %v, got %v", exp, *config)
  }
}
