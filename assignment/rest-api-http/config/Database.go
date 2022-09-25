package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

type DBConfig struct {
	Host       string
	Port       int
	User       string
	DBDatabase string
	Password   string
}

func InitDB() error {
	var err error
	Db, err = sql.Open("mysql", MysqlConfig(InsertConfig()))
	if err != nil {
		return err
	}
	// See "Important settings" section.
	Db.SetConnMaxLifetime(time.Minute * 3)
	Db.SetMaxOpenConns(10)
	Db.SetMaxIdleConns(10)
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
func GetDb() *sql.DB {
	return Db
}
