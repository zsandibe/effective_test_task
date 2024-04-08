migrate-up:
	migrate -path migrations -database "postgres://postgres:postgres@127.0.0.1:5432/effective?sslmode=disable" -verbose up 

migrate-down:
	migrate -path migrations -database "postgres://postgres:postgres@127.0.0.1:5432/effective?sslmode=disable" -verbose down


.PHONY: docker-up
docker-up:
	docker compose up

.PHONY: docker-down
docker-down:
	docker compose down


run:
	go run cmd/main.go