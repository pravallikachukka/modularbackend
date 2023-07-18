# Use the official Go image as the base image
FROM golang:1.17-alpine

WORKDIR /app

# Copy the Go application source code into the container
COPY . .

# Fetch and download the Go modules (dependencies)
RUN go mod download

# Build the Go application
RUN go build -o /modularproject/src

# Expose the port that your Go application listens on
EXPOSE 8080

# Set the entry point for the container to run your Go application
CMD ["./modularproject/src"]
