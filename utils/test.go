package utils

type TestModel struct {
	Description string

	// Test input
	Route string

	// Expected output
	ExpectedError bool
	ExpectedCode  int
	ExpectedBody  string
}
