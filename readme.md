# Dating App Backend

## Overview
This is a backend application for a dating app, implemented in Go (Golang). The application includes functionalities for user authentication, swiping profiles, and upgrading to premium status. It uses PostgreSQL as the database and Redis for caching.

## Features
1. **User Authentication**:
   - User signup with hashed passwords.
   - User login with JWT-based authentication.

2. **Swiping Profiles**:
   - Swipe on profiles with actions like "like" or "pass".
   - Limit swipes for non-premium users to 10 per day.
   - Prevent duplicate swipes on the same profile within the same day.

3. **Premium Subscription**:
   - Upgrade users to premium status.
   - Unlimited swipes for premium users.

4. **Validation and Error Handling**:
   - Request validation with structured error responses.
   - Centralized JSON response formatting.

5. **Database Management**:
   - Auto-incrementing user IDs.
   - Database migrations for schema setup.

## Technology Stack
- **Language**: Go (Golang)
- **Database**: PostgreSQL
- **Cache**: Redis
- **Web Framework**: Gin
- **JWT**: Authentication and authorization
- **Validator**: Request validation using `go-playground/validator`

## Setup and Installation

### Prerequisites
1. Install Go (1.19 or higher).
2. Install PostgreSQL and Redis.

### Configuration
Create a `config.yaml` file in the root directory with the following structure:

```yaml
server:
  port: ":8080"

redis:
  addr: "localhost:6379"
  password: ""
  db: 0

db:
  host: "localhost"
  port: "5432"
  user: "your_user"
  password: "your_password"
  name: "dating_app"
```

### Database Setup
1. Create the PostgreSQL database:
   ```sql
   CREATE DATABASE dating_app;
   ```

2. Run the migrations (automatically executed on application start).

### Running the Application
1. Clone the repository:
   ```bash
   git clone https://github.com/your-repo/dating-app-backend.git
   cd dating-app-backend
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Run the application:
   ```bash
   go run main.go
   ```

The application will start on the port specified in `config.yaml` (default: `:8080`).

## API Endpoints

### Public Endpoints
1. **Sign Up**:
   - `POST /signup`
   - Request Body:
     ```json
     {
       "name": "John Doe",
       "email": "john@example.com",
       "password": "securepassword"
     }
     ```

2. **Login**:
   - `POST /login`
   - Request Body:
     ```json
     {
       "email": "john@example.com",
       "password": "securepassword"
     }
     ```
   - Response:
     ```json
     {
       "code": 200,
       "message": "Login successful",
       "data": {
         "token": "<JWT Token>"
       },
       "errors": null
     }
     ```

### Protected Endpoints
1. **Swipe**:
   - `POST /swipe`
   - Request Body:
     ```json
     {
       "profile_id": 123,
       "action": "like"
     }
     ```

2. **Upgrade to Premium**:
   - `POST /premium`

### Error Handling
All responses follow a consistent structure:
```json
{
  "code": <HTTP Status Code>,
  "message": "<Message>",
  "data": <Response Data>,
  "errors": <Error Details>
}
```

## Project Structure
```
dating-app-backend/
.
├── config
│   └── config.go
├── config.yaml
├── docker-compose.yml
├── go.mod
├── go.sum
├── internal
│   ├── middleware
│   │   ├── authMiddleware.go
│   │   └── auth_middleware.go
│   ├── models
│   │   ├── auth.go
│   │   ├── base_response.go
│   │   ├── claims.go
│   │   ├── premium.go
│   │   ├── swipe.go
│   │   └── user.go
│   ├── module
│   │   ├── auth
│   │   │   ├── delivery
│   │   │   │   └── http
│   │   │   │       └── auth_Handler.go
│   │   │   ├── pg_repository.go
│   │   │   ├── repository
│   │   │   │   └── auth_repository.go
│   │   │   ├── usecase
│   │   │   │   └── auth_usecase.go
│   │   │   └── usecase.go
│   │   ├── premium
│   │   │   ├── delivery
│   │   │   │   └── http
│   │   │   │       └── premium_handler.go
│   │   │   ├── pg_repository.go
│   │   │   ├── repository
│   │   │   │   └── premium_repository.go
│   │   │   ├── usecase
│   │   │   │   └── premium_usecase.go
│   │   │   └── usecase.go
│   │   └── swipe
│   │       ├── delivery
│   │       │   └── http
│   │       │       └── swipe_handler.go
│   │       ├── pg_repository.go
│   │       ├── repository
│   │       │   └── swipe_repository.go
│   │       ├── usecase
│   │       │   └── swipe_usecase.go
│   │       └── usecase.go
│   ├── repository
│   │   ├── database.go
│   │   └── redis.go
│   └── utils
│       ├── jwt.go
│       └── response.go
├── main.go
├── migrations
│   ├── 001_create_users_table.sql
│   ├── 002_create_swipe_table.sql
│   └── 003_create_premium_purchases_table.sql
├── pkg
│   └── validation
│       └── validation.go
└── readme.md
```

## Future Improvements
- Add rate limiting for requests.
- Enhance logging with structured logs.
- Implement unit tests for all components.