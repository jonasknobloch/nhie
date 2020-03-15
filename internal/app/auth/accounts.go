package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Accounts(users []string) gin.Accounts {

	accounts := make(gin.Accounts)

	for _, u := range users {
		user := viper.GetString(u + "_user")
		pass := viper.GetString(u + "_pass")

		if user != "" && pass != "" {
			accounts[user] = pass
		}
	}

	return accounts
}
