# Use the official Golang image
FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Initialize and download the necessary Go modules
RUN go mod init github.com/RodrigoMS/CRUD-Go-Gin && go mod tidy

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]

