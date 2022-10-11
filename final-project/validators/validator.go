package validators

import (
	"hacktiv8-learning/final-project/models"

	"github.com/go-playground/validator/v10"
)

// check email exist in DB User
func EmailExist(fl validator.FieldLevel) bool {
	email := fl.Field().Interface().(string)
	//check email exist
	_, err := models.FindUserByEmail(email)

	return !(err == nil)

}

// check username exist in Db User
func UsernameExist(fl validator.FieldLevel) bool {
	username := fl.Field().Interface().(string)
	//check username exist
	_, err := models.FindUserByUsername(username)
	return !(err == nil)
}
