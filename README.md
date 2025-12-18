# Gin Framework Admin Panel & API

A Go application built with Gin framework featuring **GoAdmin** admin panel for CRUD operations and API endpoints for Flutter applications.

## Features

- ✅ **GoAdmin Panel** - Professional admin panel framework with:
  - Users management (Full CRUD)
  - Stores management (Full CRUD)
  - Brands management (Full CRUD)
  - Built-in authentication and authorization
  - Advanced filtering and sorting
  - Responsive design
- ✅ RESTful API endpoints for Flutter app:
  - User list API
  - Brand list API
- ✅ MySQL database integration with GORM
- ✅ Redis integration for caching
- ✅ Environment-based configuration
- ✅ CORS middleware enabled

## Prerequisites

- Go 1.21 or higher
- MySQL 5.7+ or MySQL 8.0+
- Redis server

## Installation

1. **Clone the repository** (if applicable) or navigate to the project directory:
   ```bash
   cd gin-first
   ```

2. **Install dependencies**:
   ```bash
   go mod download
   ```

3. **Configure environment variables**:
   - Copy `.env.example` to `.env` (or edit the existing `.env` file)
   - Update the database credentials and Redis configuration:
     ```env
     DB_HOST=localhost
     DB_PORT=3306
     DB_USER=root
     DB_PASSWORD=your_password
     DB_NAME=gin_first_db
     
     REDIS_HOST=localhost
     REDIS_PORT=6379
     ```

4. **Create MySQL database**:
   ```sql
   CREATE DATABASE gin_first_db;
   ```

5. **Start Redis server** (if not running):
   ```bash
   redis-server
   ```

6. **Run the application**:
   ```bash
   go run main.go
   ```

   The server will start on port 8080 (or the port specified in `.env`).

7. **Access the GoAdmin Panel**:
   - Open your browser and navigate to: `http://localhost:8080/admin`
   - **Default Login**: GoAdmin will prompt you to create an admin user on first access
   - The admin panel provides a professional interface to manage Users, Stores, and Brands

## GoAdmin Panel

The GoAdmin panel is accessible at `http://localhost:8080/admin`. It provides:

- **Professional UI**: AdminLTE theme with modern design
- **Full CRUD Operations**: Create, Read, Update, Delete for all entities
- **Advanced Features**:
  - Filtering and searching
  - Sorting capabilities
  - Pagination
  - Export functionality
  - Custom field displays (status badges, images)
- **User-friendly Forms**: Built-in form validation and field types
- **Status Management**: Toggle active/inactive status for records

### Admin Panel Sections

1. **Users Management**: Manage user accounts with name, email, phone, and status
2. **Stores Management**: Manage store information including address, contact details
3. **Brands Management**: Manage brand information with descriptions and logos

### First Time Setup

On first access to `/admin`, GoAdmin will guide you through:
1. Creating the initial admin user
2. Setting up authentication
3. Configuring the admin panel

## API Endpoints

> **Note**: The admin panel is now handled by GoAdmin at `/admin`. The API endpoints below are for the Flutter application.

### Admin Panel (GoAdmin)
- `GET /admin` - GoAdmin login and dashboard
- GoAdmin automatically provides CRUD interfaces for:
  - `/admin/info/users` - Users management
  - `/admin/info/stores` - Stores management
  - `/admin/info/brands` - Brands management

### Legacy API Endpoints (for backward compatibility)

#### Users
- `GET /admin/users` - Get all users
- `GET /admin/users/:id` - Get user by ID
- `POST /admin/users` - Create new user
- `PUT /admin/users/:id` - Update user
- `DELETE /admin/users/:id` - Delete user

#### Stores
- `GET /admin/stores` - Get all stores
- `GET /admin/stores/:id` - Get store by ID
- `POST /admin/stores` - Create new store
- `PUT /admin/stores/:id` - Update store
- `DELETE /admin/stores/:id` - Delete store

#### Brands
- `GET /admin/brands` - Get all brands
- `GET /admin/brands/:id` - Get brand by ID
- `POST /admin/brands` - Create new brand
- `PUT /admin/brands/:id` - Update brand
- `DELETE /admin/brands/:id` - Delete brand

### Flutter API Endpoints

- `GET /api/users?page=1&limit=10` - Get paginated user list (active users only)
- `GET /api/brands?page=1&limit=10` - Get paginated brand list (active brands only)

### Health Check
- `GET /health` - Server health check

## Example API Requests

### Create User (Admin)
```bash
curl -X POST http://localhost:8080/admin/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "phone": "+1234567890"
  }'
```

### Get User List (Flutter API)
```bash
curl http://localhost:8080/api/users?page=1&limit=10
```

### Create Brand (Admin)
```bash
curl -X POST http://localhost:8080/admin/brands \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Nike",
    "description": "Just Do It",
    "logo": "https://example.com/nike-logo.png"
  }'
```

### Get Brand List (Flutter API)
```bash
curl http://localhost:8080/api/brands?page=1&limit=10
```

## Project Structure

```
gin-first/
├── admin/           # GoAdmin configuration
│   ├── goadmin_setup.go  # GoAdmin initialization
│   └── tables.go         # Table definitions for Users, Stores, Brands
├── config/          # Configuration management
├── database/        # Database connections (MySQL & Redis)
├── handlers/        # API request handlers (for Flutter)
├── middleware/      # Middleware functions
├── models/          # Data models (GORM)
├── routes/          # Route definitions
├── static/          # Static files (legacy, GoAdmin handles UI now)
├── .env            # Environment variables
├── go.mod          # Go module file
├── go.sum          # Go dependencies checksum
└── main.go         # Application entry point
```

## Database Models

### User
- ID, Name, Email, Phone, Status, CreatedAt, UpdatedAt

### Store
- ID, Name, Address, Phone, Email, Status, CreatedAt, UpdatedAt

### Brand
- ID, Name, Description, Logo, Status, CreatedAt, UpdatedAt

## Environment Variables

All configuration is managed through the `.env` file:

- `SERVER_PORT` - Server port (default: 8080)
- `GIN_MODE` - Gin mode: debug or release
- `DB_HOST` - MySQL host
- `DB_PORT` - MySQL port
- `DB_USER` - MySQL username
- `DB_PASSWORD` - MySQL password
- `DB_NAME` - MySQL database name
- `REDIS_HOST` - Redis host
- `REDIS_PORT` - Redis port
- `REDIS_PASSWORD` - Redis password (optional)
- `REDIS_DB` - Redis database number
- `JWT_SECRET` - JWT secret key (for future authentication)

## Development

The application uses GORM for database migrations. Tables are automatically created when the application starts.

## License

MIT
