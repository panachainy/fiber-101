//go:build test_all || integration
// +build test_all integration

package main

import (
	"fiber-101/utils"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	tests := []utils.IntegrationTestModel{
		{
			Description:   "index route",
			Route:         "/",
			ExpectedError: false,
			ExpectedCode:  200,
			ExpectedBody:  "OK",
		},
		{
			Description:   "non existing route",
			Route:         "/i-dont-exist",
			ExpectedError: false,
			ExpectedCode:  404,
			ExpectedBody:  "Cannot GET /i-dont-exist",
		},
	}

	app := SetupApp()

	logrus.Infoln("Running tes====", tests)

	for _, test := range tests {
		req, _ := http.NewRequest(
			"GET",
			test.Route,
			nil,
		)

		res, _ := app.Test(req)

		body, err := ioutil.ReadAll(res.Body)

		assert.Equalf(t, test.ExpectedError, err != nil, test.Description)

		assert.Nilf(t, err, test.Description)
		assert.Equalf(t, test.ExpectedBody, string(body), test.Description)
	}
}
