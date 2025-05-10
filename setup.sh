#!/bin/bash

# Generate random password for PostgreSQL
DB_PASSWORD=$(openssl rand -base64 12)

# Create .env file
cat > .env << EOL
# Database Configuration
DB_HOST=db
DB_USER=postgres
DB_PASSWORD=${DB_PASSWORD}
DB_NAME=postgres
DB_PORT=5432

# PostgreSQL Configuration
POSTGRES_USER=postgres
POSTGRES_PASSWORD=${DB_PASSWORD}
POSTGRES_DB=postgres

# Application Configuration
APP_PORT=8080
EOL

# Make the script executable
chmod +x setup.sh

echo "Environment variables have been generated in .env file"
echo "Please keep your .env file secure and never commit it to version control" 