# Insider Assessment Project - Messaging

## 📌 Overview
The **Insider Assessment Project - Messaging** is a backend system designed to handle message processing efficiently using **Go, Fiber, MongoDB, and Redis**. The system consists of two primary services:

1. **API Service**: Exposes endpoints for managing messages and controlling the worker.
2. **Worker Service**: Processes queued messages and ensures no duplicate processing using **Redis locking**.

## 🚀 Features
- **REST API** using Fiber.
- **Database** is MongoDB.
- **Redis-based locking** to prevent duplicate message processing.
- **Worker control via Redis Pub/Sub** (start/stop commands).
- **Docker Compose support**
- **Swagger API documentation**.

---

## 🏗️ Project Structure
```
insider-messaging/
│── cmd/
│   ├── api/           # API service entry point
│   ├── worker/        # Worker service entry point
│── internal/
│   ├── cache/         # Redis cache
│   ├── database/      # Database configuration
│   ├── message/       # Business logic (service, repository, models)
│   ├── test/          # Unit & integration tests
│── configs/           # Centralized configuration  
│── api/
│   ├── handler/       # API request handlers
│   ├── router/        # API routes
│── worker/            # Worker service
│── docs/              # Swagger documentation
│── scripts/           # Scripts init database
│── docker-compose.yml # Docker Compose configuration
│── README.md
```

---

## ⚙️ Installation & Setup
### **1️⃣ Clone the Repository**
```sh
git clone https://github.com/brkdmn/insider-assessment.git
cd insider-assessment
```

### **2️⃣ Setup Environment**
Configure **MongoDB** and **Redis** connection details in `configs/config.yaml`.

### **3️⃣ Run with Docker Compose**
```sh
docker-compose up -d --build
```
This will start:
- API on **http://localhost:8080**
- MongoDB on **mongodb://localhost:27017**
- Redis on **redis://localhost:6379**
- Worker service in the background

---

## 📝 API Documentation (Swagger)
After starting the API, Swagger documentation will be available at:
📌 **[http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)**

Example Endpoints:
- `GET /messages` → Retrieve sending messages.
- `POST /worker/start` → Start the worker.
- `POST /worker/stop` → Stop the worker.

---

## 🔧 Worker Process & Locking Mechanism
The Worker service fetches pending messages from MongoDB every **2 minutes**, processing up to **3 messages per cycle**. To prevent duplicate processing:
- **Redis Locking (`SETNX`)** ensures only one worker processes a message.
- **Redis Pub/Sub** controls worker execution remotely (`start/stop`).

### **Start/Stop Worker via API**
```sh
curl -X POST http://localhost:8080/worker/start
curl -X POST http://localhost:8080/worker/stop
```

---


