##### Stage 1: Build Stage #####

# Use golang:1.21-alpine as the base image for building the application
FROM golang:1.21-alpine AS builder

# Install required dependencies including timezone data and CA certificates
RUN apk --no-cache add tzdata ca-certificates

# Set up the working directory
WORKDIR /app

# Copy Go application dependency files to leverage Docker layer caching
COPY go.mod go.sum ./

# Download Go application module dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Disable CGO for cross-platform builds
ENV CGO_ENABLED=0

# Build the Go application for a Linux OS
RUN go build -o app cmd/trekkstay/main.go

##### Stage 2: Deployment Stage #####

# Use a minimal Alpine Linux image as the base for the deployment stage
FROM alpine:latest

# Set the timezone to Asia/Ho_Chi_Minh
ENV TZ=Asia/Ho_Chi_Minh

# Install necessary runtime dependencies, including timezone data and CA certificates
RUN apk --no-cache add tzdata ca-certificates

# Create a directory for the application
WORKDIR /app

# Copy the built binary application, configuration files, and templates from the build stage
COPY --from=builder /app/app .
COPY --from=builder /app/go.mod .
COPY --from=builder /app/env ./env
COPY --from=builder /app/templates ./templates

# Copy necessary certificates and timezone information from the builder stage
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Set environment variables
ENV CONFIG_PATH=./env/trekkstay.env
ENV MIGRATE=false

# Command to run the binary application with specified configurations
ENTRYPOINT ./app -conf=${CONFIG_PATH} -migrate=${MIGRATE}

EXPOSE 8888