FROM golang:1.17-alpine

COPY go.mod ./

RUN go mod download

COPY *.go ./

# Build the Go application
RUN go build -o /modularproject/src

# Expose the port that your Go application listens on
EXPOSE 8080

# Set the entry point for the container to run your Go application
CMD ["./modularproject/src"]
