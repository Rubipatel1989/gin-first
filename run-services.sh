#!/bin/bash

# Script to run both microservices

echo "Starting Admin Panel Service..."
cd admin-service
go run main.go &
ADMIN_PID=$!

echo "Starting Mobile API Service..."
cd ../mobile-api-service
go run main.go &
MOBILE_API_PID=$!

echo "Both services are running!"
echo "Admin Panel: http://localhost:8080/admin"
echo "Mobile API Service: http://localhost:8081/api"
echo ""
echo "Press Ctrl+C to stop both services"

# Wait for interrupt signal
trap "kill $ADMIN_PID $MOBILE_API_PID; exit" INT TERM

# Wait for both processes
wait

