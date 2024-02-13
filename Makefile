build: 
	cd cmd && go build -o hexsatisfaction .

docker: build
	# docker image prune -af
	docker build -t hexsatisfaction:1.0 -f docker/Dockerfile --progress plain .
	docker build -t hexsatisfaction_postgres:1.0 -f docker/Dockerfile-postgres --progress plain .
	docker build -t pgbouncer:1.0 -f docker/Dockerfile-pgbouncer.yaml --progress plain .
	docker build -t busybox:1.0 -f docker/Dockerfile-busybox --progress plain .

docker-compose: docker
	# docker image prune -af
	# docker container prune -f
	docker-compose build -f docker/docker-compose.yml --no-cache
	docker-compose -f docker/docker-compose.yml up -d

down-compose:
	docker-compose -f docker/docker-compose.yml down

kdown:
	kubectl delete --recursive -f k8s/postgres
	kubectl delete -f k8s/pgbouncer.yaml
	# kubectl delete -f k8s/hexsatisfaction.yaml
	# kubectl apply -f k8s/removeJob.yaml
kup: docker
	kubectl apply --recursive -f k8s/postgres 
	kubectl apply -f k8s/pgbouncer.yaml
	# kubectl apply -f k8s/hexsatisfaction.yaml
	# kubectl delete -f k8s/removeJob.yaml
kr: kdown kup





























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

test-coverage:
	go test ./... -coverprofile coverage.out
	go tool cover -func coverage.out
	go tool cover -html=coverage.out -o coverage.html

start : lint  swagger  run

