run:
	air

tidy:
	go mod tidy -v

test:
	go test -race

test-in:
	go test -race -tags=integration
