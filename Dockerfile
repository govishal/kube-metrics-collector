# Use an appropriate base image for the build phase
FROM golang:1.24-alpine AS builder

# Install dependencies required for building Go apps and Kubernetes client
RUN apk add --no-cache git

# Set the Go workspace inside the container
WORKDIR /app

# Copy go.mod and go.sum files to download the dependencies early
COPY go.mod go.sum ./

# Download Go dependencies
RUN go mod tidy

# Copy the rest of the application source code
COPY . .

# Copy the Kubernetes configuration to the correct path
COPY --chown=root:root .kube /root/.kube

# Build the Go application
RUN go build -o my-kube-client .

# Use a smaller image for the runtime phase (multi-stage build)
FROM alpine:latest

# Install dependencies required to run the Go binary (e.g., certificates, etc.)
RUN apk --no-cache add ca-certificates

# Set the working directory
WORKDIR /root/

# Copy the built binary and Kubernetes configuration from the builder phase
COPY --from=builder /app/my-kube-client /root/
COPY --from=builder /app/.kube /root/.kube

# Expose the port that the app listens on
EXPOSE 8080

# Command to run the application
CMD ["./my-kube-client"]