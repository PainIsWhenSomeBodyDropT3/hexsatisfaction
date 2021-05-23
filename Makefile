.PHONY lint:
		golangci-lint run --config .golangci.yml

migrate_up:
	    migrate -path pkg/migrations/ -database "postgres://postgres:18051965q@localhost:5432/postgres?sslmode=disable" up

migrate_down:
	    migrate -path pkg/migrations/ -database "postgresql://postgres:18051965q@localhost:5432/postgres?sslmode=disable" down

.PHONY swagger:swagger-spec

swagger-spec:
	swag init -g cmd/main.go

gen-mocks:
	mockery --all --keeptree
run:
	go run cmd/main.go

start : lint  swagger  run
