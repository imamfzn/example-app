export LINTER_VERSION ?= 1.53.3

GO_PACKAGES ?= $(shell go list ./... | grep -v 'examples\|qtest\|mock')
UNAME       := $(shell uname)

MIGRATION_TOOL_EXISTS = 0
ifneq ("$(wildcard $(CUR_DIR)/bin/migrate)","")
    MIGRATION_TOOL_EXISTS = 1
endif

tool-lint: bin
	@test -e ./bin/golangci-lint || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ./bin v${LINTER_VERSION}

lint: tool-lint
	./bin/golangci-lint run -v --timeout 3m0s

bin:
	@mkdir -p bin

tool-migrate: bin
ifeq ($(MIGRATION_TOOL_EXISTS), 1)
	@echo "Migration tool has been existed."
else ifeq ($(UNAME), Linux)
	@curl -sSfL https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar zxf - --directory /tmp \
	&& cp /tmp/migrate bin/
else ifeq ($(UNAME), Darwin)
	@curl -sSfL https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.darwin-amd64.tar.gz | tar zxf - --directory /tmp \
	&& cp /tmp/migrate bin/
else
	@echo "Your OS is not supported."
endif

create-migration: tool-migrate
	bin/migrate create -ext sql -dir migrations/ $(NAME)

migrate: tool-migrate
	bin/migrate -path migrations/ -database "$(DB)" -verbose up

test:
	@go test -race -v ${GO_PACKAGES}

dep:
	@go mod download

tidy:
	@go mod tidy