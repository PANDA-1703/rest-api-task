# Сборка
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

# Сборка бинарника
RUN CGO_ENABLED=0 go build -o /api main.go

FROM alpine:3.21
WORKDIR /app

RUN apk add --no-cache tzdata postgresql-client

COPY --from=builder /api /api

COPY migrations/ /migrations/

CMD ["sh", "-c", "migrate -source file://migrations -database \"postgres://${DB_USER}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}?sslmode=disable\" up && ./api"]