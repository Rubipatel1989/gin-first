# GoAdmin Setup Guide

## Quick Start

1. **Install Dependencies**:
   ```bash
   go mod download
   go mod tidy
   ```

2. **Configure Environment**:
   - Update `.env` file with your MySQL and Redis credentials

3. **Create Database**:
   ```sql
   CREATE DATABASE gin_first_db;
   ```

4. **Run the Application**:
   ```bash
   go run main.go
   ```

5. **Access GoAdmin**:
   - Navigate to: `http://localhost:8080/admin`
   - On first access, GoAdmin will prompt you to create an admin user

## GoAdmin Features

### Tables Configured

1. **Users Table** (`/admin/info/users`)
   - Fields: ID, Name, Email, Phone, Status, Created At, Updated At
   - Filterable: Name, Email, Created At
   - Status badges: Active (green) / Inactive (red)

2. **Stores Table** (`/admin/info/stores`)
   - Fields: ID, Name, Address, Phone, Email, Status, Created At, Updated At
   - Filterable: Name, Email, Created At
   - Status badges: Active (green) / Inactive (red)

3. **Brands Table** (`/admin/info/brands`)
   - Fields: ID, Name, Description, Logo (image display), Status, Created At, Updated At
   - Filterable: Name, Created At
   - Logo displayed as thumbnail image
   - Status badges: Active (green) / Inactive (red)

### Customization

To customize tables, edit `admin/tables.go`:
- Add/remove fields
- Change field types
- Modify filters
- Customize displays

## API Endpoints (for Flutter)

- `GET /api/users?page=1&limit=10` - Get paginated user list
- `GET /api/brands?page=1&limit=10` - Get paginated brand list

## Troubleshooting

### GoAdmin not loading
- Ensure MySQL is running and accessible
- Check database credentials in `.env`
- Verify database exists: `CREATE DATABASE gin_first_db;`

### Dependencies issues
```bash
go mod tidy
go mod download
```

### Port already in use
- Change `SERVER_PORT` in `.env` file

