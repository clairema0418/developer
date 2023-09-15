.PHONY: build run clean

build:
	go build -o audio_processing main.go

run: build
	./audio_processing

clean:
	rm -f audio_processing

docker-build:
	docker-compose build

docker-up:
	docker-compose up

docker-down:
	docker-compose down

docker-clean: docker-down
	docker-compose rm -f
	docker volume prune -f
	docker system prune -f

test:
	go test -v ./...

lint:
	golangci-lint run

all: lint test build