# Microservices Architecture

This repository contains two separate microservices:

1. **Admin Panel Service** - GoAdmin-based admin panel for managing data
2. **API Service** - REST API for frontend applications (Flutter, Mobile, Web)

Both services share the same database but run independently on different ports.

## Project Structure

```
.
├── admin-service/          # Admin Panel Microservice
│   ├── admin/              # GoAdmin configuration
│   ├── config/             # Configuration management
│   ├── database/           # Database connections
│   ├── models/             # Data models
│   └── main.go             # Admin service entry point
│
├── mobile-api-service/     # Mobile App API Microservice
│   ├── handlers/           # API handlers
│   ├── routes/             # Route definitions
│   ├── middleware/         # Middleware (CORS, etc.)
│   ├── config/             # Configuration management
│   ├── database/           # Database connections
│   ├── models/             # Data models
│   └── main.go             # API service entry point
│
└── README.md               # This file
```

## Prerequisites

- Go 1.21.5 or higher
- MySQL 5.7+ or MySQL 8.0+
- Redis (optional, for caching)
- Environment variables configured (see `.env.example`)

## Setup

### 1. Environment Configuration

Each service has its own `.env` file. Copy the `.env.example` file in each service directory and configure it:

**Admin Service:**
```bash
cd admin-service
cp .env.example .env
# Edit .env with your configuration
```

**Mobile API Service:**
```bash
cd mobile-api-service
cp .env.example .env
# Edit .env with your configuration
```

Each service's `.env` file should contain:

**admin-service/.env:**
```env
# Server Configuration
ADMIN_SERVER_PORT=8080
GIN_MODE=debug

# Database Configuration
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=gin_first_db

# Redis Configuration (optional)
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=
REDIS_DB=0

# JWT Secret
JWT_SECRET=your_jwt_secret_key_here
```

**mobile-api-service/.env:**
```env
# Server Configuration
API_SERVER_PORT=8081
GIN_MODE=debug

# Database Configuration
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=gin_first_db

# Redis Configuration (optional)
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=
REDIS_DB=0

# JWT Secret
JWT_SECRET=your_jwt_secret_key_here
```

### 2. Install Dependencies

Navigate to each service directory and install dependencies:

```bash
# Install Admin Service dependencies
cd admin-service
go mod download
go mod tidy

# Install Mobile API Service dependencies
cd ../mobile-api-service
go mod download
go mod tidy
```

## Running the Services

### Option 1: Run Services Separately

**Terminal 1 - Admin Panel Service:**
```bash
cd admin-service
go run main.go
```

The admin panel will be available at: `http://localhost:8080/admin`

**Terminal 2 - Mobile API Service:**
```bash
cd mobile-api-service
go run main.go
```

The Mobile API will be available at: `http://localhost:8081/api`

### Option 2: Run Both Services with Scripts

Create a simple script to run both services (see `run-services.sh` below).

## API Endpoints

### Admin Panel Service
- **Admin Panel**: `http://localhost:8080/admin`
- **Health Check**: `http://localhost:8080/health`

### Mobile API Service
- **Health Check**: `http://localhost:8081/health`
- **Get Users**: `GET http://localhost:8081/api/users?page=1&limit=10`
- **Get Brands**: `GET http://localhost:8081/api/brands?page=1&limit=10`
- **Get Stores**: `GET http://localhost:8081/api/stores?page=1&limit=10`

## Features

### Admin Panel Service
- Full CRUD operations through GoAdmin interface
- Bootstrap-based admin panel (AdminLTE theme)
- User management
- Brand management
- Store management
- Database migrations

### Mobile API Service
- RESTful API endpoints for mobile app
- Pagination support
- CORS enabled
- Only returns active records
- Health check endpoint

## Database

Both services connect to the same MySQL database. The first service to start will run migrations automatically.

## Development

### Adding New Models

1. Add the model to both `admin-service/models/` and `mobile-api-service/models/`
2. Add table configuration in `admin-service/admin/tables.go`
3. Add API handler in `mobile-api-service/handlers/`
4. Add route in `mobile-api-service/routes/routes.go`

### Building for Production

```bash
# Build Admin Service
cd admin-service
go build -o admin-service main.go

# Build Mobile API Service
cd ../mobile-api-service
go build -o mobile-api-service main.go
```

## Notes

- Both services can run simultaneously on different ports
- They share the same database schema
- Admin panel is for internal use, API is for external clients
- Default admin credentials are set by GoAdmin (check GoAdmin documentation)

## Troubleshooting

1. **Port already in use**: Change `ADMIN_SERVER_PORT` or `MOBILE_API_SERVER_PORT` in the respective service's `.env` file
2. **Database connection error**: Verify MySQL is running and credentials are correct in both `.env` files
3. **Module not found**: Run `go mod tidy` in the respective service directory
4. **Environment variables not loading**: Make sure `.env` file exists in each service directory (not in root)
