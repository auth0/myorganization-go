#-----------------------------------------------------------------------------------------------------------------------
# Variables (https://www.gnu.org/software/make/manual/html_node/Using-Variables.html#Using-Variables)
#-----------------------------------------------------------------------------------------------------------------------
.DEFAULT_GOAL := help
GO_BIN ?= $(shell go env GOPATH)/bin

#-----------------------------------------------------------------------------------------------------------------------
# Rules (https://www.gnu.org/software/make/manual/html_node/Rule-Introduction.html#Rule-Introduction)
#-----------------------------------------------------------------------------------------------------------------------
.PHONY: help

help: ## Show this help
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

#-----------------------------------------------------------------------------------------------------------------------
# Dependencies
#-----------------------------------------------------------------------------------------------------------------------
.PHONY: deps

deps: ## Download dependencies
	@echo "==> Downloading dependencies..."
	@go mod download

$(GO_BIN)/govulncheck:
	${call print, "Installing govulncheck"}
	@go install -v golang.org/x/vuln/cmd/govulncheck@latest

#-----------------------------------------------------------------------------------------------------------------------
# Checks
#-----------------------------------------------------------------------------------------------------------------------
.PHONY: fmt lint check-vuln

fmt: ## Format go files
	@echo "==> Formatting go files..."
	@gofmt -w -s .
	@echo "==> Done."

lint: ## Run golangci-lint on non-generated files
	@echo "==> Running linter..."
	@golangci-lint run \
		./client/ \
		./option/ \
		./internal/auth0/ \
		./internal/transport/ \
		./internal/telemetry/

check-vuln: $(GO_BIN)/govulncheck ## Check for vulnerabilities
	@echo "==> Checking for vulnerabilities..."
	@govulncheck -show verbose ./...

#-----------------------------------------------------------------------------------------------------------------------
# Testing
#-----------------------------------------------------------------------------------------------------------------------
.PHONY: test test-unit test-fuzz test-wiremock coverage

WIREMOCK_COMPOSE_FILE := wiremock/docker-compose.test.yml
WIREMOCK_PROJECT_NAME := wiremock-auth0sdk
COVERAGE_DIR := .coverage
FUZZ_TIME ?= 10s

test: test-unit test-fuzz test-wiremock coverage ## Run all tests (unit + fuzz + WireMock) and generate coverage. Usage: `make test FILTER="TestName"`

test-unit: ## Run unit tests (no Docker required). Usage: `make test-unit FILTER="TestName"`
	@echo "==> Running unit tests..."
	@mkdir -p $(COVERAGE_DIR)
	@go test \
		-run "$(FILTER)" \
		-cover \
		-covermode=atomic \
		-coverprofile=$(COVERAGE_DIR)/unit.out \
		./client/ \
		./internal/auth0/ \
		./internal/telemetry/ \
		./internal/transport/

test-fuzz: ## Run fuzz tests for a fixed duration. Usage: `make test-fuzz FUZZ_TIME=30s`
	@echo "==> Running fuzz tests ($(FUZZ_TIME) each)..."
	@go test -fuzz=FuzzSanitizeDomain              -fuzztime=$(FUZZ_TIME) ./internal/auth0/
	@go test -fuzz=FuzzDeriveURLs                   -fuzztime=$(FUZZ_TIME) ./internal/auth0/
	@go test -fuzz=FuzzValidateOptions              -fuzztime=$(FUZZ_TIME) ./internal/auth0/
	@go test -fuzz=FuzzAudienceFromTokenURL         -fuzztime=$(FUZZ_TIME) ./internal/transport/
	@go test -fuzz=FuzzNewPrivateKeyJwtTokenSource  -fuzztime=$(FUZZ_TIME) ./internal/transport/
	@go test -fuzz=FuzzEncodeAuth0ClientInfo        -fuzztime=$(FUZZ_TIME) ./internal/telemetry/
	@echo "==> Fuzz tests passed."

test-wiremock: ## Run WireMock integration tests. Usage: `make test-wiremock FILTER="TestName"`
	@echo "==> Starting WireMock container..."
	@mkdir -p $(COVERAGE_DIR)
	@docker compose -p "$(WIREMOCK_PROJECT_NAME)" -f $(WIREMOCK_COMPOSE_FILE) up -d
	@WIREMOCK_PORT=$$(docker compose -p "$(WIREMOCK_PROJECT_NAME)" -f $(WIREMOCK_COMPOSE_FILE) port wiremock 8080 | cut -d: -f2); \
		WIREMOCK_URL="http://localhost:$$WIREMOCK_PORT"; \
		echo "==> Waiting for WireMock at $$WIREMOCK_URL..."; \
		until curl -s "$$WIREMOCK_URL/__admin/mappings" > /dev/null 2>&1; do \
			sleep 1; \
		done; \
		echo "==> WireMock is ready at $$WIREMOCK_URL"; \
		echo "==> Running WireMock integration tests..."; \
		WIREMOCK_URL="$$WIREMOCK_URL" go test \
			-run "$(FILTER)" \
			-cover \
			-coverpkg=./... \
			-covermode=atomic \
			-coverprofile=$(COVERAGE_DIR)/wiremock.out \
			./... ; \
		EXIT_CODE=$$?; \
		echo "==> Stopping WireMock container..."; \
		docker compose -p "$(WIREMOCK_PROJECT_NAME)" -f $(WIREMOCK_COMPOSE_FILE) down -v; \
		exit $$EXIT_CODE

coverage: ## Merge coverage profiles and generate report
	@echo "==> Merging coverage profiles..."
	@echo "mode: atomic" > coverage.out
	@for f in $(COVERAGE_DIR)/*.out; do \
		if [ -f "$$f" ]; then \
			tail -n +2 "$$f" >> coverage.out; \
		fi; \
	done
	@echo "==> Coverage report:"
	@go tool cover -func=coverage.out | tail -1
	@rm -rf $(COVERAGE_DIR)

#-----------------------------------------------------------------------------------------------------------------------
# Tidy
#-----------------------------------------------------------------------------------------------------------------------
.PHONY: tidy

tidy: ## Tidy go modules
	@echo "==> Tidying go modules..."
	@go mod tidy
