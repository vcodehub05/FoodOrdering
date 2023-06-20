TEST_REPORT=coverage.out
TEST_REPORT_TMP=coverage.tmp.out

help:
	@echo  ""
	@echo  "make lint       --run all linters against HEAD"
	@echo  "make lint-diff  --run all linters against diff from merge-base"
	@echo  "make lint-clean --cleans linter cache"

.PHONY: lint
lint:
	golangci-lint run ./... -v

.PHONY: lint-diff
lint-diff:
	golangci-lint run ./... --new-from-rev=$$(git merge-base origin/demo HEAD)

.PHONY: lint-clean
lint-clean:
	golangci-lint cache clean

.PHONY: test
test:
	go test -race -v $$(go list ./... | grep -v mocks*) -coverprofile=${TEST_REPORT_TMP}
	cat ${TEST_REPORT_TMP} | grep -v "_gen.go" > ${TEST_REPORT}
	go tool cover -func=${TEST_REPORT}

.PHONY: test
generate:
	buf generate
run:
	go run application/customer/*.go
	go run application/resturant/*.go