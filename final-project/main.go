package main

import (
	"fmt"
	"hacktiv8-learning/final-project/config"
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
