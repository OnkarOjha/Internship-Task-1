# Use the official Go base image
FROM golang:1.19

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application source code to the container's working directory
COPY . .

# Run the go mod tidy command to clean up and download dependencies
RUN go mod tidy

# Build the Go application binary
RUN go build -o myapp

# Specify the command to run your application
CMD ["./myapp"]
