# Use the official Go image as a base
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app


# Copy go.mod and go.sum files to the working directory
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go mod download

# Copy the entire project to the working directory
COPY . .

# Install MySQL client library
RUN apt-get update && apt-get install -y default-mysql-client

# Install Redis client library
RUN go get github.com/go-redis/redis/v8

# Build the Go application
RUN go build -o wallet-api cmd/wallet-api/main.go

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./wallet-api"]
