package router

import (
	"hacktiv8-learning/assignment/rest-api-http/handler"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("api")
	{
		api.GET("/orders", handler.ListOrder)
		api.POST("/orders", handler.CreateOrder)
		api.GET("/orders/:id", handler.DetailOrder)
		api.PUT("/orders/:id", handler.UpdateOrder)
		api.DELETE("/orders/:id", handler.DeleteOrder)
	}
	return router

}
