FROM golang:1.16 as build

WORKDIR /app

COPY . .

WORKDIR /app/cmd
RUN CGO_ENABLED=0 go build -o hexsatisfaction .

FROM alpine:3 

WORKDIR /app/cmd

COPY --from=build /app/cmd/hexsatisfaction .

CMD ["./hexsatisfaction"]


