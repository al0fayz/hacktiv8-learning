package models

import "time"

type Comment struct {
	Id        int64     `json:"id"`
	UserId    int64     `json:"user_id"`
	PhotoId   int64     `json:"photo_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
