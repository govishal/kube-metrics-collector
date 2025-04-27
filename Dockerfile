# Use an appropriate base image
FROM golang:1.24-alpine AS builder

# Install required dependencies
RUN apk add --no-cache git

# Copy the Go module files
COPY go.mod go.sum /app/

# Set the working directory
WORKDIR /app

# Download dependencies
RUN go mod tidy

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o my-kube-client .

# Expose the required port
EXPOSE 8080

# Run the application
CMD ["./my-kube-client"]