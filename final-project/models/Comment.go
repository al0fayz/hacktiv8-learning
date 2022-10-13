package models

import (
	"hacktiv8-learning/final-project/config"
	"hacktiv8-learning/final-project/midlleware"
	"time"

	"github.com/gin-gonic/gin"
)

type Comment struct {
	Id        int64     `json:"id" gorm:"primaryKey"`
	UserId    int64     `json:"user_id"`
	PhotoId   int64     `json:"photo_id"`
	Message   string    `json:"message" gorm:"not null;type:varchar(191)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      *User     `json:"user" gorm:"ForeignKey:UserId"`
	Photo     *Photo    `json:"photo" gorm:"ForeignKey:PhotoId"`
}

func CreateComment(com *Comment) error {
	com.CreatedAt = time.Now()
	com.UpdatedAt = time.Now()
	db := config.GetDb()
	err := db.Create(com).Error
	return err
}
func GetAllCommentByUserId(com *[]Comment, c *gin.Context) error {
	db := config.GetDb()
	user_id, err := midlleware.ExtractTokenID(c)
	if err != nil {
		return err
	}
	err = db.Model(&Comment{}).Preload("User").Preload("Photo").Where("user_id = ?", user_id).Find(com).Error
	return err
}
func FindCommentById(id string) (Comment, error) {
	db := config.GetDb()
	var model Comment
	err := db.Where("id = ?", id).First(&model).Error
	return model, err
}

func UpdateComment(com *Comment) error {
	com.UpdatedAt = time.Now()
	db := config.GetDb()
	err := db.Model(&Comment{}).Where("id = ?", com.Id).Updates(com).Error
	return err
}

func DeleteComment(id string) error {
	db := config.GetDb()
	var model Comment
	err := db.Where("id = ?", id).Delete(&model).Error
	return err
}
