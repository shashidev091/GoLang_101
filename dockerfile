# --- Stage 1: Build the application ---
FROM golang:tip-trixie AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files first (for better caching)
COPY go.mod ./
# COPY go.sum ./  # Uncomment if you have dependencies

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
# -o main: name the output binary "main"
RUN go build -o main .

# --- Stage 2: Run the application ---
FROM alpine:latest

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Expose the port the app runs on
EXPOSE 8080

# Command to run the executable
CMD ["./main"]