# Use the official Go image as the base image
FROM golang:1.19

WORKDIR /app

#COPY go.mod go.sum ./
RUN go mod init

# Copy the Go application source code into the container
COPY src/ .

# Build the Go application
#RUN go build -o /modularproject/src
RUN CGO_ENABLED=0 go build -o modularproject

# Expose the port that your Go application listens on
EXPOSE 8080

# Set the entry point for the container to run your Go application
CMD ["./modularproject"]
