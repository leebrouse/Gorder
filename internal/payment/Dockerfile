# step1: build the go app
FROM golang:1.24.2-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o payment

# step2: copy the binary to a new image
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/payment .

# 5004 grpc / 8285 http
EXPOSE 5004 8285

ENTRYPOINT ["./payment"]