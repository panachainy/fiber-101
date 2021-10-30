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
	go test -tags=test_all -v -cover ./...

test.cov:
	$(ENV_LOCAL_TEST) \
	go test -tags=test_all -v -race -covermode=atomic -coverprofile=covprofile ./...

test.it:
	$(ENV_LOCAL_TEST) \
	go test -tags=integration -v -cover ./...

cov.htm:
	go tool cover -html=covprofile

cov.func:
	go tool cover -func=covprofile

publish:
	publish.sh



# this command will trigger integration test
# INTEGRATION_TEST_SUITE_PATH is used for run specific test in Golang, if it's not specified
# it will run all tests under ./it directory
test.integration:
	$(ENV_LOCAL_TEST) \
	go test -tags=integration $(INTEGRATION_TEST_PATH) -count=1 # -run=$(INTEGRATION_TEST_SUITE_PATH)

# this command will trigger integration test with verbose mode
test.integration.debug:
	$(ENV_LOCAL_TEST) \
	go test -tags=integration $(INTEGRATION_TEST_PATH) -count=1 -v -run=$(INTEGRATION_TEST_SUITE_PATH)
