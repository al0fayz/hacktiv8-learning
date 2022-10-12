package handler

import (
	"errors"
	"hacktiv8-learning/final-project/midlleware"
	"hacktiv8-learning/final-project/models"
	"hacktiv8-learning/final-project/rules"
	"hacktiv8-learning/final-project/utils"
	"hacktiv8-learning/final-project/validators"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func RegisterHadnler(c *gin.Context) {
	req := rules.UserCreateValidator()
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
		user := models.User{
			Username: req.Username,
			Email:    req.Email,
			Password: req.Password,
			Age:      req.Age,
		}
		//create user
		err := models.CreateUser(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":     http.StatusBadRequest,
				"errors":   err.Error(),
				"messages": "Bad Request",
			})
			return
		}
		//success
		c.JSON(http.StatusCreated, gin.H{
			"id":       user.Id,
			"username": user.Username,
			"email":    user.Email,
			"age":      user.Age,
		})
	}
}
func Login(c *gin.Context) {
	req := rules.UserLogin()
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
		//check user is exist
		user, err := models.FindUserByEmail(req.Email)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":     http.StatusBadRequest,
				"errors":   err.Error(),
				"messages": "Bad Request",
			})
			return
		}
		//check password is correct
		err = utils.CompareHash(user.Password, req.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":     http.StatusBadRequest,
				"messages": "Email or Password incorect",
			})
			return
		}
		//generate token
		token, err := midlleware.GenerateToken(uint(user.Id))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":     http.StatusInternalServerError,
				"errors":   err.Error(),
				"messages": "Internal Server Error",
			})
			return
		}
		//success
		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	}
}
