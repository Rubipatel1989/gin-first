# Fixing GoAdmin Dependencies

The issue is with invalid version numbers in `go.mod`. Follow these steps to fix:

## Step 1: Clean go.mod

The `go.mod` file has been updated to remove invalid versions. The adapter and driver are sub-packages of go-admin, not separate modules.

## Step 2: Install Dependencies

Run these commands in order:

```bash
# Install main GoAdmin package (this includes adapter and drivers)
go get github.com/GoAdminGroup/go-admin@latest

# Install GoAdmin themes
go get github.com/GoAdminGroup/themes@latest

# Clean up dependencies
go mod tidy
```

## Step 3: Verify Installation

After running the above commands, verify that `go.mod` and `go.sum` are updated:

```bash
go mod verify
```

## Step 4: Run the Application

```bash
go run main.go
```

## Alternative: Use the Install Script

You can also use the provided script:

```bash
chmod +x install-deps.sh
./install-deps.sh
```

## If Issues Persist

If you still encounter issues, try:

1. **Clear Go module cache:**
   ```bash
   go clean -modcache
   ```

2. **Remove go.sum and regenerate:**
   ```bash
   rm go.sum
   go mod tidy
   ```

3. **Check Go version:**
   ```bash
   go version
   ```
   Ensure you have Go 1.21 or higher.

## Note

The adapter (`github.com/GoAdminGroup/go-admin/adapter/gin`) and driver (`github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql`) are **sub-packages** of the main `go-admin` package. They don't need separate `require` statements in `go.mod` - they're automatically included when you install the main package.

