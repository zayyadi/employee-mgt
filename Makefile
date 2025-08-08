# Variables
BINARY=employee-management
MAIN_FILE=cmd/server/main.go
DOCKER_COMPOSE=docker-compose.yml

# Default target
.PHONY: all
all: build

# Build the application
.PHONY: build
build:
	go build -o ${BINARY} ${MAIN_FILE}

# Run the application
.PHONY: run
run:
	go run ${MAIN_FILE}

# Install dependencies
.PHONY: deps
deps:
	go mod tidy

# Run tests
.PHONY: test
test:
	go test -v ./...

# Run tests with coverage
.PHONY: test-coverage
test-coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

# Clean build files
.PHONY: clean
clean:
	rm -f ${BINARY}

# Format code
.PHONY: fmt
fmt:
	go fmt ./...

# Vet code
.PHONY: vet
vet:
	go vet ./...

# Run linter
.PHONY: lint
lint:
	golint ./...

# Start docker containers
.PHONY: docker-up
docker-up:
	docker-compose -f ${DOCKER_COMPOSE} up -d

# Stop docker containers
.PHONY: docker-down
docker-down:
	docker-compose -f ${DOCKER_COMPOSE} down

# Build and start docker containers
.PHONY: docker-build
docker-build:
	docker-compose -f ${DOCKER_COMPOSE} up -d --build

# View docker logs
.PHONY: docker-logs
docker-logs:
	docker-compose -f ${DOCKER_COMPOSE} logs -f

# Run database migrations
.PHONY: migrate-up
migrate-up:
	go run cmd/migrate/main.go up

# Run database rollback
.PHONY: migrate-down
migrate-down:
	go run cmd/migrate/main.go down

# Generate Swagger documentation
.PHONY: swagger
swagger:
	swag init -g ${MAIN_FILE} --output ./docs/swagger

# Help
.PHONY: help
help:
	@echo "Available commands:"
	@echo "  make build        - Build the application"
	@echo "  make run          - Run the application"
	@echo "  make deps         - Install dependencies"
	@echo "  make test         - Run tests"
	@echo "  make test-coverage- Run tests with coverage"
	@echo "  make clean        - Clean build files"
	@echo "  make fmt          - Format code"
	@echo "  make vet          - Vet code"
	@echo "  make lint         - Run linter"
	@echo "  make docker-up    - Start docker containers"
	@echo "  make docker-down  - Stop docker containers"
	@echo "  make docker-build - Build and start docker containers"
	@echo "  make docker-logs  - View docker logs"
	@echo "  make migrate-up   - Run database migrations"
	@echo "  make migrate-down - Run database rollback"
	@echo "  make swagger      - Generate Swagger documentation"
	@echo "  make help         - Show this help message"
