package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/neverhaveiever-io/api/middleware/error"
	"github.com/neverhaveiever-io/api/routers/v1"
)

func InitRouter(auth gin.HandlerFunc) *gin.Engine {

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

	return router
}
