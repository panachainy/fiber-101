package products

import (
	"fiber-101/database"
	"fiber-101/products/entities"
	"fmt"

	"gorm.io/gorm"
)

type ProductRepository interface {
	GetAll() []entities.Product

	RemoveById(uid string) error
	AddNewRegister(u entities.Product) error
}

func NewProductRepo(DBConn *gorm.DB) ProductRepository {
	return &productRepoImp{DBConn: DBConn}
}

type productRepoImp struct {
	DBConn *gorm.DB
}

func (r *productRepoImp) AddNewRegister(product entities.Product) error {
	r.DBConn.Create(&product)

	return nil
}

func (r *productRepoImp) RemoveById(uid string) error {
	var product entities.Product
	r.DBConn.First(&product, uid)

	if product.ID == 0 {
		return fmt.Errorf("not found product id %v", uid)
	}

	r.DBConn.Delete(&product)
	return nil
}

func (r *productRepoImp) GetAll() []entities.Product {
	var product []entities.Product
	database.DBConn.Find(&product)
	return product
}
