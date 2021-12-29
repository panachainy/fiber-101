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

		logrus.Debug("[DB-CONFIG] DATABASE_HOST: ", config.Get("DATABASE_HOST"))
		logrus.Debug("[DB-CONFIG] DATABASE_USER: ", config.Get("DATABASE_USER"))
		logrus.Debug("[DB-CONFIG] DATABASE_PASSWORD: ", config.Get("DATABASE_PASSWORD"))
		logrus.Debug("[DB-CONFIG] DATABASE_NAME: ", config.Get("DATABASE_NAME"))
		logrus.Debug("[DB-CONFIG] DATABASE_PORT: ", config.Get("DATABASE_PORT"))
		logrus.Debug("[DB-CONFIG] DATABASE_SSLMODE: ", config.Get("DATABASE_SSLMODE"))
		logrus.Debug("[DB-CONFIG] DATABASE_TIMEZONE: ", config.Get("DATABASE_TIMEZONE"))

		dsn = fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v",
			config.Get("DATABASE_HOST"),
			config.Get("DATABASE_USER"),
			config.Get("DATABASE_PASSWORD"),
			config.Get("DATABASE_NAME"),
			config.Get("DATABASE_PORT"),
			config.Get("DATABASE_SSLMODE"),
			config.Get("DATABASE_TIMEZONE"))
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
