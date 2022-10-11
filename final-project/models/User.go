package models

import (
	"hacktiv8-learning/final-project/config"
	"time"
)

type User struct {
	Id        int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	Username  string    `json:"username" gorm:"not null;uniqe;type:varchar(191)"`
	Email     string    `json:"email" gorm:"not null;uniqe;type:varchar(191)"`
	Password  string    `json:"password" gorm:"not null;type:varchar(191)"`
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
