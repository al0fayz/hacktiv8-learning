package rules

import (
	"hacktiv8-learning/final-project/validators"

	"github.com/gin-gonic/gin"
)

type CommentCreate struct {
	Message string `form:"message" json:"message" binding:"required"`
	PhotoId int64  `form:"photo_id" json:"photo_id"`
}

func (com *CommentCreate) Bind(c *gin.Context) error {
	err := validators.Bind(c, com)
	return err
}

// create
func CommentCreateValidator() CommentCreate {
	ret := CommentCreate{}
	return ret
}

type CommentUpdate struct {
	Message string `form:"message" json:"message" binding:"required"`
}

func (com *CommentUpdate) Bind(c *gin.Context) error {
	err := validators.Bind(c, com)
	return err
}

// create
func CommentUpdateValidator() CommentUpdate {
	ret := CommentUpdate{}
	return ret
}
