package auth

import (
	"github.com/gin-gonic/gin"
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
		expected gin.Accounts
	}{
		{
			[]string{"test"},
			gin.Accounts{
				"foo": "bar",
			},
		},
		{
			[]string{"baz"},
			gin.Accounts{},
		},
	}

	for k, c := range cases {
		result := Accounts(c.input)
		if !reflect.DeepEqual(result, c.expected) {
			t.Fatalf("%d: Unexpected output: %+v", k, result)
		}
	}
}
