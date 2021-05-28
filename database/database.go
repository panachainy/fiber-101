package database

import (
	"fiber-101/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func InitDatabase() {
	var err error
	DBConn, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DBConn.AutoMigrate(&model.Product{})
}
