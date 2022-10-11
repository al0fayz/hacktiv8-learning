package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

type DBConfig struct {
	Host       string
	Port       int
	User       string
	DBDatabase string
	Password   string
}

func InitDB() error {
	var err error
	Db, err = gorm.Open(mysql.Open(MysqlConfig(InsertConfig())), &gorm.Config{})

	db, _ := Db.DB()
	db.SetMaxIdleConns(100)
	db.SetMaxOpenConns(750)
	db.SetConnMaxIdleTime(time.Minute)

	if err != nil {
		return err
	}
	return nil
}
func InsertConfig() *DBConfig {
	//covert string to int beacuse in DBConfig struct we defined port int
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	dbConfig := DBConfig{
		Host:       os.Getenv("DB_HOST"),
		Port:       port,
		User:       os.Getenv("DB_USERNAME"),
		DBDatabase: os.Getenv("DB_DATABASE"),
		Password:   os.Getenv("DB_PASSWORD"),
	}
	return &dbConfig
}
func MysqlConfig(dbConfig *DBConfig) string {
	//sprintf return a string format
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBDatabase,
	)
}

// use in models
func GetDb() *gorm.DB {
	return Db
}
