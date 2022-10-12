package handler

import (
	"errors"
	"hacktiv8-learning/final-project/midlleware"
	"hacktiv8-learning/final-project/models"
	"hacktiv8-learning/final-project/rules"
	"hacktiv8-learning/final-project/validators"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreatePhoto(c *gin.Context) {
	req := rules.PhotoCreateValidator()
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
		user_id, err := midlleware.ExtractTokenID(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":     http.StatusInternalServerError,
				"errors":   err.Error(),
				"messages": "Internal Server Error",
			})
			return
		}
		//create photo
		photo := models.Photo{
			Title:    req.Title,
			Caption:  req.Caption,
			PhotoUrl: req.PhotoUrl,
			UserId:   int64(user_id),
		}
		err = models.CreatePhoto(&photo)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":     http.StatusInternalServerError,
				"errors":   err.Error(),
				"messages": "Internal Server Error",
			})
			return
		}
		c.JSON(http.StatusCreated, gin.H{
			"id":         photo.Id,
			"title":      photo.Title,
			"caption":    photo.Caption,
			"photo_url":  photo.PhotoUrl,
			"user_id":    photo.UserId,
			"created_at": photo.CreatedAt,
		})
	}
}
func GetAllPhoto(c *gin.Context) {
	var photo []models.Photo
	err := models.GetAllPhotoByUserId(&photo, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":     http.StatusInternalServerError,
			"errors":   err.Error(),
			"messages": "Internal Server Error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": photo,
	})
}
