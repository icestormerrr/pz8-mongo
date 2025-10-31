install:
	go mod tidy

run:
	go run ./cmd/api

build:
	go build -o pz8-mongo.exe ./cmd/api

test:
	go test ./... -v
