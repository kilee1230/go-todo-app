APP_NAME = go-todo-api
BUILD_DIR = $(PWD)/build

clean:
	rm -rf ./build

run-dev:
	air -c .air.toml

build: clean
	go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) main.go

tidy:
	go mod tidy

lint:
	golangci-lint run ./...