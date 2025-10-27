# Go-Kafka-Example

🔥 A simple and practical example of using **Kafka with Golang** using the official Go Kafka package.

---

## 📌 Project Overview

This repository is designed for Go developers who want to:

- Work with Kafka using the official Go package ([kafka-go](https://github.com/segmentio/kafka-go))  
- Build event-driven services  
- Produce and consume messages efficiently and reliably  

Using Golang ensures **high performance** and **low memory usage** for message processing.

---

## ⚡ Features

- Simple Producer and Consumer implemented in Go  
- Uses the official [kafka-go](https://github.com/segmentio/kafka-go) package  
- Easily extendable for microservices  
- Ideal for learning and prototyping Kafka-based applications

---

## 🚀 Quick Start

### Prerequisites

- Go >= 1.20  
- Docker & Docker Compose (to run Kafka and Zookeeper)

### Installation & Run

1. Clone the repository:

```bash

# Clone repo
git clone git@github.com:rsulkln/Go-Kafka-Example.git
cd Go-Kafka-Example

# Start Kafka and Zookeeper
docker-compose up -d

# Create and list topics
docker exec -it kafka /bin/bash
docker-topics --create --bootstrap-server localhost:9092 --topic orders --partitions 2
docker-topics --list --bootstrap-server localhost:9092
# Expected output:
# __consumer_offsets
# orders

# Run Producer and Consumers
go run consumer/main.go
go run consumer2/main.go
go run producer/main.go
