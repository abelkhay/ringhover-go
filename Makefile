# --- Config (change if needed) ---
HOST    ?= host.docker.internal
PORT    ?= 4000
DB_NAME ?= tasking
SCHEMA  ?= db/migrations/001_init.sql

# --- Quick help ---
.PHONY: help
help:
	@echo "make db        - start TiDB (MySQL-compatible) in background"
	@echo "make stop-db   - stop/remove TiDB container"
	@echo "make migrate   - apply initial schema ($(SCHEMA))"
	@echo "make verify    - show databases/tables"
	@echo "make deps      - install Go deps"
	@echo "make run       - start API (go run ./cmd/api)"
	@echo "make dsn       - print DSN for .env"

# --- Database ---
.PHONY: db
db:
	docker run -d --name tidb -p $(PORT):4000 pingcap/tidb:v8.5.2

.PHONY: stop-db
stop-db:
	- docker stop tidb
	- docker rm tidb

.PHONY: migrate
migrate:
	cat $(SCHEMA) | docker run --rm -i mysql:8 \
	  mysql -h $(HOST) -P $(PORT) -u root --protocol=tcp

.PHONY: verify
verify:
	docker run --rm -it mysql:8 \
	  mysql -h $(HOST) -P $(PORT) -u root --protocol=tcp \
	  -e "SHOW DATABASES; USE $(DB_NAME); SHOW TABLES;"

# --- Go ---
.PHONY: deps
deps:
	go get github.com/gin-gonic/gin
	go get go.uber.org/zap
	go get github.com/jmoiron/sqlx
	go get github.com/go-sql-driver/mysql
	go get github.com/joho/godotenv

.PHONY: run
run:
	go run ./cmd/api

.PHONY: dsn
dsn:
	@echo "DB_DSN=root:@tcp(127.0.0.1:$(PORT))/$(DB_NAME)?charset=utf8mb4&parseTime=True&loc=Local"
