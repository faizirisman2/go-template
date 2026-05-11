GOLANGCI_CMD := $(shell command -v golangci-lint 2> /dev/null)
GOLANGCI_VERSION := $(shell $(GOLANGCI_CMD) version 2>/dev/null | grep -oE '[0-9]+\.[0-9]+\.[0-9]+' | head -n1)
GOLANGCI_MAJOR := $(shell echo $(GOLANGCI_VERSION) | cut -d. -f1)

check-golangci:
ifndef GOLANGCI_CMD
	$(error "Please install golangci linters from https://golangci-lint.run/usage/install/")
endif

ifeq ($(GOLANGCI_MAJOR),)
	$(error "Could not determine golangci-lint version")
else ifeq ($(GOLANGCI_MAJOR),1)
	$(error "Please update to golangci-lint version 2")
else
	@echo "Linting using : golangci-lint $(GOLANGCI_MAJOR)"
endif

lint: check-golangci
	@echo -e "$(OK_COLOR)==> linting projects$(NO_COLOR)..."
	@golangci-lint run --fix -v -c golangci.yaml
	@echo -e "$(OK_COLOR)==> done, all ok$(NO_COLOR)..."

deps:
	@go mod tidy && go mod vendor

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