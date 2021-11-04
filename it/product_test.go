//go:build test_all || integration
// +build test_all integration

package it

import (
	"fiber-101/utils"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
			description:   "index route",
			route:         "/products",
			method:        "GET",
			expectedError: false,
			expectedCode:  200,
			expectedBody:  "{\"data\":[]}",
		}, {
			description:   "index route",
			route:         "/products",
			method:        "POST",
			body:          "",
			expectedError: false,
			expectedCode:  200,
			expectedBody:  "{\"data\":[]}",
		},
	}

	app := utils.SetupApp()

	for _, tt := range tests {

		// if tt.mock {
		// 	tt.mock()
		// }

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
