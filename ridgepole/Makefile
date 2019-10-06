.PHONY: setup
setup: tools db.init db.migrate

.PHONY: tools
tools:
	go get github.com/ArcanjoQueiroz/wait-for@b8061231466559274f181d59664051a3b46ad64a
	go mod tidy

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

DB_URL := postgres://$(DB_USER):$(DB_PASS)@$(DB):5432/$(DB_NAME)


.PHONY: db.init
db.init: db.start
	$(DB_DOCKER_COMPOSE) exec $(DB) $(DROP_DATABASE_CMD)
	$(DB_DOCKER_COMPOSE) exec $(DB) $(CREATE_DATABASE_CMD)

.PHONY: db.start
db.start:
	$(DB_DOCKER_COMPOSE) ps --filter "status=running" --service | grep $(DB) || $(DB_DOCKER_COMPOSE) up -d $(DB)
	$(DB_WAIT_FOR)

.PHONY: db.stop
db.stop:
	$(DB_DOCKER_COMPOSE) kill $(DB)

.PHONY: db.restart
db.restart: db.stop db.start

.PHONY: db.schema.export
db.schema.export: db.start
	$(IGNORE)docker-compose run --rm ridgepole -c $(DB_URL) --export -o Schemafile

.PHONY: db.schema.migrate.dryrun
db.schema.migrate.dryrun: db.start
	docker-compose run --rm ridgepole -c $(DB_URL) --apply --dry-run

.PHONY: db.schema.migrate
db.schema.migrate: db.start
	docker-compose run --rm ridgepole -c $(DB_URL) --apply

.PHONY: db.console
db.console: db.start
	$(DB_DOCKER_COMPOSE) exec $(DB) $(DB_CLI)
