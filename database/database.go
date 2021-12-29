package database

import (
	"fiber-101/config"
	"fiber-101/model"
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

// TODO: upgrade to support sqlite for test

func Init() {
	var err error
	var dsn string

	if config.Get("DATABASE_DSN") != "" {
		logrus.Infoln("[DB-CONFIG] Use config DATABASE_DSN")
		logrus.Debug("[DB-CONFIG] DATABASE_DSN: ", config.Get("DATABASE_DSN"))

		dsn = config.Get("DATABASE_DSN")

	} else {
		logrus.Infoln("[DB-CONFIG] Use split config DATABASE")

		logrus.Debug("[DB-CONFIG] DB_HOST: ", config.Get("DB_HOST"))
		logrus.Debug("[DB-CONFIG] DB_USER: ", config.Get("DB_USER"))
		logrus.Debug("[DB-CONFIG] DB_PASSWORD: ", config.Get("DB_PASSWORD"))
		logrus.Debug("[DB-CONFIG] DB_NAME: ", config.Get("DB_NAME"))
		logrus.Debug("[DB-CONFIG] DB_PORT: ", config.Get("DB_PORT"))
		logrus.Debug("[DB-CONFIG] DB_SSLMODE: ", config.Get("DB_SSLMODE"))
		logrus.Debug("[DB-CONFIG] DB_TIMEZONE: ", config.Get("DB_TIMEZONE"))

		dsn = fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v",
			config.Get("DB_HOST"),
			config.Get("DB_USER"),
			config.Get("DB_PASSWORD"),
			config.Get("DB_NAME"),
			config.Get("DB_PORT"),
			config.Get("DB_SSLMODE"),
			config.Get("DB_TIMEZONE"))
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
