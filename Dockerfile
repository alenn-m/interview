# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Install required system packages
RUN apk add --no-cache git

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./main.go

# Final stage
FROM alpine:latest

WORKDIR /app

# Install make and go in the final stage
RUN apk add --no-cache make go

# Install goose and add it to PATH
ENV PATH="/root/go/bin:${PATH}"
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

# Copy the binary from builder
COPY --from=builder /app/main .
COPY --from=builder /app/api ./api
COPY --from=builder /app/Makefile .
COPY --from=builder /app/svc/migrations ./migrations

# I'm aware this will be visible on Github, but this is a test project
ENV PORT=8080
ENV DB_HOST=postgres
ENV DB_PORT=5432
ENV DB_USER=postgres
ENV DB_PASSWORD=postgres
ENV DB_SCHEMA=interview

# Expose port 8080
EXPOSE 8080

# Create a startup script
COPY --from=builder /app/scripts/start.sh .
RUN chmod +x start.sh

# Command to run the application
ENTRYPOINT ["./start.sh"]
