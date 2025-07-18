[![License](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

<!-- ![logo](./docs/images/logo.png) -->

[English](README_en.md) | [简体中文](README_cn.md)

# Gorder

Gorder-v2 is a distributed microservices e-commerce order system, covering core modules such as order, stock, payment, and kitchen. It supports high concurrency, observability, service discovery, message queue, and other modern cloud-native features. Suitable for both learning and production.

## ✨ Features

1. Multiple business modules: decoupled order, stock, payment, and kitchen services
2. Cloud-native support: service discovery (Consul), tracing (Jaeger), monitoring (Prometheus & Grafana)
3. Multiple storage backends: MySQL, MongoDB, Redis
4. High concurrency & reliability: message queue (RabbitMQ), distributed architecture
5. Rich scripts & one-click deployment: Docker Compose, init scripts

## 🏗️ Architecture

![Architecture](.\images\architecture.jpg)

## 🚀 Quick Start

### Prerequisites

- Install Docker and Docker Compose
- Go 1.18+ (for local development)

### Start dependencies

```bash
docker-compose up -d
```

### Initialize database

- MySQL will automatically execute `init.sql`

### Start microservices

```bash
cd internal/order && go run main.go
cd internal/stock && go run main.go
cd internal/payment && go run main.go
cd internal/kitchen && go run main.go
```

### Access services

- Consul UI: http://localhost:8500
- RabbitMQ UI: http://localhost:15672
- Jaeger UI: http://localhost:16686
- Prometheus: http://localhost:9090
- Grafana: http://localhost:3000
- Order HTTP API: http://localhost:8282/api

## 📂 Project Structure

```text
api/                # API definitions (OpenAPI, Protobuf)
internal/           # Core microservices
public/             # Frontend static resources
prometheus/         # Prometheus config
scripts/            # Utility scripts
docker-compose.yml  # One-click start for all dependencies
init.sql            # MySQL initialization script
```

## 📝 License

Gorder is licensed under the [MIT License](LICENSE)
