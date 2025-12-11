# Greenlight - Movie API

This is a REST API for managing movies, built with Go as part of the "Let's Go Further" book. It's designed to demonstrate best practices for building production-ready APIs in Go.

## Project Overview

Greenlight is a JSON API that allows users to manage a collection of movies. Currently, it supports:
- Health check endpoint
- Creating movies (validation only, not persisted yet)
- Retrieving movies by ID (returns hardcoded data)

The project follows Go best practices and includes features like structured logging, input validation, custom JSON marshaling, and database connection pooling.

## Project Structure

```
├── cmd/api/           # Application entry point and HTTP handlers
│   ├── main.go        # Server setup, configuration, database connection
│   ├── routes.go      # HTTP routing configuration
│   ├── movies.go      # Movie-related HTTP handlers
│   ├── helpers.go     # JSON read/write helpers, parameter parsing
│   ├── errors.go      # Error response helpers
│   └── healthcheck.go # Health check endpoint
├── internal/          # Private application code
│   ├── data/          # Data models and database interactions
│   │   ├── models.go  # Model initialization
│   │   ├── movies.go  # Movie model and validation
│   │   └── runtime.go # Custom Runtime type with JSON marshaling
│   └── validator/     # Input validation logic
│       └── validator.go
├── migrations/        # Database schema migrations
└── go.mod             # Go module dependencies
```

## Key Concepts and Patterns

### 1. Application Structure
- Uses an `application` struct to hold shared dependencies (config, logger, models)
- Dependency injection pattern for testability and modularity

### 2. Configuration Management
- Command-line flags for runtime configuration
- Environment variables for sensitive data (database DSN)
- Structured config with defaults

### 3. HTTP Routing and Handlers
- Uses `httprouter` for efficient routing
- RESTful API design with JSON responses
- Custom error responses with appropriate HTTP status codes

### 4. JSON Handling
- Custom `envelop` type for consistent API responses
- Robust JSON parsing with error handling
- Custom types with JSON marshaling (e.g., `Runtime` type)

### 5. Input Validation
- Custom validator package with fluent API
- Field-level validation with descriptive error messages
- Business rule validation (e.g., movie year constraints)

### 6. Database Layer
- Repository pattern with model structs
- Connection pooling configuration
- Placeholder methods ready for implementation
- Database migrations for schema management

### 7. Error Handling
- Structured error responses
- Logging of internal errors
- Graceful degradation

## Dependencies

- `github.com/julienschmidt/httprouter` - HTTP router
- `github.com/lib/pq` - PostgreSQL driver
- `github.com/joho/godotenv` - Environment variable loading

## Database Schema

The `movies` table includes:
- `id` - Primary key (bigserial)
- `created_at` - Timestamp
- `title` - Movie title
- `year` - Release year (with constraints)
- `runtime` - Runtime in minutes
- `genres` - Array of genres (1-5 items)
- `version` - Optimistic locking version

## API Endpoints

### Health Check
```
GET /v1/healthcheck
```
Returns system status and version info.

### Create Movie
```
POST /v1/movies
```
Accepts JSON payload with movie data. Currently validates input but doesn't persist.

### Get Movie
```
GET /v1/movies/:id
```
Returns movie details. Currently returns hardcoded data.

## Running the Application

1. Set up PostgreSQL database
2. Create `.env` file with `GREENLIGHT_DB_DSN=your_connection_string`
3. Run migrations (you'll need a migration tool like `migrate`)
4. Build and run: `go run ./cmd/api`

## Configuration Options

- `--port` - Server port (default: 4000)
- `--env` - Environment (development/staging/production)
- `--db-dsn` - Database DSN (can use GREENLIGHT_DB_DSN env var)
- `--db-max-open-conns` - Max open DB connections (default: 25)
- `--db-max-idle-conns` - Max idle DB connections (default: 25)
- `--db-max-idle-time` - Max idle time (default: 15m)

## Next Steps

The project is in early development. To complete it:

1. Implement the database methods in `internal/data/movies.go` (Insert, Get, Update, Delete)
2. Add more endpoints (list movies, update, delete)
3. Add authentication/authorization
4. Add rate limiting and CORS
5. Add comprehensive tests
6. Add logging and monitoring

## Learning Resources

This project is based on "Let's Go Further" by Alex Edwards. For beginners learning Go:

- Focus on understanding interfaces and dependency injection
- Learn about HTTP middleware patterns
- Study the validator package for custom validation logic
- Understand database/sql patterns for data access
- Practice with JSON encoding/decoding and custom types

## Beginner Tips

- The `application` struct is a common Go pattern for web apps
- Always check errors - Go doesn't have exceptions
- Use `go mod tidy` to manage dependencies
- The `internal` directory is private to your module
- JSON tags control serialization: `json:"field,omitempty"`
- Database migrations ensure consistent schema across environments 
