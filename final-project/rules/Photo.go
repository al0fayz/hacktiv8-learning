package rules

import (
	"hacktiv8-learning/final-project/validators"

	"github.com/gin-gonic/gin"
)

type PhotoCreate struct {
	Title    string `form:"title" json:"title" binding:"required"`
	Caption  string `form:"caption" json:"caption"`
	PhotoUrl string `form:"photo_url" json:"photo_url" binding:"required"`
}

func (photo *PhotoCreate) Bind(c *gin.Context) error {
	err := validators.Bind(c, photo)
	return err
}

// create
func PhotoCreateValidator() PhotoCreate {
	ret := PhotoCreate{}
	return ret
}
