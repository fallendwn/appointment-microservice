# Medical Appointment Microservices

A Go-based microservices system for managing doctors and medical appointments, using MongoDB.

## Architecture

The project consists of two independent microservices that communicate over HTTP:

```
┌─────────────────────┐        HTTP        ┌──────────────────────┐
│  appointment-service │ ────────────────► │    doctor-service     │
│      :8082           │  (validates doctor)│       :8081           │
└──────────┬──────────┘                    └──────────┬────────────┘
           │                                           │
           └──────────────────┬────────────────────────┘
                              ▼
                        ┌──────────┐
                        │  MongoDB │
                        │  :27017  │
                        └──────────┘
```

- **doctor-service** — doctor managements (CRUD)
- **appointment-service** — appointments management

## Tech Stack

- **Language:** Go
- **Framework:** Gin
- **Database:** MongoDB
- **Containerization:** Docker + Docker Compose

## Project Structure

```
microservices/
├── docker-compose.yaml
├── doctor-service/
│   ├── cmd/main.go
│   ├── Dockerfile
│   ├── go.mod
│   └── internal/
│       ├── app/          # cfg + app initialization
│       ├── model/        # Doctor model
│       ├── repository/   # MongoDB repository + crud
│       ├── transport/    # handlers + http
│       └── usecase/      # buisness-logic
└── appoitment-service/
    ├── cmd/main.go
    ├── Dockerfile
    ├── go.mod
    └── internal/
        ├── app/         
        ├── client/       
        ├── dto/          # DTO
        ├── model/        # Appointment + Status
        ├── repository/   # 
        ├── transport/    # 
        └── usecase/      # 
```

## Getting Started

### Prerequisites

- Docker
- Docker Compose

### Run with Docker Compose

```bash
cd microservices
docker-compose up --build
```



| Service             | URL                        |
|---------------------|----------------------------|
| doctor-service      | http://localhost:8081       |
| appointment-service | http://localhost:8082       |
| MongoDB             | mongodb://localhost:27017   |



`

## API Reference

### Doctor Service — `http://localhost:8081`

| Method | Endpoint              | Description          |
|--------|-----------------------|----------------------|
| POST   | `/api/v1/doctors/`    | Create doctor        |
| GET    | `/api/v1/doctors/`    | List of all doctors  |
| GET    | `/api/v1/doctors/:id` | Get doctor by ID     |

**Doctor model:**
```json
{
  "id": "ObjectID",
  "name": "string",
  "specialization": "string",
  "email": "string"
}
```

---

### Appointment Service — `http://localhost:8082`

| Method | Endpoint                         | Description                  |
|--------|----------------------------------|------------------------------|
| POST   | `/api/v1/appointments`           | Create an appointment        |
| GET    | `/api/v1/appointments/`          | List of all appointment      |
| GET    | `/api/v1/appointments/:id`       | Get an appointment by ID     |
| PATCH  | `/api/v1/appointments/:id/status`| Patch status                 |

**Appointment model:**
```json
{
  "id": "ObjectID",
  "title": "string",
  "description": "string",
  "doctor_id": "ObjectID",
  "status": "new | in_progress | done",
  "created_at": "datetime",
  "updated_at": "datetime"
}
```

**PATCH body:**
```json
{
  "status": "in_progress"
}
```

## Environment Variables

| Variable            | Description                            | Default                    |
|---------------------|----------------------------------------|-----------------------------|
| `MONGO_URI`         | MongoDB connection string              | `mongodb://localhost:27017` |
| `PORT`              | HTTP server port                       | `:8080`                     |
| `DOCTOR_SERVICE_URL`| URL doctor-service (appointment only)  | `http://doctor-service:8080`|
