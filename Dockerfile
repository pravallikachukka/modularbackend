# Use the official Go image as the base image
FROM golang:1.19

WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the Go application source code into the container
#COPY src/ .
COPY *.go ./

# Build the Go application
#RUN go build -o /modularproject/src
RUN CGO_ENABLED=0 GOOS=linux go build -o modularproject

# Expose the port that your Go application listens on
EXPOSE 8080

# Set the entry point for the container to run your Go application
CMD ["./modularproject"]

# Copy the Go application source code into the container
COPY src/ .
