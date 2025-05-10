# Development stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache gcc musl-dev git

# Copy source code
COPY . .

# Download dependencies and build
RUN go mod tidy

# Remove git after build
RUN apk del git

# Deployment stage
FROM alpine:latest

WORKDIR /app

RUN go build -o main .

# Expose port
EXPOSE 8080

# Run the application
CMD ["./main"]
