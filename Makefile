run:
	air

tidy:
	go mod tidy -v

test:
	go test -v -cover ./...

test-cov:
	go test -race -covermode=atomic -coverprofile=covprofile ./...

cov-htm:
	go tool cover -html=covprofile

cov-func:
	go tool cover -func=covprofile

publish:
	publish.sh
ENV_LOCAL_TEST=\
	DB_HOST=localhost \
	DB_PORT=5432 \
	DB_USER=postgres \
	DB_PASSWORD=1234 \
	DB_NAME=test_db \
	DB_SSLMODE=disable \
	DB_TIMEZONE=Asia/Shanghai


test.it:
	$(ENV_LOCAL_TEST) \
	go test -tags=integration -v -cover ./...
