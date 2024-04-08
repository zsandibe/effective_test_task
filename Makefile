migrate-up:
	migrate -path migrations -database "postgres://postgres:postgres@127.0.0.1:5432/effective?sslmode=disable" -verbose up 

migrate-down:
	migrate -path migrations -database "postgres://postgres:postgres@127.0.0.1:5432/effective?sslmode=disable" -verbose down
