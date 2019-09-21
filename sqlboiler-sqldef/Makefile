GREP := grep
ifeq ($(OS),Windows_NT)
	GREP := findstr
endif

.PHONY: setup
setup:
	go get -u golang.org/x/tools/cmd/goimports
	go get -u golang.org/x/lint/golint
	go get -u github.com/jwilder/dockerize
	go get -u github.com/kei2100/wait-for@use-password
	go get github.com/k0kubun/sqldef/cmd/psqldef@v0.5.2
	go mod tidy

#
# lint
#
PACKAGES := $(shell go list ./...)

.PHONY: fmt
fmt:
	goimports -w .

.PHONY: lint
lint:
	goimports -d . | $(GREP) "^" && exit 1 || exit 0
	golint -set_exit_status $(PACKAGES)

.PHONY: vet
vet:
	go vet $(PACKAGES)

#
# db
#
DB := postgres
DB_PORT := 25432
DB_USER := postgres
DB_PASS := postgres
DB_NAME := develop
TEST_DB_NAME := test
DB_ENV := DB_PORT=$(DB_PORT) DB_USER=$(DB_USER) DB_PASS=$(DB_PASS)
DOCKER_COMPOSE_DB := $(DB_ENV) docker-compose
DB_CLI := psql -U $(DB_USER)

CREATE_DATABASE_OPTS := ENCODING 'UTF8' LC_CTYPE 'C' LC_COLLATE 'C' TEMPLATE 'template0'
CREATE_DATABASE_CMD := $(DB_CLI) -d postgres -c "CREATE DATABASE $(DB_NAME) $(CREATE_DATABASE_OPTS)"
CREATE_TEST_DATABASE_CMD := $(DB_CLI) -d postgres -c "CREATE DATABASE $(TEST_DB_NAME) $(CREATE_DATABASE_OPTS)"
DROP_DATABASE_CMD := $(DB_CLI) -d postgres -c "DROP DATABASE IF EXISTS $(DB_NAME)"
DROP_TEST_DATABASE_CMD := $(DB_CLI) -d postgres -c "DROP DATABASE IF EXISTS $(TEST_DB_NAME)"

DB_WAIT_FOR_OPTS := --port=$(DB_PORT) --user=$(DB_USER) --password=$(DB_PASS) --seconds=3 --maxAttempts=10

.PHONY: db.init
db.init: db.start
	$(DOCKER_COMPOSE_DB) exec $(DB) $(DROP_DATABASE_CMD)
	$(DOCKER_COMPOSE_DB) exec $(DB) $(CREATE_DATABASE_CMD)
	$(MAKE) db.test.init

.PHONY: db.test.init
db.test.init: db.start
	$(DOCKER_COMPOSE_DB) exec $(DB) $(DROP_TEST_DATABASE_CMD)
	$(DOCKER_COMPOSE_DB) exec $(DB) $(CREATE_TEST_DATABASE_CMD)

.PHONY: db.start
db.start:
	$(DOCKER_COMPOSE_DB) ps --filter "status=running" --service | $(GREP) $(DB) || $(DOCKER_COMPOSE_DB) up -d $(DB)
	wait-for $(DB_WAIT_FOR_OPTS)

.PHONY: db.stop
db.stop:
	$(DOCKER_COMPOSE_DB) kill $(DB)

.PHONY: db.restart
db.restart: db.stop db.start

.PHONY: db.migrate
db.migrate: db.start
	# TODO
	echo "setup"

.PHONY: db.cli
db.cli: db.start
	$(DOCKER_COMPOSE_DB) exec $(DB) $(DB_CLI)
