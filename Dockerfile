# Development stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache gcc musl-dev git

# Copy go.mod and go.sum first
COPY go.mod go.sum ./

# Download dependencies and tidy
RUN go mod tidy
RUN go mod download

# Copy the rest of the source code
COPY . .

# Install gqlgen and generate code
RUN go install github.com/99designs/gqlgen@latest
RUN gqlgen generate

# Build the application
RUN go build -o main .

# Remove git after build
RUN apk del git

# Deployment stage
FROM alpine:latest

WORKDIR /app

# Copy only the binary from builder
COPY --from=builder /app/main .

# Expose port
EXPOSE 8080

# Run the application
CMD ["./main"]
