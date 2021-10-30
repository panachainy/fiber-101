INTEGRATION_TEST_PATH?=./it

ENV_LOCAL_TEST=\
	DB_HOST=localhost \
	DB_PORT=5432 \
	DB_USER=postgres \
	DB_PASSWORD=1234 \
	DB_NAME=test_db \
	DB_SSLMODE=disable \
	DB_TIMEZONE=Asia/Shanghai

run:
	air

tidy:
	go mod tidy -v

test:
	$(ENV_LOCAL_TEST) \
	go test -tags=test_all -cover ./...

test.debug:
	$(ENV_LOCAL_TEST) \
	go test -tags=test_all -v -cover ./...

test.cov:
	$(ENV_LOCAL_TEST) \
	go test -tags=test_all -v -race -covermode=atomic -coverprofile=covprofile ./...

test.it:
	$(ENV_LOCAL_TEST) \
	go test -tags=integration -v -cover ./...

test.unit:
	$(ENV_LOCAL_TEST) \
	go test -tags=unit -v -cover ./...

cov.htm:
	go tool cover -html=covprofile

cov.func:
	go tool cover -func=covprofile

publish:
	publish.sh
