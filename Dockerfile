FROM --platform=linux/amd64 debian:stable-slim

# Install ca-certificates for HTTPS requests
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Create non-root user
RUN groupadd -g 1001 notely && \
    useradd -r -u 1001 -g notely notely

WORKDIR /app

# Copy the pre-built binary (built by buildprod.sh)
COPY notely .

# Copy static files
COPY static ./static

# Change ownership to non-root user
RUN chown -R notely:notely /app
USER notely

# Expose port
EXPOSE 8080

# Command to run
CMD ["./notely"]
