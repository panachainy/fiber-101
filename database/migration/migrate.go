package migration

import (
	"fiber-101/pkg/products"

	"gorm.io/gorm"
)

func Migrate(DBConn *gorm.DB) {
	err := DBConn.AutoMigrate(&products.Product{})
	if err != nil {
		panic("Failed to auto migrate database")
	}
}
