# syntax=docker/dockerfile:1

# step1: build the go app
FROM golang:1.24.2-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o stock

# step2: copy the binary to a new image
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/stock .

# 5003 grpc / 8284 http
EXPOSE 5003 8284

ENTRYPOINT ["./stock"]
