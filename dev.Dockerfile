# Stage 1: Build the Go application
FROM golang:1.23 as builder

# Set the working directory inside the container
WORKDIR /app

# Cache the module dependencies
COPY go.mod ./
RUN go mod download

# Copy the rest of the application source code
COPY . ./

# Build the application
RUN go build -o app .

# Stage 2: Create a minimal image for running the application
FROM gcr.io/distroless/base-debian10

# Set the working directory inside the container
WORKDIR /

# Copy the built binary from the builder stage
COPY --from=builder /app/app .

# Expose the port the application runs on
EXPOSE 8080

# Command to run the executable
CMD ["./app"]