.PHONY lint:
		golangci-lint run --config .golangci.yml

migrate_up:
	    migrate -path migrations/ -database "postgres://postgres:18051965q@localhost:5432/postgres?sslmode=disable" up

migrate_down:
	    migrate -path migrations/ -database "postgresql://postgres:18051965q@localhost:5432/postgres?sslmode=disable" down

.PHONY swagger:swagger-spec | swagger-validate

swagger-spec:
	SWAGGER_GENERATE_EXTENSION=false swagger -q generate spec -m -o ./swagger.yaml

swagger-validate:
	swagger -q validate ./swagger.yaml

run:
	go run cmd/main.go

start : lint  swagger  run
