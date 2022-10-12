package utils

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GenerateHash(password string) (string, error) {
	pwd := []byte(password)
	hashResult, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	//to be string
	hashString := string(hashResult)

	return hashString, nil
}

// data from DB and password input
func CompareHash(data string, password string) error {
	myPassword := []byte(data)
	inputPassword := []byte(password)

	err := bcrypt.CompareHashAndPassword(myPassword, inputPassword)

	return err
}

func GetContentType(c *gin.Context) string {
	return c.Request.Header.Get("Content-Type")
}
