package models

import (
	"hacktiv8-learning/final-project/config"
	"hacktiv8-learning/final-project/midlleware"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Photo struct {
	Id        int64     `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null;type:varchar(191)"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url" gorm:"not null;type:varchar(191)"`
	UserId    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      *User     `json:"user" gorm:"ForeignKey:UserId"`
}

func CreatePhoto(photo *Photo) error {
	photo.CreatedAt = time.Now()
	photo.UpdatedAt = time.Now()
	db := config.GetDb()
	err := db.Create(photo).Error
	return err
}
func GetAllPhotoByUserId(photo *[]Photo, c *gin.Context) error {
	db := config.GetDb()
	user_id, err := midlleware.ExtractTokenID(c)
	if err != nil {
		return err
	}
	err = db.Model(&Photo{}).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("Id,Email,Username")
	}).Where("user_id = ?", user_id).Find(photo).Error

	return err
}
func FindPhotoById(id string) (Photo, error) {
	db := config.GetDb()
	var model Photo
	err := db.Where("id = ?", id).First(&model).Error
	return model, err
}

func UpdatePhoto(photo *Photo) error {
	photo.UpdatedAt = time.Now()
	db := config.GetDb()
	err := db.Model(&Photo{}).Where("id = ?", photo.Id).Updates(photo).Error
	return err
}

func DeletePhoto(id string) error {
	db := config.GetDb()
	var model Photo
	err := db.Where("id = ?", id).Delete(&model).Error
	return err
}
