package app

import (
	"github.com/spf13/viper"
)

func accounts(users []string) map[string]string {

	accounts := make(map[string]string)

	for _, u := range users {
		user := viper.GetString(u + "_user")
		pass := viper.GetString(u + "_pass")

		if user != "" && pass != "" {
			accounts[user] = pass
		}
	}

	return accounts
}
