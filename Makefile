run:
	go run ./cmd/app

test:
	go test ./...

build:
	go build -o app ./cmd/app

docker-build:
	docker build -t task-api .

up:
	docker compose up --build