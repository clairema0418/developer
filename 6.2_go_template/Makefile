.PHONY: build run migrate

build:
	go build -o app main.go

run: build
	./app

migrate:
	go run config/migration.go

generate_crud:
	./templates/crud_api_template.sh ${TABLE_NAME}