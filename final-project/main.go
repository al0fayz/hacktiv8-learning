package main

import (
	"fmt"
	"hacktiv8-learning/final-project/config"
	"hacktiv8-learning/final-project/models"
	"hacktiv8-learning/final-project/router"
	"hacktiv8-learning/final-project/validators"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func main() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Can't load .env file!")
	}
	err = config.InitDB()
	if err != nil {
		log.Fatal("Error when connect to DB!")
	}
	//migrate table
	go AutoMigrate()

	//init validator
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("emailExist", validators.EmailExist)
		v.RegisterValidation("usernameExist", validators.UsernameExist)
	}
	port := os.Getenv("APP_PORT")
	//init router
	r := router.InitRouter()

	//serve server
	app := &http.Server{
		Addr:           ":" + port,
		Handler:        r,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	text := fmt.Sprintf("app run on port: %s", port)
	fmt.Println(text)
	//
	err = app.ListenAndServe()
	fmt.Println(err.Error())
}

func AutoMigrate() {
	//migrate table
	db := config.GetDb()
	userExist := db.Migrator().HasTable(&models.User{})
	if !userExist {
		// Create table for `User`
		db.Migrator().CreateTable(&models.User{})
		fmt.Println("table users created")
	}
	socialExist := db.Migrator().HasTable(&models.SocialMedia{})
	if !socialExist {
		// Create table for `SocialMedia`
		db.Migrator().CreateTable(&models.SocialMedia{})
		fmt.Println("table socialmedia created")
	}
	photoExist := db.Migrator().HasTable(&models.Photo{})
	if !photoExist {
		// Create table for `Photo`
		db.Migrator().CreateTable(&models.Photo{})
		fmt.Println("table photo created")
	}
	commentExist := db.Migrator().HasTable(&models.Comment{})
	if !commentExist {
		// Create table for `Comment`
		db.Migrator().CreateTable(&models.Comment{})
		fmt.Println("table comment created")
	}
}
