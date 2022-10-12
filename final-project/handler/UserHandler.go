package handler

import (
	"errors"
	"fmt"
	"hacktiv8-learning/final-project/midlleware"
	"hacktiv8-learning/final-project/models"
	"hacktiv8-learning/final-project/rules"
	"hacktiv8-learning/final-project/validators"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func UpdateUser(c *gin.Context) {
	req := rules.UserUpdateValidator()
	if err := req.Bind(c); err != nil {
		var errValidation validator.ValidationErrors
		if errors.As(err, &errValidation) {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code":     http.StatusUnprocessableEntity,
				"errors":   validators.NewValidatorError(errValidation),
				"messages": "invalid input",
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":     http.StatusInternalServerError,
				"errors":   err.Error(),
				"messages": "Internal Server Error",
			})
			return
		}
	} else {
		user, err := models.FindUserByEmail(req.Email)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":     http.StatusBadRequest,
				"errors":   err.Error(),
				"messages": "Bad Request",
			})
			return
		}
		//update user
		user.Username = req.Username
		err = models.UpdateUser(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":     http.StatusBadRequest,
				"errors":   err.Error(),
				"messages": "Bad Request",
			})
			return
		}
		//success
		c.JSON(http.StatusOK, gin.H{
			"id":         user.Id,
			"username":   user.Username,
			"email":      user.Email,
			"age":        user.Age,
			"updated_at": user.UpdatedAt,
		})
	}
}

func DeleteUser(c *gin.Context) {
	//get id from jwt
	user_id, err := midlleware.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":     http.StatusInternalServerError,
			"errors":   err.Error(),
			"messages": "Internal Server Error",
		})
		return
	}
	//delete on db
	err = models.DeleteUser(fmt.Sprint(user_id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":     http.StatusInternalServerError,
			"errors":   err.Error(),
			"messages": "Internal Server Error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"messages": "Your account has been successfuly deleted",
	})
}
