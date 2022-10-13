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

func CreateSocialMedia(c *gin.Context) {
	req := rules.SocialMediaCreateValidator()
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
		//create social media
		soc := models.SocialMedia{
			Name:           req.Name,
			SocialMediaUrl: req.SocialMediaUrl,
			UserId:         int64(user_id),
		}
		err = models.CreateSocialMedia(&soc)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":     http.StatusBadRequest,
				"errors":   err.Error(),
				"messages": "Bad Request",
			})
			return
		}
		c.JSON(http.StatusCreated, gin.H{
			"id":               soc.Id,
			"name":             soc.Name,
			"social_media_url": soc.SocialMediaUrl,
			"user_id":          soc.UserId,
			"created_at":       soc.CreatedAt,
		})
	}
}

func GetALlSocialMedia(c *gin.Context) {
	var socialMedia []models.SocialMedia
	err := models.GetAllSocialMediaByUserId(&socialMedia, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":     http.StatusInternalServerError,
			"errors":   err.Error(),
			"messages": "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": socialMedia,
	})
}

func UpdateSocialMedia(c *gin.Context) {
	req := rules.SocialMediaCreateValidator()
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
		//search social media
		id := c.Params.ByName("id")
		soc, err := models.FindSocialMediaById(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":     http.StatusBadRequest,
				"errors":   err.Error(),
				"messages": "Bad Request, Social Media not found",
			})
			return
		}
		soc.Name = req.Name
		soc.SocialMediaUrl = req.SocialMediaUrl

		err = models.UpdateSocialMedia(&soc)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":     http.StatusBadRequest,
				"errors":   err.Error(),
				"messages": "Bad Request",
			})
			return
		}
		c.JSON(http.StatusCreated, gin.H{
			"id":               soc.Id,
			"name":             soc.Name,
			"social_media_url": soc.SocialMediaUrl,
			"user_id":          soc.UserId,
			"updated_at":       soc.UpdatedAt,
		})
	}
}

func DeleteSocialMedia(c *gin.Context) {
	//search social media is exist
	id := c.Params.ByName("id")
	_, err := models.FindSocialMediaById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":     http.StatusBadRequest,
			"errors":   err.Error(),
			"messages": "Bad Request",
		})
		return
	}
	//delete
	err = models.DeleteSocialMedia(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":     http.StatusBadRequest,
			"errors":   err.Error(),
			"messages": "Bad Request",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"messages": "Your social media has been successfuly deleted",
	})
}
