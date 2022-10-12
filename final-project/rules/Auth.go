package rules

import (
	"hacktiv8-learning/final-project/validators"

	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (login *LoginForm) Bind(c *gin.Context) error {
	err := validators.Bind(c, login)
	return err
}

// create
func UserLogin() LoginForm {
	ret := LoginForm{}
	return ret
}
