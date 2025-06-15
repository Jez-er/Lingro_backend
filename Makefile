# Имя проекта (можно изменить под твоё имя модуля)
PROJECT_NAME := Astral-Back-End

# Пути в проекте
CMD_DIR := cmd/app
SRC_DIR := .
DOCS_DIR := $(SRC_DIR)/docs

# Go параметры
GO := go
GOFLAGS := -mod=mod
GOPATH := $(shell go env GOPATH)
SWAG := $(GOPATH)/bin/swag

# Цели по умолчанию
.PHONY: all
all: deps swagger build

# Установка зависимостей
.PHONY: deps
deps:
	$(GO) mod tidy
	$(GO) install github.com/swaggo/swag/cmd/swag@latest

# Генерация Swagger документации
.PHONY: swagger
swagger:
	$(SWAG) init -g $(SRC_DIR)/internal/app/main.go -d ./

# Сборка проекта
.PHONY: build
build: swagger
	$(GO) build $(GOFLAGS) -o $(PROJECT_NAME) $(CMD_DIR)/main.go

# Запуск проекта
.PHONY: run
run: swagger
	$(GO) run $(GOFLAGS) $(CMD_DIR)/main.go

# Тестирование
.PHONY: test
test:
	$(GO) test $(GOFLAGS) ./...

# Очистка сгенерированных файлов
.PHONY: clean
clean:
	del /Q /F $(DOCS_DIR)
	del /Q /F $(PROJECT_NAME)

# Проверка, установлен ли swag
.PHONY: check-swag
check-swag:
	@if [ ! -f $(SWAG) ]; then \
		echo "Swag не установлен. Устанавливаем..."; \
		$(GO) install github.com/swaggo/swag/cmd/swag@latest; \
	fi

# Помощь
.PHONY: help
help:
	@echo "All commands:"
	@echo "  all      - It will install dependencies, generate Swagger and build the project"
	@echo "  deps     - Install project dependencies"
	@echo "  swagger  - Generate Swagger documentation"
	@echo "  build    - Build the project"
	@echo "  run      - Run the project"
	@echo "  test     - Run tests"
	@echo "  clean    - Clean the generated files"
	@echo "  help     - Show this help"

# По умолчанию проверяем, установлен ли swag перед выполнением
$(SWAG):
	$(GO) install github.com/swaggo/swag/cmd/swag@latest
