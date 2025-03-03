run:
	go run cmd/api-server/main.go

tidy:
	go mod tidy

compose:
	docker compose up -d

compose-build:
	docker compose up -d --build

compose-down:
	docker compose down

# make migrate-create NAME=create_users_table
migrate-create:
	docker run --rm --network host -v ./db:/db migrate/migrate \
		create -ext sql -dir db/migrations ${NAME}

migrate-local:
	docker run --rm --network host -v ./db:/db migrate/migrate \
		-path db/migrations -database postgres://appuser:1234@localhost:5432/testdb?sslmode=disable up
