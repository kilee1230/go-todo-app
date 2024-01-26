APP_NAME = go-todo-api
BUILD_DIR = $(PWD)/build

.PHONY: run-dev
run-dev:
	air main.go -b 0.0.0.0

.PHONY: build
build: clean
	go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) main.go

.PHONY: clean
clean:
	rm -rf ./build

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: lint
lint:
	golangci-lint run ./...