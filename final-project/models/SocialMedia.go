package models

type SocialMedia struct {
	Id             int64  `json:"id" gorm:"primaryKey"`
	Name           string `json:"name" gorm:"not null;type:varchar(191)"`
	SocialMediaUrl string `json:"social_media_url" gorm:"not null;type:varchar(191)"`
	UserId         int64  `json:"user_id"`
}
