package error

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// use _ = c.Error(err) to pass error to error handler
// use c.Error(err).Err.Error() to pass error to handler and keep original error

func Error() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		for _, err := range c.Errors {
			fmt.Println(err.Error())
		}
	}
}
