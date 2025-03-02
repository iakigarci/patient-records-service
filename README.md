# Patient Records Service

A Go-based microservice for managing patient records and medical diagnoses, built using Domain-Driven Design (DDD) principles and a clean architecture approach.

## ⚡ Features

- JWT-based authentication and authorization
- Comprehensive patient management (CRUD operations)
- Medical diagnostic records with history
- Interactive Swagger/OpenAPI documentation
- PostgreSQL database with versioned migrations
- Clean architecture and DDD implementation
- Robust error handling and logging
- Unit and integration tests
- Rate limiting and security middleware

## 🛠 Tech Stack

- Go 1.23.4
- Gin Web Framework
- PostgreSQL
- Goose (database migrations)
- Swagger/OpenAPI
- Docker & Docker Compose
- JWT Authentication
- Zap Logger
- Testify (testing)
- Viper (configuration)

## 📋 Prerequisites

- Docker and Docker Compose
- Go 1.23.4 or higher
- Make
- Git

## 🏗 Project Structure 
```bash
.
├── cmd/
│ └── api/ # Application entry point
├── internal/
│ ├── adapters/ # Adapters layer (REST, DB)
│ ├── domain/ # Domain layer (business logic)
│ └── ports/ # Interfaces
├── docs/ # Swagger documentation
├── .env # Environment variables
```


## ⚐ Getting Started

1. Clone the repository:
```bash
git clone https://github.com/iakigarci/patient-records-service.git
cd patient-records-service
```

2. Start the services:
```bash
docker-compose up -d
```

3. Run database migrations (if not already run):
```bash
make migrate-up
```

4. Check migration status:
```bash
make migrate-status
```

5. Check container status:
```bash
docker ps
```


## 🛣️ API Documentation

Once the service is running, you can access the Swagger documentation at:
http://localhost:8080/v1/swagger/index.html

### Available Endpoints

- **POST** `/v1/auth/login` - User authentication
- **GET** `/v1/diagnostics` - List diagnostics (protected)
- **POST** `/v1/diagnostics` - Create diagnostic (protected)

## 🛠️ Testing

There are unit tests for the diagnostic service.
```bash
cd internal/domain/services/diagnostic
```