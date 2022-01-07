package products

import (
	"fiber-101/models"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code          string `validate:"required,min=3,max=10"`
	Price         uint
	PriceDetailJa int
}

func ValidateStruct(product Product) []*models.ErrorResponse {
	var errors []*models.ErrorResponse
	validate := validator.New()
	err := validate.Struct(product)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element models.ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
