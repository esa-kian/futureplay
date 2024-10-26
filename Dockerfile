# Stage 1: Build the Go application
FROM golang:1.22.2 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o futureplay ./cmd/server.go

# Stage 2: Run the Go application
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/futureplay .

# Copy the config file into the container
COPY config.yaml .  

# Expose the port on which the application will run
EXPOSE 8080

# Command to run the executable
CMD ["./futureplay"]
