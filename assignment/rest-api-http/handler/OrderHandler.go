package handler

import (
	"hacktiv8-learning/assignment/rest-api-http/model"
	"hacktiv8-learning/assignment/rest-api-http/request"
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
	var order request.OrderCreate
	err := c.ShouldBindJSON(&order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	err = model.CreateOrder(&order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    nil,
		"message": "Create data success",
	})
}
func DetailOrder(c *gin.Context) {
	id := c.Params.ByName("id")
	order, err := model.DetailOrder(id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"data":     order,
		"messages": "Get Detail Data Success",
	})
}
func UpdateOrder(c *gin.Context) {

}
func DeleteOrder(c *gin.Context) {
	id := c.Params.ByName("id")
	err := model.DeleteOrder(id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"data":     nil,
		"messages": "Delete Data Success",
	})
}
