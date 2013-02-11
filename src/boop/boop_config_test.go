package boop

import (
	"github.com/gerow/gotest"
	"testing"
)

func TestLoadConfigFromFileWithDefaults(t *testing.T) {
	t.Logf("Starting TestLoadConfigFromFileWithDefaults")
	const filename = "test.empty.config.json"

	var expected_config Config

	expected_config.Port = 9180

	config, err := LoadConfigFromFile(filename)

	if err != nil {
		t.Errorf("Got error from LoadConfigFromFile:", err)
	}

	gotest.AssertDeepEqual(expected_config, *config, t)

}

func TestLoadConfigFromFile(t *testing.T) {
	t.Logf("Starting TestLoadConfigFromFileWithDefaults")
	const filename = "test.config.json"

	var exp Config

	exp.Port = 9180
	exp.OnlyAllowIps = []string{"192.168.1.1", "192.169.1.1"}
	exp.Commands = []Command{
		Command{
			"GET /here/is/get/path",
			"touch hello",
			[]string{"192.188.1.5", "192.162.1.55"},
			120},
		Command{
			"POST /here/is/post/path",
			"touch hello2",
			[]string{"192.188.2.5", "192.162.2.55"},
			320},
		Command{
			"DELETE /here/is/delete/path",
			"sleep 200",
			nil,
			0}}

	config, err := LoadConfigFromFile(filename)

	if err != nil {
		t.Errorf("Got error from LoadConfigFromFile:", err)
	}

	gotest.AssertDeepEqual(exp, *config, t)

}
