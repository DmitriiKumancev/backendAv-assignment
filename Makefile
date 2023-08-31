MIGRATIONS_DIR = ./migrations
DATABASE_URL = postgres://postgres:mypassword@localhost:5432/avitoSegmentsDb?sslmode=disable
APP_NAME = backend-trainee-assignment-app

build:
	docker-compose build $(APP_NAME)

run:
	docker-compose up $(APP_NAME)

test:
	go test -v ./...

migrate-new:
	migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq init

migrate-up:
	migrate -path $(MIGRATIONS_DIR) -database $(DATABASE_URL) up

migrate-down:
	migrate -path $(MIGRATIONS_DIR) -database $(DATABASE_URL) down