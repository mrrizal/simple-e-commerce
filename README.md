# E-Commerce API
Example golang app that following clean architecture.

If you take a look the source code, i'm spliting between the controller and fiber handler,

just in case you want to change the implemention for example using other framework or other protocol like grpc,

its can easly replaced

## Run database
`docker-compose up`

## Migrate up
`migrate -path db/migrations -database "postgresql://postgres:postgres@localhost/{DB_NAME}?sslmode=disable" up`

## Migrate down
`migrate -path db/migrations -database "postgresql://postgres:postgres@localhost/{DB_NAME}?sslmode=disable" down`

## Run
`go run main.go`

## test with coverage
`go clean -testcache && go test -v -coverprofile=coverage.out ./...`

## see coverage in html format
`go tool cover -html=coverage.out`
