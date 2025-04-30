FROM leebrouse7/common-base as common

FROM golang:1.24.2-alpine as builder

WORKDIR /app
COPY go.mod go.sum /app/order/
RUN go mod download

COPY --from=common /app/common /app/common
COPY . /app/order

RUN go build -o order /app/order

# step2: copy the binary to a new image
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/order .

# 暴露服务端口
EXPOSE 5002 8283

ENTRYPOINT ["./order"]

