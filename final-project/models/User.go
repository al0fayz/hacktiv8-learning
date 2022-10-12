package models

import (
	"hacktiv8-learning/final-project/config"
	"hacktiv8-learning/final-project/utils"
	"time"

	"gorm.io/gorm"
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

func (u *User) BeforeCreate(tx *gorm.DB) error {
	hash, err := utils.GenerateHash(u.Password)
	if err != nil {
		return err
	}
	u.Password = hash
	return nil
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

func CreateUser(user *User) error {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	db := config.GetDb()
	err := db.Create(user).Error
	return err
}
