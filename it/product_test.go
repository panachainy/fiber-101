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
		description   string
		route         string
		expectedError bool
		expectedCode  int
		expectedBody  string
	}{
		{
			description:   "index route",
			route:         "/products",
			expectedError: false,
			expectedCode:  200,
			expectedBody:  "{\"data\":[]}",
		},
	}

	app := utils.SetupApp()

	for _, test := range tests {
		req, _ := http.NewRequest(
			"GET",
			test.route,
			nil,
		)

		res, _ := app.Test(req)

		body, err := ioutil.ReadAll(res.Body)

		assert.Equalf(t, test.expectedError, err != nil, test.description)

		assert.Nilf(t, err, test.description)
		assert.Equalf(t, test.expectedBody, string(body), test.description)
	}
}
