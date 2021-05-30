run:
	air

tidy:
	go mod tidy -v

test:
	go test -v -cover

test-cov:
	go test -coverprofile=coverage.out

cov-htm:
	go tool cover -html=coverage.out

cov-func:
	go tool cover -func=coverage.out
