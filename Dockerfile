# Build stage
FROM golang:1.23.4-alpine AS builder

# Install git and swag
RUN apk add --no-cache git && \
    go install github.com/swaggo/swag/cmd/swag@latest

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Generate swagger docs
RUN swag init -g cmd/api/main.go --parseDependency --parseInternal

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api/main.go

# Final stage
FROM alpine:3.18

# Install certificates for HTTPS
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy the binary and make it executable
COPY --from=builder /app/main ./main
RUN chmod +x ./main

# Copy swagger docs
COPY --from=builder /app/docs ./docs

# Set environment variables
ENV GIN_MODE=release
ENV PORT=8080

# Expose port
EXPOSE 8080

# Run the application
CMD ["./main"] 