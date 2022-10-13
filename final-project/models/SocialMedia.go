package models

import (
	"hacktiv8-learning/final-project/config"
	"hacktiv8-learning/final-project/midlleware"

	"github.com/gin-gonic/gin"
)

type SocialMedia struct {
	Id             int64  `json:"id" gorm:"primaryKey"`
	Name           string `json:"name" gorm:"not null;type:varchar(191)"`
	SocialMediaUrl string `json:"social_media_url" gorm:"not null;type:varchar(191)"`
	UserId         int64  `json:"user_id"`
}

func CreateSocialMedia(soc *SocialMedia) error {
	db := config.GetDb()
	err := db.Create(soc).Error
	return err
}
func GetAllSocialMediaByUserId(soc *[]SocialMedia, c *gin.Context) error {
	db := config.GetDb()
	user_id, err := midlleware.ExtractTokenID(c)
	if err != nil {
		return err
	}
	err = db.Model(&SocialMedia{}).Where("user_id = ?", user_id).Find(soc).Error
	return err
}
func FindSocialMediaById(id string) (SocialMedia, error) {
	db := config.GetDb()
	var model SocialMedia
	err := db.Where("id = ?", id).First(&model).Error
	return model, err
}

func UpdateSocialMedia(soc *SocialMedia) error {
	db := config.GetDb()
	err := db.Model(&SocialMedia{}).Where("id = ?", soc.Id).Updates(soc).Error
	return err
}

func DeleteSocialMedia(id string) error {
	db := config.GetDb()
	var model SocialMedia
	err := db.Where("id = ?", id).Delete(&model).Error
	return err
}
