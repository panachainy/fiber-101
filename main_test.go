//go:build test_all || integration
// +build test_all integration

package main

import (
	"fiber-101/utils"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	tests := []struct {
		description   string
		route         string
		expectedError bool
		expectedCode  int
		expectedBody  string
	}{
		{
			description:   "index route",
			route:         "/",
			expectedError: false,
			expectedCode:  200,
			expectedBody:  "OK",
		},
		{
			description:   "non existing route",
			route:         "/i-dont-exist",
			expectedError: false,
			expectedCode:  404,
			expectedBody:  "Cannot GET /i-dont-exist",
		},
	}

	app := SetupApp()

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
