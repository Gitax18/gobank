# GoBank

A RESTful banking API built with Go. GoBank provides essential banking operations including user authentication, account management, and transaction handling. Built using Go Fiber v3 framework with PostgreSQL database and follows clean architecture principles.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
  - [Prerequisites](#prerequisites)
  - [Steps](#steps)
- [Usage](#usage)
  - [Installing Dependencies](#installing-dependencies)
  - [Setting Up Environment](#setting-up-environment)
  - [Creating Database](#creating-database)
  - [Starting the Application](#starting-the-application)
- [API Endpoints](#api-endpoints)
  - [Authentication](#authentication)
  - [User Management](#user-management)
  - [Transactions](#transactions)
- [Project Structure](#project-structure)

## Features

**Authentication**
- User registration with email and password
- JWT based authentication with HTTP-only cookies
- Protected routes with middleware authorization
- User login and logout functionality

**User Management**
- Create new bank users with account details
- Retrieve authenticated user information
- Update user profile data
- Delete user accounts

**Transactions**
- Transfer money between accounts
- Automatic balance deduction from sender
- Automatic balance addition to receiver

**Architecture**
- Clean, modular codebase with separated concerns
- Dependency injection pattern for testability
- Auto-migration for database schemas
- Environment based configuration

## Installation

### Prerequisites

- Go (v1.21 or higher)
- PostgreSQL (v13 or higher)
- Git

### Steps

1. Clone the repository

```bash
git clone https://github.com/Gitax18/gobank.git
cd gobank
```

2. Install dependencies

```bash
go mod download
```

3. Set up environment variables

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
JWT_SECRET=your_jwt_secret_key
```

4. Create the database

Connect to PostgreSQL and create the database:

```sql
CREATE DATABASE gobank;
```

5. Run the application

```bash
go run ./cmd/gobank
```

The server will start on `http://localhost:8080`

## Usage

### Installing Dependencies

```bash
go mod download
```

This command downloads all the dependencies specified in `go.mod` file.

### Setting Up Environment

1. Create a copy of the environment template:

```bash
cp .env.template .env
```

2. Open the `.env` file and fill in your database credentials and JWT secret.

### Creating Database

Connect to your PostgreSQL server using psql or any database client:

```bash
psql -U postgres
```

Then create the database:

```sql
CREATE DATABASE gobank;
```

The application will automatically create the required tables when it starts.

### Starting the Application

```bash
go run ./cmd/gobank
```

The server runs on port 8080 by default.

## API Endpoints

### Authentication

#### POST /user/register

Creates a new user account.

**Request Body**

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| email | string | Yes | User email address |
| password | string | Yes | User password |
| name | string | Yes | User full name |
| number | integer | Yes | User phone number |
| account_number | integer | Yes | Bank account number |
| balance | integer | Yes | Initial account balance |

**Request Example**

```json
{
  "email": "john@example.com",
  "password": "securepassword123",
  "name": "John Doe",
  "number": 1234567890,
  "account_number": 9876543210,
  "balance": 5000
}
```

**Response**

Status: `200 OK`

```json
{
  "message": "user created successfully"
}
```

**Errors**

| Status Code | Response |
|-------------|----------|
| 400 Bad Request | `{"message": "bad request"}` |
| 500 Internal Server Error | `{"message": "error creating user", "err": "error details"}` |

---

#### POST /user/login

Authenticates a user and sets JWT token in HTTP-only cookie.

**Request Body**

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| email | string | Yes | User email address |
| password | string | Yes | User password |

**Request Example**

```json
{
  "email": "john@example.com",
  "password": "securepassword123"
}
```

**Response**

Status: `200 OK`

Sets `Authorization` cookie with JWT token.

```json
{
  "message": "user logined successfully"
}
```

**Errors**

| Status Code | Response |
|-------------|----------|
| 400 Bad Request | `{"message": "bad request"}` |
| 400 Bad Request | `{"message": "mail not found", "err": "error details"}` |
| 400 Bad Request | `{"message": "password not found", "err": "error details"}` |
| 400 Bad Request | `{"message": "fail to generate token", "err": "error details"}` |

---

#### POST /user/logout

Logs out the authenticated user by clearing the authorization cookie.

**Headers**

| Header | Required | Description |
|--------|----------|-------------|
| Cookie: Authorization | Yes | JWT token from login |

**Response**

Status: `200 OK`

```json
{
  "message": "user logout successfully"
}
```

**Errors**

| Status Code | Response |
|-------------|----------|
| 401 Unauthorized | `{"message": "resource access forbidden"}` |
| 401 Unauthorized | `{"message": "invalid or expired token"}` |

---

### User Management

#### GET /user

Retrieves the authenticated user information.

**Headers**

| Header | Required | Description |
|--------|----------|-------------|
| Cookie: Authorization | Yes | JWT token from login |

**Response**

Status: `200 OK`

```json
{
  "message": "user found successfully",
  "data": {
    "id": 1,
    "email": "john@example.com",
    "name": "John Doe",
    "number": 1234567890,
    "account_number": 9876543210,
    "balance": 5000
  }
}
```

**Errors**

| Status Code | Response |
|-------------|----------|
| 401 Unauthorized | `{"message": "resource access forbidden"}` |
| 401 Unauthorized | `{"message": "invalid or expired token"}` |
| 403 Forbidden | `{"message": "token expired, please login"}` |
| 500 Internal Server Error | `{"message": "error getting user", "err": "error details"}` |

---

#### PUT /user/:id

Updates user profile information.

**Headers**

| Header | Required | Description |
|--------|----------|-------------|
| Cookie: Authorization | Yes | JWT token from login |

**Path Parameters**

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| id | integer | Yes | User ID |

**Request Body**

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| name | string | No | Updated user name |
| number | integer | No | Updated phone number |

**Request Example**

```json
{
  "name": "John Smith",
  "number": 9876543210
}
```

**Response**

Status: `200 OK`

```json
{
  "message": "user updated successfully"
}
```

**Errors**

| Status Code | Response |
|-------------|----------|
| 400 Bad Request | `{"message": "Id should not be empty"}` |
| 400 Bad Request | `{"message": "improper data", "data": {...}}` |
| 401 Unauthorized | `{"message": "resource access forbidden"}` |
| 401 Unauthorized | `{"message": "invalid or expired token"}` |
| 500 Internal Server Error | `{"message": "error occured while updating user", "err": "error details"}` |

---

#### DELETE /user/:id

Deletes a user account.

**Headers**

| Header | Required | Description |
|--------|----------|-------------|
| Cookie: Authorization | Yes | JWT token from login |

**Path Parameters**

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| id | integer | Yes | User ID |

**Response**

Status: `200 OK`

```json
{
  "message": "user deleted successfully"
}
```

**Errors**

| Status Code | Response |
|-------------|----------|
| 400 Bad Request | `{"message": "Id should not be empty"}` |
| 401 Unauthorized | `{"message": "resource access forbidden"}` |
| 401 Unauthorized | `{"message": "invalid or expired token"}` |
| 500 Internal Server Error | `{"message": "Error deleting user", "err": "error details"}` |

---

### Transactions

#### POST /transaction/

Creates a new transaction to transfer money between accounts.

**Headers**

| Header | Required | Description |
|--------|----------|-------------|
| Cookie: Authorization | Yes | JWT token from login |

**Request Body**

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| receiver_id | integer | Yes | ID of the receiving user |
| amount | integer | Yes | Amount to transfer |

Note: The sender ID is automatically extracted from the authenticated user token.

**Request Example**

```json
{
  "receiver_id": 2,
  "amount": 500
}
```

**Response**

Status: `200 OK`

```json
{
  "message": "Transaction completed successfully"
}
```

**Errors**

| Status Code | Response |
|-------------|----------|
| 400 Bad Request | `{"message": "Bad request", "err": "error details"}` |
| 400 Bad Request | `{"message": "Bad request", "err": "Incorrect data for making transaction"}` |
| 401 Unauthorized | `{"message": "resource access forbidden"}` |
| 401 Unauthorized | `{"message": "invalid or expired token"}` |
| 500 Internal Server Error | `{"message": "Internal server error", "err": "error details"}` |

---

## Project Structure

```
gobank/
├── cmd/
│   └── gobank/
│       └── main.go                 # Application entry point
├── internal/
│   ├── database/
│   │   └── postgres.go             # Database connection setup
│   ├── middleware/
│   │   └── checkAuth.go            # JWT authentication middleware
│   ├── modules/
│   │   ├── user/
│   │   │   ├── user.handler.go     # HTTP request handlers
│   │   │   ├── user.model.go       # User data model
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
│       └── router.go               # Main router setup
├── .env.template                   # Environment variables template
├── go.mod                          # Go module dependencies
├── go.sum                          # Dependency checksums
├── LICENSE                         # MIT License
└── README.md                       # Project documentation
```
