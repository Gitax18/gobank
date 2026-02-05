# GoBank

A modern, lightweight banking REST API built with Go. GoBank provides essential banking operations including user management and transaction handling, following clean architecture principles.

![Go Version](https://img.shields.io/badge/Go-1.25.6-00ADD8?style=flat&logo=go&logoColor=white)
![License](https://img.shields.io/badge/License-MIT-green.svg)
![Framework](https://img.shields.io/badge/Framework-Fiber%20v3-blue?logo=go)
![Database](https://img.shields.io/badge/Database-PostgreSQL-336791?logo=postgresql&logoColor=white)

---

## Description

**GoBank** is a RESTful banking API service designed to handle fundamental banking operations. Built with performance and scalability in mind, it leverages:

- **[Go Fiber v3](https://gofiber.io/)** – An Express-inspired web framework for high-performance HTTP routing
- **[GORM](https://gorm.io/)** – A powerful ORM for seamless PostgreSQL database interactions
- **Clean Architecture** – Modular design with separated concerns (handlers, services, repositories)

The application follows a domain-driven design pattern, organizing code by feature modules (user, transaction) with clear separation between business logic and data access layers.

---

## Features

- **User Management**
  - Create new bank users with account details
  - Retrieve user information by ID
  - Update user profile and account data
  - Delete user accounts

- **Transaction Support**
  - Transaction model with sender/receiver tracking

- **Architecture**
  - Clean, modular codebase following best practices
  - Dependency injection pattern for testability
  - Auto-migration for database schemas
  - Environment-based configuration

- **Performance**
  - Built on Fiber – one of the fastest Go web frameworks
  - Efficient PostgreSQL connection pooling via GORM
  - Lightweight and memory-efficient

---

## Installation

### Prerequisites

Ensure you have the following installed:

- [Go](https://go.dev/dl/) (v1.21 or higher)
- [PostgreSQL](https://www.postgresql.org/download/) (v13 or higher)
- Git

### Steps

1. **Clone the repository**

   ```bash
   git clone https://github.com/Gitax18/gobank.git
   cd gobank
   ```

2. **Set up environment variables**

   Copy the template and configure your database credentials:

   ```bash
   cp .env.template .env
   ```

   Edit the `.env` file with your PostgreSQL configuration:

   ```env
   DB_HOST=localhost
   DB_USER=your_postgres_username
   DB_PWD=your_postgres_password
   DB_PORT=5432
   DB_NAME=gobank
   DB_SSLM=disable
   ```

3. **Create the database**

   ```sql
   CREATE DATABASE gobank;
   ```

4. **Install dependencies**

   ```bash
   go mod download
   ```

5. **Run the application**

   ```bash
   go run ./cmd/gobank
   ```

   The server will start on `http://localhost:8080`

---

## Usage

### API Endpoints

#### User Endpoints

| Method   | Endpoint             | Description              |
|----------|----------------------|--------------------------|
| `GET`    | `/user/:id`          | Retrieve user by ID      |
| `POST`   | `/user/create`       | Create a new user        |
| `PUT`    | `/user/update/:id`   | Update user by ID        |
| `DELETE` | `/user/delete/:id`   | Delete user by ID        |

### Example Requests

#### Create a User

```bash
curl -X POST http://localhost:8080/user/create \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "number": 1234567890,
    "account_number": 9876543210,
    "balance": 5000
  }'
```

#### Get User by ID

```bash
curl http://localhost:8080/user/1
```

#### Update User

```bash
curl -X PUT http://localhost:8080/user/update/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Smith",
    "balance": 7500
  }'
```

#### Delete User

```bash
curl -X DELETE http://localhost:8080/user/delete/1
```

---

## Project Structure

```
gobank/
├── cmd/
│   └── gobank/
│       └── main.go          # Application entry point
├── internal/
│   ├── database/
│   │   └── postgres.go      # Database connection setup
│   ├── modules/
│   │   ├── user/
│   │   │   ├── user.handler.go     # HTTP handlers
│   │   │   ├── user.model.go       # Data models
│   │   │   ├── user.repository.go  # Database operations
│   │   │   ├── user.routes.go      # Route definitions
│   │   │   └── user.service.go     # Business logic
│   │   └── transaction/
│   │       ├── transaction.handler.go
│   │       ├── transaction.model.go
│   │       ├── transaction.repository.go
│   │       ├── transaction.routes.go
│   │       └── transaction.service.go
│   └── router/
│       └── router.go        # Main router setup
├── .env.template            # Environment variables template
├── go.mod                   # Go module dependencies
├── go.sum                   # Dependency checksums
└── README.md
```

---

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

---

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
