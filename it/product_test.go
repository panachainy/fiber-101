package it

import (
	"fiber-101/utils"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: try to make integration test.
func TestGetProducts(t *testing.T) {
	tests := []struct {
		mock          func()
		description   string
		route         string
		method        string
		body          string
		expectedError bool
		expectedCode  int
		expectedBody  string
	}{
		{
			mock: func() {
				// product := &models.Product{
				// 	Code:          "",
				// 	Price:         2,
				// 	PriceDetailJa: 1,
				// }

				// db := utils.DBConn
				CleanDatabase()
				// db.Create(&product)
			},
			description:   "index route",
			route:         "/products",
			method:        "GET",
			expectedError: false,
			expectedCode:  200,
			expectedBody:  "{\"data\":[]}",
		},
		{
			mock:          func() {},
			description:   "index route",
			route:         "/products",
			method:        "POST",
			body:          "",
			expectedError: false,
			expectedCode:  200,
			expectedBody:  "{\"message\":\"Unprocessable Entity\"}",
		},
	}

	app := utils.SetupApp("../example.env")

	for _, tt := range tests {
		tt.mock()

		req, _ := http.NewRequest(
			tt.method,
			tt.route,
			nil,
		)

		res, _ := app.Test(req)

		body, err := ioutil.ReadAll(res.Body)

		assert.Equalf(t, tt.expectedError, err != nil, tt.description)

		assert.Nilf(t, err, tt.description)
		assert.Equalf(t, tt.expectedBody, string(body), tt.description)
	}
}
