run:
	air

tidy:
	go mod tidy -v

test:
	go test -v -cover ./...

test-cov:
	go test -race -covermode atomic -coverprofile=covprofile ./...

cov-htm:
	go tool cover -html=covprofile

cov-func:
	go tool cover -func=covprofile

publish:
	publish.sh
