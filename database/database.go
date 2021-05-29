package database

import (
	"fiber-101/config"
	"fiber-101/model"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func Init() {
	var err error

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v",
		config.Get("DB_HOST"),
		config.Get("DB_USER"),
		config.Get("DB_PASSWORD"),
		config.Get("DB_NAME"),
		config.Get("DB_PORT"),
		config.Get("DB_SSLMODE"),
		config.Get("DB_TIMEZONE"))

	DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	err = DBConn.AutoMigrate(&model.Product{})
	if err != nil {
		panic("Failed to auto migrate database")
	}
}
