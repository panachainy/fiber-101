package utils

type IntegrationTestModel struct {
	Description   string
	Route         string
	ExpectedError bool
	ExpectedCode  int
	ExpectedBody  string
}
