FROM --platform=linux/amd64 debian:stable-slim
# # Build stage
# FROM --platform=linux/amd64 golang:1.23-alpine AS builder

# # Set working directory
# WORKDIR /app

# # Install git (needed for go mod download)
# RUN apk add --no-cache git

# # Copy go mod files
# COPY go.mod go.sum ./

# # Download dependencies
# RUN go mod download

# # Copy source code
# COPY . .

# # Build the application
# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o notely .

# # Final stage
# FROM --platform=linux/amd64 alpine:latest

# # Install ca-certificates for HTTPS requests
# RUN apk --no-cache add ca-certificates tzdata

# # Create non-root user
# RUN addgroup -g 1001 -S notely && \
#     adduser -S notely -u 1001

# WORKDIR /app

# # Copy binary from builder stage
# COPY --from=builder /app/notely .

# # Copy static files
# COPY --from=builder /app/static ./static

# # Change ownership to non-root user
# RUN chown -R notely:notely /app
# USER notely

# # Expose port
# EXPOSE 8080

# # Command to run
# CMD ["./notely"]
