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

func CreateComment(c *gin.Context) {
	req := rules.CommentCreateValidator()
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
		//search photo by id
		_, err := models.FindPhotoById(fmt.Sprint(req.PhotoId))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":     http.StatusBadRequest,
				"errors":   err.Error(),
				"messages": "Bad Request, photo not found",
			})
			return
		}
		user_id, err := midlleware.ExtractTokenID(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":     http.StatusInternalServerError,
				"errors":   err.Error(),
				"messages": "Internal Server Error",
			})
			return
		}
		comment := models.Comment{
			UserId:  int64(user_id),
			PhotoId: req.PhotoId,
			Message: req.Message,
		}
		//create comment
		err = models.CreateComment(&comment)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":     http.StatusBadRequest,
				"errors":   err.Error(),
				"messages": "Bad Request",
			})
			return
		}
		c.JSON(http.StatusCreated, gin.H{
			"id":         comment.Id,
			"message":    comment.Message,
			"photo_id":   comment.PhotoId,
			"user_id":    comment.UserId,
			"created_at": comment.CreatedAt,
		})
	}
}

func GetAllComment(c *gin.Context) {
	var comment []models.Comment
	err := models.GetAllCommentByUserId(&comment, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":     http.StatusInternalServerError,
			"errors":   err.Error(),
			"messages": "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": comment,
	})
}

func UpdateComment(c *gin.Context) {
	req := rules.CommentUpdateValidator()
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
		//check comment is exist
		id := c.Params.ByName("id")
		com, err := models.FindCommentById(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":     http.StatusBadRequest,
				"errors":   err.Error(),
				"messages": "Bad Request",
			})
			return
		}
		com.Message = req.Message
		err = models.UpdateComment(&com)
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
			"id":         com.Id,
			"message":    com.Message,
			"photo_id":   com.PhotoId,
			"user_id":    com.UserId,
			"updated_at": com.UpdatedAt,
		})
	}
}
func DeleteComment(c *gin.Context) {
	//search photo is exist
	id := c.Params.ByName("id")
	_, err := models.FindCommentById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":     http.StatusBadRequest,
			"errors":   err.Error(),
			"messages": "Bad Request",
		})
		return
	}
	//delete
	err = models.DeleteComment(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":     http.StatusBadRequest,
			"errors":   err.Error(),
			"messages": "Bad Request",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"messages": "Your comment has been successfuly deleted",
	})
}
