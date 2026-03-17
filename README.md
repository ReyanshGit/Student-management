# Student Management API

A simple REST API for managing students built with Golang and MySQL.

## Tech Stack
- **Golang** — Backend language
- **Gin** — Web framework
- **MySQL** — Database
- **GORM** — ORM

## Project Structure
```
student-management/
├── main.go
├── .env
├── config/
│   └── database.go
├── models/
│   └── student.go
├── controllers/
│   └── student.go
└── routes/
    └── routes.go
```

## Setup

### 1. Database banao
```sql
CREATE DATABASE studentdb;
```

### 2. .env file banao
```
DB_USER=root
DB_PASSWORD=yourpassword
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=studentdb
```

### 3. Dependencies install karo
```bash
go mod tidy
```

### 4. Run karo
```bash
go run main.go
```

## API Endpoints

| Method | URL | Description |
|--------|-----|-------------|
| GET | /api/students | Sab students dekho |
| GET | /api/students/:id | Ek student dekho |
| POST | /api/students | Naya student add karo |
| PUT | /api/students/:id | Student update karo |
| DELETE | /api/students/:id | Student delete karo |

## Author
**Rupendra Kumar** — Golang Backend Developer
GitHub: [ReyanshGit](https://github.com/ReyanshGit)