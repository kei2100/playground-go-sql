GREP := grep
CAT := cat
ifeq ($(OS),Windows_NT)
GREP := findstr
CAT := type
endif

.PHONY: setup
setup: tools db.init db.migrate

.PHONY: tools
tools:
	go get -u golang.org/x/tools/cmd/goimports
	go get -u golang.org/x/lint/golint
	go get -u github.com/rakyll/gotest
	go get github.com/ArcanjoQueiroz/wait-for@b8061231466559274f181d59664051a3b46ad64a
	go get github.com/k0kubun/sqldef/cmd/psqldef@v0.5.6
	go get github.com/volatiletech/sqlboiler@v3.5.0
	go get github.com/volatiletech/sqlboiler/drivers/sqlboiler-psql@v3.5.0
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
# test
#
GO_TEST ?= gotest

.PHONY: test
test: db.start
	$(GO_TEST) -v -race $(PACKAGES)

.PHONY: test.nocache
test.nocache: db.start
	$(GO_TEST) -count=1 -v -race $(PACKAGES)

#
# db
#
DB := postgres
DB_PORT ?= 25432
DB_USER ?= postgres
DB_PASS ?= postgres
DB_NAME ?= develop
DB_ENV := DB_PORT=$(DB_PORT) DB_USER=$(DB_USER) DB_PASS=$(DB_PASS)
DB_DOCKER_COMPOSE := $(DB_ENV) docker-compose
DB_CLI := psql -U $(DB_USER)
DB_WAIT_FOR := wait-for --port=$(DB_PORT) --user=$(DB_USER) --password=$(DB_PASS) --seconds=3 --maxAttempts=10
DB_PGSSLMODE ?= disable

CREATE_DATABASE_CMD := $(DB_CLI) -d postgres -c "CREATE DATABASE $(DB_NAME) ENCODING 'UTF8' LC_CTYPE 'C' LC_COLLATE 'C' TEMPLATE 'template0'"
DROP_DATABASE_CMD := $(DB_CLI) -d postgres -c "DROP DATABASE IF EXISTS $(DB_NAME)"

SQLDEF := PGSSLMODE=$(DB_PGSSLMODE) psqldef --port=$(DB_PORT) --user=$(DB_USER) --password=$(DB_PASS)
SQLDEF_SCHEMA_FILE := ./postgres/schema.sql
SQLDEF_DRYRUN_FILE := ./postgres/tmp/dryrun.sql


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

.PHONY: db.migrate
db.migrate: db.start
	$(SQLDEF) $(DB_NAME) --file=$(SQLDEF_SCHEMA_FILE)

.PHONY: db.migrate.dryrun
db.migrate.dryrun: db.start
	$(SQLDEF) $(DB_NAME) --file=$(SQLDEF_SCHEMA_FILE) --dry-run > $(SQLDEF_DRYRUN_FILE)
	$(CAT) $(SQLDEF_DRYRUN_FILE)

.PHONY: db.console
db.console: db.start
	$(DB_DOCKER_COMPOSE) exec $(DB) $(DB_CLI)

#
# app
#
APP_DB_DRIVER := psql
APP_ENTITY_PKG := entity
APP_ENTITY_DIR := app/infra/db/$(APP_ENTITY_PKG)

.PHONY: app.entity.gen
app.entity.gen: db.start
	sqlboiler $(APP_DB_DRIVER) -p $(APP_ENTITY_PKG) -o $(APP_ENTITY_DIR)
