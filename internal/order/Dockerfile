# step1: build the go app
FROM golang:1.24.2-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o order

# step2: copy the binary to a new image
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/order .

# 5002 grpc / 8283 http
EXPOSE 5002 8283

ENTRYPOINT ["./order"]
