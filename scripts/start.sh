#!/bin/sh

# Wait for postgres to be ready
echo "Waiting for postgres..."
while ! nc -z $DB_HOST $DB_PORT; do
  sleep 1
done
echo "Postgres is up"

# Run migrations
echo "Running migrations..."
goose -dir ./migrations postgres "user=$DB_USER password=$DB_PASSWORD dbname=$DB_SCHEMA host=$DB_HOST port=$DB_PORT sslmode=disable" up

# Start the application
echo "Starting application..."
./main api 