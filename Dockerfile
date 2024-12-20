# Stage 2: Create the final image
FROM golang:1.23.4 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy Go Modules manifests
COPY go.mod go.sum ./

# Download Go Modules
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o /bin/app ./cmd/app
RUN go build -o /bin/migrator ./cmd/migrator

# Stage 2: Create the final image
FROM debian:bullseye-slim

# Install PostgreSQL client, curl, and bash for wait-for-it
RUN apt-get update && apt-get install -y \
    postgresql-client \
    curl \
    bash \
    && rm -rf /var/lib/apt/lists/*

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the pre-built binary files from the builder stage
COPY --from=builder /bin/app /bin/app
COPY --from=builder /bin/migrator /bin/migrator

# Copy the wait-for-it script into the container
#COPY wait-for-it.sh /wait-for-it.sh

# Set environment variables
ENV APP_ENV=dev
ENV POSTGRES_HOST=postgres
ENV POSTGRES_PORT=5432
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=password
ENV POSTGRES_DB=maker_checker
ENV RABBIT_HOST=rabbitmq
ENV RABBIT_PORT=5672
ENV RABBIT_USER=guest
ENV RABBIT_PASSWORD=guest
ENV RABBIT_VHOST=/

# Expose the required ports
EXPOSE 3000 5672 5432

# Default entry point (you can override in docker-compose.yml)
CMD ["/bin/app"]