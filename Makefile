include .env.local
export

APP_NAME = server
BUILD_DIR = $(PWD)/build

clean:
	rm -rf ./build

lint:
	golangci-lint run ./...

build:
	CGO_enabled=0 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) main.go

run:
	$(BUILD_DIR)/$(APP_NAME)

docker-build:
	docker build -t go-rest-test:latest .

docker-run:
	docker-compose up

docker-stop:
	docker-compose down