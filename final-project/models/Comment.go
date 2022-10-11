package models

import "time"

type Comment struct {
	Id        int64     `json:"id" gorm:"primaryKey"`
	UserId    int64     `json:"user_id"`
	PhotoId   int64     `json:"photo_id"`
	Message   string    `json:"message" gorm:"not null;type:varchar(191)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
