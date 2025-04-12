# Etapa de build
FROM golang:1.24.2-alpine3.21 AS builder

WORKDIR /build

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o app

FROM alpine:3

WORKDIR /app

COPY --from=builder /build/app .

EXPOSE 8080

RUN echo "DATABASE_URL=postgres://postgres:root@db:5432/postgres?sslmode=disable" > .env

CMD ["./app"]
