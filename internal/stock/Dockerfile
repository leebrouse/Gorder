# Dockerfile for stock
FROM golang:1.24.2-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# 拷贝所有代码（包括 common 和 stock）
COPY internal ./internal

# 编译 stock 服务
RUN go build -o stock ./internal/stock

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/stock .

EXPOSE 5003 8284
ENTRYPOINT ["./stock"]

