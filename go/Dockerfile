# Stage 1 - Build the templ tool and the app
FROM golang:1.23 AS builder

# Set the working directory
WORKDIR /app

# Copy the Go modules and install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the application
RUN go build -o main ./main.go

# Stage 2 - Create the final runtime container
FROM golang:1.23

# Set the working directory
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/main .

# Expose the port the app runs on
EXPOSE 8080

# Run the Go app
CMD ["/app/main"]
