FROM golang:latest

WORKDIR /go/src/app
COPY . .

# install psql
RUN apt-get update
RUN apt-get -y install postgresql-client

# make wait-for-postgres.sh executable
RUN chmod +x wait-for-postgres.sh

# build go app
RUN go get -d -v ./...
RUN go install -v ./...

CMD go run main.go
