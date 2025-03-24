run:
	go run main.go

test:
	go test -v ./...

mod-vendor:
	go mod vendor

golangci-lint:
	@golangci-lint run

gosec:
	@gosec -quiet ./...

validate: golangci-lint gosec

docker:
	docker-compose build
	docker-compose up

migrate-create:
	@goose -dir=migrations create "$(name)" sql

migrate-up:
	@goose -dir=migrations postgres "host=${POSTGRES_HOST} user=${POSTGRES_USER} password=${POSTGRES_PASSWORD} dbname=${POSTGRES_DB} sslmode=disable" up

migrate-down:
	@goose -dir=migrations postgres "host=${POSTGRES_HOST} user=${POSTGRES_USER} password=${POSTGRES_PASSWORD} dbname=${POSTGRES_DB} sslmode=disable" down

install-swag-deps:
	go get -u github.com/swaggo/swag/cmd/swag

swagger:
	swag fmt
	swag init -g cmd/api/main.go -o ./docs

docs: swagger
	redocly build-docs docs/swagger.json -o docs/index.html