package models

import "time"

type Photo struct {
	Id        int64     `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null;type:varchar(191)"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url" gorm:"not null;type:varchar(191)"`
	UserId    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
