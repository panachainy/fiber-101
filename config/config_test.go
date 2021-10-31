//go:build test_all || unit
// +build test_all unit

package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConfig(t *testing.T) {
	tests := []struct {
		description      string
		key              string
		envPath          string
		expectedError    bool
		expectedResponse string
	}{
		{
			description:      "when get existing env key from file should be 5432",
			key:              "DB_PORT",
			envPath:          ".env.example",
			expectedError:    false,
			expectedResponse: "5432",
		},
		{
			description:      "when get not existing env key from file should be empty",
			key:              "xxxxxx",
			envPath:          ".env.example",
			expectedError:    false,
			expectedResponse: "",
		},
	}

	for _, test := range tests {
		res := Get(test.key, test.envPath)

		assert.Equalf(t, test.expectedResponse, res, test.description)
	}
}
