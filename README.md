## Getting Started

### Running with Docker Compose

1. Make sure you have Docker and Docker Compose installed on your system
2. Clone this repository
3. Run the application using Docker Compose:

```bash
docker-compose up -d
```

This will build and start all the required services in detached mode.

### API Documentation (Swagger)

The API documentation is available through Swagger UI. After starting the application, you can access the Swagger documentation at:

```
http://localhost:8080/swagger/index.html
```

This provides an interactive interface to explore and test the API endpoints.

## Development

To stop the application and remove the containers:

```bash
docker-compose down
```

To rebuild and restart the services:

```bash
docker-compose up -d --build
```
