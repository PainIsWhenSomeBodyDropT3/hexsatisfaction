.PHONY lint:
		golangci-lint run --config .golangci.yml

migrate_up:
	    migrate -path migrations/ -database "postgres://postgres:18051965q@localhost:5432/postgres?sslmode=disable" up

migrate_down:
	    migrate -path migrations/ -database "postgresql://postgres:18051965q@localhost:5432/postgres?sslmode=disable" down
