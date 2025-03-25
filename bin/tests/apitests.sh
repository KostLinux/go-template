#!/bin/bash

set -e  # Exit on error
BASE_URL="http://localhost:8080/api/v1"

# Function to extract ID from response
get_id() {
    echo $1 | grep -o '"id":[0-9]*' | grep -o '[0-9]*'
}

# Create a user
echo "Creating user..."
CREATE_RESPONSE=$(curl -s -X POST "${BASE_URL}/users" \
-H 'Content-Type: application/json' \
-d '{
    "name": "Test User",
    "email": "test.'"$(date +%s)"'@example.com",
    "password": "secretpassword123"
}')
echo $CREATE_RESPONSE
echo

# Extract user ID from response
USER_ID=$(get_id "$CREATE_RESPONSE")
echo "Created user ID: $USER_ID"
echo

# Get all users
echo "Getting all users..."
curl -s -X GET "${BASE_URL}/users"
echo

# Get user by ID
echo "Getting user with ID $USER_ID..."
curl -s -X GET "${BASE_URL}/users/$USER_ID"
echo

# Update user
echo "Updating user with ID $USER_ID..."
curl -s -X PUT "${BASE_URL}/users/$USER_ID" \
-H 'Content-Type: application/json' \
-d '{
    "name": "Updated User",
    "email": "updated.'"$(date +%s)"'@example.com"
}'
echo

# Get updated user
echo "Getting updated user..."
curl -s -X GET "${BASE_URL}/users/$USER_ID"
echo

# Delete user
echo "Deleting user with ID $USER_ID..."
curl -s -X DELETE "${BASE_URL}/users/$USER_ID"
echo

# Verify deletion
echo "Verifying deletion..."
curl -s -X GET "${BASE_URL}/users/$USER_ID"
echo