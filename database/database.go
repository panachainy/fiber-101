package database

import (
	"fiber-101/config"
	"fiber-101/model"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func Init() {
	var err error
	var dsn string

	if config.Get("DATABASE_DSN") != "" {
		dsn = config.Get("DATABASE_DSN")
		log.Println("Use config DATABASE_DSN")
	} else {
		dsn = fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v",
			config.Get("DB_HOST"),
			config.Get("DB_USER"),
			config.Get("DB_PASSWORD"),
			config.Get("DB_NAME"),
			config.Get("DB_PORT"),
			config.Get("DB_SSLMODE"),
			config.Get("DB_TIMEZONE"))
		log.Println("Use split config DATABASE")
	}
	DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	err = DBConn.AutoMigrate(&model.Product{})
	if err != nil {
		panic("Failed to auto migrate database")
	}
}
