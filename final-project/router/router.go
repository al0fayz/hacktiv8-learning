package router

import (
	"hacktiv8-learning/final-project/handler"
	"hacktiv8-learning/final-project/midlleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Final Project Hacktiv-8",
		})
	})
	//api
	api := router.Group("api")
	{
		api.POST("/users/register", handler.RegisterHadnler)
		api.POST("/users/login", handler.Login)
		api.Use(midlleware.InitUserMiddleware())
		api.PUT("/users", handler.UpdateUser)
		api.DELETE("/users", handler.DeleteUser)
	}
	return router
}
