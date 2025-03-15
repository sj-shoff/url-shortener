FROM golang:1.23-alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux

WORKDIR /api

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -ldflags="-w -s" -o url-shortener ./cmd/url-shortener/main.go

FROM alpine:3.19

RUN apk add --no-cache \
    postgresql15-client \
    tzdata

WORKDIR /api
COPY --from=builder /api/config ./config
COPY --from=builder /api/url-shortener .
COPY --from=builder /api/wait-for-postgres.sh .
COPY --from=builder /api/.env . 

RUN chmod +x wait-for-postgres.sh

CMD ["./wait-for-postgres.sh", "db", "5432", "./url-shortener"]