package auth

import (
	"github.com/spf13/viper"
	"os"
	"reflect"
	"testing"
)

func TestAccounts(t *testing.T) {
	_ = os.Setenv("NHIE_TEST_USER", "foo")
	_ = os.Setenv("NHIE_TEST_PASS", "bar")

	viper.SetEnvPrefix("NHIE")
	viper.AutomaticEnv()

	cases := []struct {
		input    []string
		expected map[string]string
	}{
		{
			[]string{"test"},
			map[string]string{
				"foo": "bar",
			},
		},
		{
			[]string{"baz"},
			map[string]string{},
		},
	}

	for k, c := range cases {
		result := Accounts(c.input)
		if !reflect.DeepEqual(result, c.expected) {
			t.Fatalf("%d: Unexpected output: %+v", k, result)
		}
	}
}
