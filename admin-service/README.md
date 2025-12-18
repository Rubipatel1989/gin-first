# Admin Panel Service

A complete admin panel built with GoAdmin framework providing full CRUD operations for managing Users, Stores, and Brands.

## Features

### Full CRUD Operations
- ✅ **Create** - Add new records through intuitive forms
- ✅ **Read** - View all records in paginated tables
- ✅ **Update** - Edit existing records inline
- ✅ **Delete** - Remove records with confirmation

### UI Features
- 🎨 **Bootstrap-based AdminLTE Theme** - Modern and responsive design
- 📊 **Data Tables** - Sortable, filterable, and searchable
- 📄 **Pagination** - Navigate through large datasets
- 🔍 **Advanced Filtering** - Filter by multiple criteria
- 📤 **Export Functionality** - Export data to various formats
- 🎯 **Field Validation** - Required fields and format validation
- 🖼️ **Image Preview** - Logo preview for brands
- 🏷️ **Status Badges** - Visual status indicators

## Available Modules

### 1. Users Management (`/admin/info/users`)
- View all users in a table
- Create new users
- Edit user information
- Delete users
- Filter by name, email, status
- Sort by ID, name, email, created date

**Fields:**
- ID (auto-generated)
- Name (required)
- Email (required, validated)
- Phone (optional)
- Status (Active/Inactive)
- Created At / Updated At (auto-managed)

### 2. Stores Management (`/admin/info/stores`)
- View all stores
- Create new stores
- Edit store details
- Delete stores
- Filter by name, address, email, status
- Sort by ID, name, email, created date

**Fields:**
- ID (auto-generated)
- Name (required)
- Address (optional, textarea)
- Phone (optional)
- Email (optional, validated)
- Status (Active/Inactive)
- Created At / Updated At (auto-managed)

### 3. Brands Management (`/admin/info/brands`)
- View all brands with logo preview
- Create new brands
- Edit brand information
- Delete brands
- Filter by name, status
- Sort by ID, name, created date

**Fields:**
- ID (auto-generated)
- Name (required)
- Description (optional, textarea)
- Logo URL (optional, with preview)
- Status (Active/Inactive)
- Created At / Updated At (auto-managed)

## Accessing the Admin Panel

1. Start the admin service:
   ```bash
   cd admin-service
   go run main.go
   ```

2. Open your browser and navigate to:
   ```
   http://localhost:8080/admin
   ```

3. Default login credentials:
   - Username: `admin`
   - Password: `admin`
   
   > **Note:** Change the default credentials in production!

## CRUD Operations Guide

### Creating a New Record
1. Navigate to any module (Users, Stores, or Brands)
2. Click the **"New"** button in the top right
3. Fill in the required fields (marked with *)
4. Click **"Submit"** to save

### Viewing Records
- All records are displayed in a table format
- Use pagination controls at the bottom to navigate
- Click column headers to sort
- Use the filter panel to search/filter records

### Editing a Record
1. Find the record in the table
2. Click the **"Edit"** icon in the row
3. Modify the fields
4. Click **"Submit"** to save changes

### Deleting a Record
1. Find the record in the table
2. Click the **"Delete"** icon in the row
3. Confirm the deletion

### Filtering and Searching
- Use the filter panel on the left
- Enter text in filter fields for text-based filtering
- Select from dropdowns for status filtering
- Use date range pickers for date filtering

### Exporting Data
- Click the **"Export"** button
- Choose export format (CSV, Excel, etc.)
- Download the exported file

## Configuration

All configuration is done through the `.env` file:

```env
ADMIN_SERVER_PORT=8080
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=gin_first_db
```

## Customization

### Adding New Tables
1. Create a new table generator function in `admin/tables.go`
2. Add it to the generators map in `admin/goadmin_setup.go`
3. Restart the service

### Modifying Table Display
Edit the table configuration in `admin/tables.go`:
- Add/remove fields
- Change field display format
- Modify filters and sorting
- Customize form layouts

### Theme Customization
The admin panel uses AdminLTE theme. You can:
- Change color scheme in `main.go` (ColorScheme field)
- Customize logo and title
- Modify animation settings

## Database

The admin panel automatically:
- Connects to MySQL database
- Runs migrations on startup
- Creates tables if they don't exist
- Manages timestamps (created_at, updated_at)

## Security Notes

⚠️ **Important for Production:**
1. Change default admin credentials
2. Use strong database passwords
3. Enable HTTPS
4. Restrict access to admin panel
5. Set up proper authentication/authorization
6. Configure firewall rules

## Troubleshooting

**Can't access admin panel:**
- Check if service is running on correct port
- Verify database connection
- Check `.env` file configuration

**Tables not showing:**
- Verify database connection
- Check if tables exist in database
- Review migration logs

**Forms not submitting:**
- Check browser console for errors
- Verify required fields are filled
- Check database constraints

## Support

For GoAdmin framework documentation, visit:
- [GoAdmin GitHub](https://github.com/GoAdminGroup/go-admin)
- [GoAdmin Documentation](https://go-admin.com)

