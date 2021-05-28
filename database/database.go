package database

import (
	"fiber-101/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func Init() {
	var err error

	dsn := "host=localhost user=postgres password=1234 dbname=gorm_test port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DBConn.AutoMigrate(&model.Product{})
}
