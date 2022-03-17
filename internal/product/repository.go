package product

import (
	"fmt"

	"fiber-101/database"

	"gorm.io/gorm"
)

type ProductRepository interface {
	GetAll() []Product

	// TODO: change interface follow handler requirement
	RemoveById(uid string) error
	AddNewRegister(u Product) error
}

func NewProductRepo(DBConn *gorm.DB) ProductRepository {
	return &productRepoImp{DBConn: DBConn}
}

type productRepoImp struct {
	DBConn *gorm.DB
}

func (r *productRepoImp) AddNewRegister(product Product) error {
	r.DBConn.Create(&product)

	return nil
}

func (r *productRepoImp) RemoveById(uid string) error {
	var product Product
	r.DBConn.First(&product, uid)

	if product.ID == 0 {
		return fmt.Errorf("not found product id %v", uid)
	}

	r.DBConn.Delete(&product)
	return nil
}

func (r *productRepoImp) GetAll() []Product {
	var product []Product
	database.DBConn.Find(&product)
	return product
}
