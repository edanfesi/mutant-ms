make install:
	go install

build:
	go build -o bin/main main.go

run:
	migrate -source file:./migrations -database "postgres://postgres:123456@localhost:5432/postgres?sslmode=disable" up
	go run main.go

migrate:
	migrate -source file:./migrations -database "postgres://postgres:123456@localhost:5432/postgres?sslmode=disable" up

test:
	go test ./... -v