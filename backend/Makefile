GOCMD = go
GOBUILD = $(GOCMD) build
GOMOD = $(GOCMD) mod
GOTEST = $(GOCMD) test
BINARY_NAME = main
# Define the timestamp format
TIMESTAMP := $(shell date +%Y%m%d%H%M%S)

# Define the default environment
ENV ?= debug

# Define variables
MIGRATION_DIR=../db/migration
MIGRATE_CMD=migrate create -ext sql -dir $(MIGRATION_DIR) -seq

migration:
	@read -p "Enter migration name: " name; \
	$(MIGRATE_CMD) $$name
migration_up:
	ENV=$(ENV) . ./load_env.sh && migrate -path $(MIGRATION_DIR) -database "postgresql://$$DB_USER:$$DB_PASSWORD@$$DB_HOST:$$DB_PORT/$$DB_NAME?sslmode=disable" -verbose up 1

migration_down:
	ENV=$(ENV) . ./load_env.sh && migrate -path $(MIGRATION_DIR) -database "postgresql://$$DB_USER:$$DB_PASSWORD@$$DB_HOST:$$DB_PORT/$$DB_NAME?sslmode=disable" -verbose down 1

# Usage: make migrate_force version=2
migration_force:
	@if [ -z "$(version)" ]; then \
		echo "Error: You must specify a version number, e.g., 'make migrate_force version=2'"; \
		exit 1; \
	fi
	ENV=$(ENV) . ./load_env.sh && migrate -path $(MIGRATION_DIR) -database "postgresql://$$DB_USER:$$DB_PASSWORD@$$DB_HOST:$$DB_PORT/$$DB_NAME?sslmode=disable" -verbose force $(version)
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
server:
	go run main.go
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o ./$(BINARY_NAME) -v ./
swag:
	swag init
reload:
	sudo systemctl restart iot-backend.service
redis:
	docker run --name redis_7.2 -p 6379:6379  -d redis:7.2-alpine
backup:
	cp main ../backup/backend/main_bk_$(TIMESTAMP)
	pg_dump -U nuodi -h 0.0.0.0 -W -F c -b -c -f ../backup/backend/db_$(TIMESTAMP).backup iot_backend

