FROM golang:1.21-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git gcc musl-dev

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod ./

# Download dependencies and create go.sum if it doesn't exist
RUN go mod download && go mod verify

# Copy source code
COPY . .

# Build the application with verbose output
RUN echo "Building application..." && \
    go build -v -o main . && \
    echo "Build completed successfully"

# Final stage - use smaller image
FROM alpine:latest

# Install runtime dependencies
RUN apk add --no-cache ca-certificates

WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/main .

# Expose port
EXPOSE 8080

# Run the application
CMD ["./main"]
