package rules

import (
	"hacktiv8-learning/final-project/validators"

	"github.com/gin-gonic/gin"
)

type UserCreate struct {
	Username string `form:"username" json:"username" binding:"required,usernameExist"`
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required,min=6"`
	Age      int32  `form:"age" json:"age" binding:"required,gte=8"`
}

func (user *UserCreate) Bind(c *gin.Context) error {
	err := validators.Bind(c, user)
	return err
}

// create
func UserCreateValidator() UserCreate {
	ret := UserCreate{}
	return ret
}
