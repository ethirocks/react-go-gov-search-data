# Use the official Go image as the base image
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /app

# Install any required dependencies
RUN apk add --no-cache curl

# Copy go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code to the container
COPY . .

# Build the Go application
RUN go build -o react-go-gov-search-data .

# Expose the port the application listens on
EXPOSE 8080

# Command to run the application
CMD ["./react-go-gov-search-data"]
