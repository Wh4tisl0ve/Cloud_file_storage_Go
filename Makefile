BINARY_NAME=cloud-file-storage

MAIN_FILE=cmd/cloud-file-storage/main.go
MIGRATION_FILE=cmd/migrate/main.go

build:
	go build -o $(BINARY_NAME) $(MAIN_FILE)

start_app:
	go run $(MAIN_FILE)

migration_up:
	go run $(MIGRATION_FILE) up

migration_down:
	go run $(MIGRATION_FILE) down

# test, lint, docker compose up
