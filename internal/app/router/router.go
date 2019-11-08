package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/neverhaveiever-io/api/internal/app/middleware/error"
	v1 "github.com/neverhaveiever-io/api/internal/app/router/v1"
)

func Init(auth gin.HandlerFunc) {

	router := gin.Default()

	router.Use(cors.Default())
	router.Use(error.Error())

	router.GET("/v1/statements/:id", v1.GetStatement)

	authorized := router.Group("/v1", auth)
	{
		authorized.POST("/statements", v1.AddStatement)
		authorized.PUT("/statements/:id", v1.EditStatement)
		authorized.DELETE("/statements/:id", v1.DeleteStatement)
	}

	if err := router.Run(); err != nil {
		panic("unable initialize router")
	}
}
