# Use the official golang image as a base image
FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the workspace
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code to the container
COPY . .

# Build the Go application
RUN go build -o main .

# Expose port 3001 to the outside world
EXPOSE 3001

# Define environment variable
ENV PORT=3001

# Command to run the executable
CMD ["./dd"]