package handler

import (
	"hacktiv8-learning/assignment/rest-api-http/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListOrder(c *gin.Context) {
	orders, err := model.GetAllData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    orders,
		"message": "Get data success",
	})
}

func CreateOrder(c *gin.Context) {

}
func DetailOrder(c *gin.Context) {

}
func UpdateOrder(c *gin.Context) {

}
func DeleteOrder(c *gin.Context) {

}
