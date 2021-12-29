package database

import (
	"fiber-101/config"
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

	if config.C.DATABASE_DSN != "" {
		logrus.Infoln("[DB-CONFIG] Use config DATABASE_DSN")
		dsn = config.C.DATABASE_DSN
	} else {
		logrus.Infoln("[DB-CONFIG] Use split config DATABASE")

		dsn = fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v",
			config.C.DATABASE_HOST,
			config.C.DATABASE_USER,
			config.C.DATABASE_PASSWORD,
			config.C.DATABASE_NAME,
			config.C.DATABASE_PORT,
			config.C.DATABASE_SSLMODE,
			config.C.DATABASE_TIMEZONE)
	}

	logrus.Debug("[DB-CONFIG] DSN: ", dsn)

	DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	err = DBConn.AutoMigrate(&models.Product{})
	if err != nil {
		panic("Failed to auto migrate database")
	}
}
