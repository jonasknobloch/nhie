package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nhie-io/api/internal/app/auth"
	"github.com/nhie-io/api/internal/app/middleware/prometheus"
	v1 "github.com/nhie-io/api/internal/app/router/v1"
)

func Init() {

	router := gin.Default()

	router.Use(cors.Default())

	// initialize prometheus metrics
	prometheus.UseWithAuth(router, auth.Accounts([]string{"admin", "metrics"}))

	router.GET("/v1/statements/:id", v1.GetStatement)

	authorized := router.Group("/v1", gin.BasicAuth(auth.Accounts([]string{"admin"})))
	{
		authorized.POST("/statements", v1.AddStatement)
		authorized.PUT("/statements/:id", v1.EditStatement)
		authorized.DELETE("/statements/:id", v1.DeleteStatement)
	}

	if err := router.Run(); err != nil {
		panic("unable initialize router")
	}
}
