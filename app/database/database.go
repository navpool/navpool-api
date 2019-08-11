package database

import (
	"errors"
	"fmt"
	"github.com/NavPool/navpool-api/app/config"
	"github.com/NavPool/navpool-api/app/helpers"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

var (
	ErrorDatabaseConnection = errors.New("Failed to connect to the database")
)

var DB *gorm.DB

func GetConnection() *gorm.DB {
	return DB
}

func CreateConnection() (err error) {
	args := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=%s",
		config.Get().DB.Host,
		config.Get().DB.Port,
		config.Get().DB.DbName,
		config.Get().DB.Username,
		config.Get().DB.Password,
		config.Get().DB.SSLMode)

	DB, err = gorm.Open(config.Get().DB.Dialect, args)
	if err != nil {
		helpers.LogError(err)
		return ErrorDatabaseConnection
	}

	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
	DB.DB().SetConnMaxLifetime(time.Hour)

	if config.Get().Debug == true {
		DB.LogMode(config.Get().DB.LogMode)
	}

	return
}
