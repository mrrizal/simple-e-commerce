# E-Commerce API
Example golang app that following clean architecture.

## Run database
`docker-compose up`

## Migrate up
`migrate -path db/migrations -database "postgresql://postgres:postgres@localhost/{DB_NAME}?sslmode=disable" up`

## Migrate down
`migrate -path db/migrations -database "postgresql://postgres:postgres@localhost/{DB_NAME}?sslmode=disable" down`

## Run
`go run main.go`


todo:
- write unittest
- use interfaces/abstract for database
- dockerize golang app
