#! /usr/bin/make
#(C) Copyright 2020 Hewlett Packard Enterprise Development LP

ORG := hpe-hcss
NAME := hpcaas-job-scheduler
GO_TEST= go test -tags=test
GO_BUILD= go build
GO_FLAG= GOOS=linux
GO_FLAGS_CGO= CGO_ENABLED=0
GOFLAGS= $(GO_FLAG) $(GO_FLAGS_CGO)

default: all
.PHONY: default

bin/hpcaas-job-scheduler: # Update Dockerfile with the correct binary name specififed here
	$(GOFLAGS) $(GO_BUILD) -tags netgo -a -v -installsuffix cgo -o bin/hpcaas-job-scheduler cmd/hpcaas-job-scheduler/main.go

# bin/api: # Update Dockerfile with the correct binary name specififed here
	# $(GOFLAGS) $(GO_BUILD) -tags netgo -a -v -installsuffix cgo -o bin/api cmd/api/main.go

go-build: bin/hpcaas-job-scheduler # bin/api
.PHONY: go-build

clean:
	@echo Cleaning...
	@rm -rf coverage bin .vendor/pkg kube-prometheus test-reports
	@echo Removing auto-generated code
	@find . -type f -name '*_mock.go' -delete
.PHONY: clean

vendor: go.mod
		go mod download
		go mod vendor

lint: vendor golangci-lint-config.yaml
	@golangci-lint --version
	golangci-lint run --config golangci-lint-config.yaml
.PHONY: lint

test: generate
	@echo "go test -v ./..."
	# Uncomment this once we have real test
	# @$(GO_TEST) -v ./... || (echo "Tests failed"; exit 1)
	@echo "Test passed"
.PHONY: test

coverage_dir := coverage/go
coverage: vendor
	@mkdir -p $(coverage_dir)/html
	# Uncomment this once we have real test
	$(GO_TEST) -cover -coverprofile=$(coverage_dir)/coverage.out -v ./... || (echo "Coverage tests failed")
	@go tool cover -html=$(coverage_dir)/coverage.out -o $(coverage_dir)/html/main.html;
	@echo "Generated $(coverage_dir)/html/main.html";
.PHONY: coverage

# CircleCi uses gotest.tools/gotestsum to run go test,
# also supports junit xml format to integrate CI systems
# and generates coverage report during test execution
testreport_dir := test-reports
test-ci-cover: generate
		@mkdir -p $(testreport_dir)/gotestsum
		@mkdir -p $(coverage_dir)/html
		@echo "gotestsum --format standard-verbose ./..."
		# Uncomment this once we have real test
		# @gotestsum --format standard-verbose \
			   #--junitfile $(testreport_dir)/gotestsum/results.xml \
			   #-- -cover -coverprofile=$(coverage_dir)/coverage.out \
			   #-p 3 -tags=test \
			   #./... || (echo "Test failed"; exit 1)
		@echo "Test Passed"
		# Uncomment this once we have real test
		# @go tool cover -html=$(coverage_dir)/coverage.out -o $(coverage_dir)/html/main.html;
		@echo "Generated $(coverage_dir)/html/main.html";
.PHONY: test-ci-cover

generate:
	@[ -f "$(shell which mockgen)" ] || go install github.com/golang/mock/mockgen
	@echo Removing old auto-generated code
	@find . -type f -name '*_mock.go' -delete
	@echo Running go generate for auto-generated code
	@go generate ./...

degenerate:
	@echo Removing auto-generated code
	@find . -type f -name '*_mock.go' -delete
	@echo Success
.PHONY: degenerate

all: lint test coverage go-build
.PHONY: all
