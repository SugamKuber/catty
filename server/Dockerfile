# Use Go 1.22.4 as the base image
FROM golang:1.22.4

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the workspace
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the entire project into the container
COPY . .

# Build the Go app
RUN go build -o main ./cmd/main.go

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
