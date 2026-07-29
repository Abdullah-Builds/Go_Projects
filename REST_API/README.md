# Go REST API - Student Management System

![Go](https://img.shields.io/badge/Go-1.26%2B-00ADD8?logo=go&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-Supported-2496ED?logo=docker&logoColor=white)
![GitHub](https://img.shields.io/badge/GitHub-Repository-181717?logo=github&logoColor=white)
![License](https://img.shields.io/badge/License-MIT-green)

A production-ready REST API built with Go, featuring clean architecture, Docker support, and SQLite database integration. This project demonstrates best practices in building scalable Go applications.

## Overview

This is a comprehensive RESTful API for managing student records. It's built with Go's standard library and follows clean architecture principles, making it maintainable, testable, and scalable. The API handles CRUD operations for student records with validation and error handling.

## Features

- ✅ **Clean Architecture** - Well-organized code structure with separation of concerns
- ✅ **RESTful API** - Standard HTTP methods for CRUD operations
- ✅ **SQLite Database** - Lightweight, file-based database perfect for learning and small deployments
- ✅ **Input Validation** - Comprehensive validation using `go-playground/validator`
- ✅ **Configuration Management** - YAML-based configuration with environment variable support
- ✅ **Docker Support** - Ready for containerization with Docker and Docker Compose
- ✅ **Error Handling** - Structured error responses with proper HTTP status codes
- ✅ **Logging** - Structured logging using Go's `slog` package

## Prerequisites

- **Go** 1.26 or higher
- **Docker** and **Docker Compose** (optional, for containerized deployment)
- **Git** (for cloning the repository)

## Project Structure

```
Go_REST_API/
├── cmd/
│   └── server/
│       └── main.go                 # Application entry point
├── internal/
│   ├── config/
│   │   └── config.go              # Configuration loading and parsing
│   ├── http/
│   │   └── handlers/
│   │       └── student/
│   │           └── student.go     # Student HTTP handlers
│   ├── storage/
│   │   ├── storage.go             # Storage interface definitions
│   │   └── sqlite/
│   │       └── sqlite.go          # SQLite implementation
│   ├── types/
│   │   └── type.go                # Data models and structs
│   └── utils/
│       └── response/
│           └── response.go        # Response formatting utilities
├── config/
│   └── local.yaml                 # Local configuration file
├── storage/                        # Database storage directory
├── Dockerfile                      # Docker container configuration
├── docker-compose.yml             # Docker Compose orchestration
├── go.mod                         # Go module file
└── README.md                      # This file
```

## Installation

### 1. Clone the Repository

```bash
git clone https://github.com/Abdullah-Builds/Go_Projects.git
cd Go_REST_API
```

### 2. Install Dependencies

```bash
go mod download
```

### 3. Verify Installation

```bash
go mod verify
```

## Configuration

The application uses YAML-based configuration located in `config/local.yaml`:

```yaml
env: "dev"                           # Environment: dev, prod
storage_path: "storage/storage.db"   # Path to SQLite database file
http_server:
  address: "localhost:8082"          # Server host and port
```

### Environment Variables

You can override configuration values using environment variables:

- `CONFIG_PATH` - Path to the configuration file (required if not using `-config` flag)
- `ENV` - Environment type
- Individual fields can be set via environment variables matching the YAML structure

## Running the Application

### Method 1: Direct Execution

```bash
go run ./cmd/server/main.go -config config/local.yaml
```

### Method 2: Build and Run

```bash
# Build the application
go build -o bin/server ./cmd/server

# Run the binary
./bin/server -config config/local.yaml
```

### Method 3: Docker Compose

```bash
docker-compose up --build
```

The API will start on `http://localhost:8082`

## API Endpoints

### 1. Create a Student
**POST** `/api/student`

Create a new student record.

**Request Body:**
```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "age": 20
}
```

**Response:**
```json
{
  "status": "ok",
  "statusCode": 201,
  "data": {
    "id": 1
  }
}
```

**Status Codes:**
- `201` - Student created successfully
- `400` - Bad request (validation error or invalid JSON)

---

### 2. Get All Students
**GET** `/api/student`

Retrieve all student records from the database.

**Response:**
```json
{
  "status": "ok",
  "statusCode": 200,
  "data": [
    {
      "id": 1,
      "name": "John Doe",
      "email": "john@example.com",
      "age": 20
    },
    {
      "id": 2,
      "name": "Jane Smith",
      "email": "jane@example.com",
      "age": 21
    }
  ]
}
```

**Status Codes:**
- `200` - Successfully retrieved all students
- `400` - Bad request

---

### 3. Get Student by ID
**GET** `/api/student/{id}`

Retrieve a specific student by their ID.

**URL Parameters:**
- `id` (integer) - Student ID

**Response:**
```json
{
  "status": "ok",
  "statusCode": 200,
  "data": {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com",
    "age": 20
  }
}
```

**Status Codes:**
- `200` - Student found
- `400` - Bad request or student not found

---

## Data Models

### Student

| Field | Type | Required | Validation | Description |
|-------|------|----------|-----------|-------------|
| `name` | string | Yes | Required | Student's full name |
| `email` | string | Yes | Required | Student's email address |
| `age` | integer | Yes | Required | Student's age |

## Development

### Project Dependencies

- `github.com/go-playground/validator/v10` - Input validation
- `github.com/ilyakaznacheev/cleanenv` - Configuration loading
- `modernc.org/sqlite` - SQLite driver

### Key Technologies

- **Standard Library** - Uses Go's built-in `net/http` package for HTTP handling
- **YAML** - Configuration format
- **SQLite** - Lightweight database

### Code Style

The project follows Go conventions:
- Standard `go fmt` formatting
- Clear variable naming
- Proper error handling
- Package-based organization

## Docker Deployment

### Build Docker Image

```bash
docker build -t go-rest-api:latest .
```

### Run with Docker Compose

```bash
docker-compose up
```

### Run Standalone Container

```bash
docker run -p 8082:8082 -v ./storage:/app/storage go-rest-api:latest
```

## Error Handling

The API returns structured error responses for validation failures:

```json
{
  "status": "error",
  "statusCode": 400,
  "error": {
    "field": "email",
    "message": "Field validation for 'email' failed on the 'required' tag"
  }
}
```

## Testing

To test the API endpoints, you can use:

- **cURL**
  ```bash
  curl -X POST http://localhost:8082/api/student \
    -H "Content-Type: application/json" \
    -d '{"name":"John Doe","email":"john@example.com","age":20}'
  ```

- **Postman** - Import endpoints and test with GUI
- **REST Client** - VS Code extension for inline testing

## Troubleshooting

### Port Already in Use
If port 8082 is already in use, edit `config/local.yaml` and change the address:
```yaml
http_server:
  address: "localhost:8083"
```

### Database Permission Issues
Ensure the `storage/` directory exists and has write permissions:
```bash
mkdir -p storage
chmod 755 storage
```

### Configuration Not Found
Set the `CONFIG_PATH` environment variable:
```bash
export CONFIG_PATH=config/local.yaml
go run ./cmd/server/main.go
```

## Future Enhancements

- [ ] Add update (PUT/PATCH) endpoints
- [ ] Add delete (DELETE) endpoint
- [ ] Implement authentication and authorization
- [ ] Add pagination support
- [ ] Create comprehensive unit and integration tests
- [ ] Add database migrations
- [ ] Implement caching layer
- [ ] Add API rate limiting
- [ ] Create OpenAPI/Swagger documentation

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Author

Created as part of the Go_Projects learning series by Abdullah Builds.

---

For more Go learning examples, check out the [Golang folder](../Golang/) in this repository.
