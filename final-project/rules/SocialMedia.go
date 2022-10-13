package rules

import (
	"hacktiv8-learning/final-project/validators"

	"github.com/gin-gonic/gin"
)

type SocialMediaCreate struct {
	Name           string `form:"name" json:"name" binding:"required"`
	SocialMediaUrl string `form:"social_media_url" json:"social_media_url" binding:"required"`
}

func (soc *SocialMediaCreate) Bind(c *gin.Context) error {
	err := validators.Bind(c, soc)
	return err
}

// create
func SocialMediaCreateValidator() SocialMediaCreate {
	ret := SocialMediaCreate{}
	return ret
}
