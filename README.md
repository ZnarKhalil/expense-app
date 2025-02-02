# Expense Tracker API

A RESTful API built with Go (Golang) for tracking personal expenses. This application allows users to manage their expenses by categories, featuring JWT authentication and MySQL database integration.

## Features

- User authentication (register/login) with JWT tokens
- Refresh token mechanism for enhanced security
- Personal expense categories management
- CRUD operations for expenses
- MySQL database integration using GORM
- RESTful API endpoints
- Protected routes with middleware authentication

## Prerequisites

- Go 1.16 or higher
- MySQL 5.7 or higher
- Git

## Installation

1. Clone the repository:

```bash
git clone https://github.com/ZnarKhalil/expense-app
cd expense-app
```

2. Install dependencies:
```bash
go mod download
```

3. Create a `.env` file based on `.env.example`:
```bash
cp .env.example .env
```

4. Configure your environment variables in `.env`:
```env
DB_USER=your_database_user
DB_PASSWORD=your_database_password
DB_HOST=localhost
DB_PORT=3306
DB_NAME=expense_tracker
JWT_SECRET=your_jwt_secret
```

5. Create the MySQL database:
```sql
CREATE DATABASE expense_tracker;
```

6. Run the application:
```bash
go run main.go
```

The server will start on `http://localhost:8080`

## API Endpoints

### Authentication

- **POST** `/register` - Register a new user
- **POST** `/login` - Login and receive JWT tokens
- **POST** `/refresh` - Refresh access token
- **POST** `/api/logout` - Logout (invalidate refresh token)

### Categories (Protected Routes)

- **GET** `/api/categories` - Get all categories
- **POST** `/api/categories` - Create a new category
- **PUT** `/api/categories/:id` - Update a category
- **DELETE** `/api/categories/:id` - Delete a category

### Expenses (Protected Routes)

- **GET** `/api/expenses` - Get all expenses
- **POST** `/api/expenses` - Create a new expense
- **PUT** `/api/expenses/:id` - Update an expense
- **DELETE** `/api/expenses/:id` - Delete an expense

## API Usage Examples

### Register a New User

```bash
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "john_doe",
    "email": "john@example.com",
    "password": "secure_password"
  }'
```

### Login

```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "secure_password"
  }'
```

### Create a New Expense Category (Protected)

```bash
curl -X POST http://localhost:8080/api/categories \
  -H "Authorization: Bearer YOUR_ACCESS_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Groceries",
    "description": "Daily grocery expenses"
  }'
```

### Create a New Expense (Protected)

```bash
curl -X POST http://localhost:8080/api/expenses \
  -H "Authorization: Bearer YOUR_ACCESS_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "expense_category_id": 1,
    "amount": 50.25,
    "date": "2024-02-20",
    "note": "Weekly groceries"
  }'
```

## Error Handling

The API returns appropriate HTTP status codes and error messages in JSON format:

```json
{
  "error": "error message here"
}
```

Common status codes:
- 200: Success
- 201: Created
- 400: Bad Request
- 401: Unauthorized
- 404: Not Found
- 500: Internal Server Error

## Security

- Passwords are hashed using bcrypt
- JWT tokens are used for authentication
- Refresh tokens are stored in the database
- Protected routes require valid JWT tokens
- Each user can only access their own data


## License

This project is licensed under the MIT License - see the LICENSE file for details

## Acknowledgments

- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [GORM](https://gorm.io)
- [JWT-Go](https://github.com/golang-jwt/jwt)
- [Godotenv](https://github.com/joho/godotenv)
