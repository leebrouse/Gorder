.PHONY: gen
gen: genproto genopenapi

#Create gRPC code such as Service interface and Service struct
.PHONY: genproto
genproto:
	@./scripts/genproto.sh

#Create openapi code
.PHONY: genopenapi
genopenapi:
	@./scripts/genopenapi.sh


.PHONY: fmt
fmt:
	goimports -l -w internal/

.PHONY: lint
lint:
	@./scripts/lint.sh

webhook:
	stripe listen --forward-to localhost:8285/api/webhook


# Docker deploy automaticlly:
ORDER_BINARY=orderApp
STOCK_BINARY=stockApp
PAYMENT_BINARY=paymentApp
# kitchen_BINARY

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up:
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

build_all: build_order build_stock build_payment
	@echo "All Service were built."

## build_broker: builds the broker binary as a linux executable
build_order:
	@echo "Building order binary..."
	cd ./internal/order && env GOOS=linux CGO_ENABLED=0 go build -o ${ORDER_BINARY} .
	@echo "Done!"

build_stock:
	@echo "Building stock binary..."
	cd ./internal/stock && env GOOS=linux CGO_ENABLED=0 go build -o ${STOCK_BINARY} .
	@echo "Done!"

build_payment:
	@echo "Building payment binary..."
	cd ./internal/payment && env GOOS=linux CGO_ENABLED=0 go build -o ${PAYMENT_BINARY} .
	@echo "Done!"



