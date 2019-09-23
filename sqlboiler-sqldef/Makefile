GREP := grep
CAT := cat
ifeq ($(OS),Windows_NT)
GREP := findstr
CAT := type
endif

.PHONY: setup
setup:
	go get -u golang.org/x/tools/cmd/goimports
	go get -u golang.org/x/lint/golint
	go get -u github.com/jwilder/dockerize
	go get github.com/ArcanjoQueiroz/wait-for@b8061231466559274f181d59664051a3b46ad64a
	go get github.com/k0kubun/sqldef/cmd/psqldef@v0.5.3
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
DB_ENV := DB_PORT=$(DB_PORT) DB_USER=$(DB_USER) DB_PASS=$(DB_PASS)
DB_DOCKER_COMPOSE := $(DB_ENV) docker-compose
DB_CLI := psql -U $(DB_USER)
DB_WAIT_FOR := wait-for --port=$(DB_PORT) --user=$(DB_USER) --password=$(DB_PASS) --seconds=3 --maxAttempts=10
ifndef DB_PGSQLMODE
DB_PGSQLMODE := disable
endif
SQLDEF := PGSSLMODE=$(DB_PGSQLMODE) psqldef --port=$(DB_PORT) --user=$(DB_USER) --password=$(DB_PASS)
SQLDEF_SCHEMA_FILE := ./postgres/schema.sql
SQLDEF_DRYRUN_FILE := ./postgres/tmp/dryrun.sql

CREATE_DATABASE_CMD := $(DB_CLI) -d postgres -c "CREATE DATABASE $(DB_NAME) ENCODING 'UTF8' LC_CTYPE 'C' LC_COLLATE 'C' TEMPLATE 'template0'"
DROP_DATABASE_CMD := $(DB_CLI) -d postgres -c "DROP DATABASE IF EXISTS $(DB_NAME)"


.PHONY: db.init
db.init: db.start
	$(DB_DOCKER_COMPOSE) exec $(DB) $(DROP_DATABASE_CMD)
	$(DB_DOCKER_COMPOSE) exec $(DB) $(CREATE_DATABASE_CMD)

.PHONY: db.start
db.start:
	$(DB_DOCKER_COMPOSE) ps --filter "status=running" --service | $(GREP) $(DB) || $(DB_DOCKER_COMPOSE) up -d $(DB)
	$(DB_WAIT_FOR)

.PHONY: db.stop
db.stop:
	$(DB_DOCKER_COMPOSE) kill $(DB)

.PHONY: db.restart
db.restart: db.stop db.start

.PHONY: db.schema.migrate
db.schema.migrate: db.start
	$(SQLDEF) $(DB_NAME) --file=$(SQLDEF_SCHEMA_FILE)

.PHONY: db.schema.migrate.dryrun
db.schema.migrate.dryrun: db.start
	$(SQLDEF) $(DB_NAME) --file=$(SQLDEF_SCHEMA_FILE) --dry-run > $(SQLDEF_DRYRUN_FILE)
	$(CAT) $(SQLDEF_DRYRUN_FILE)

.PHONY: db.cli
db.cli: db.start
	$(DB_DOCKER_COMPOSE) exec $(DB) $(DB_CLI)
