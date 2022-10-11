package models

import (
	"hacktiv8-learning/final-project/config"
	"time"
)

type User struct {
	Id        int64     `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Age       int32     `json:"age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FindUserByEmail(email string) (User, error) {
	db := config.GetDb()
	var model User
	err := db.Where("email = ?", email).First(&model).Error
	return model, err
}
func FindUserByUsername(username string) (User, error) {
	db := config.GetDb()
	var model User
	err := db.Where("username = ?", username).First(&model).Error
	return model, err
}
