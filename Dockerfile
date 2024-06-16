# Stage 1: Build the Go binary
FROM golang:1.22 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/pdfServer

# Stage 2: Copy the binary to a small image
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/pdfServer .
COPY pdfs /root/pdfs

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./pdfServer", "--listenAddr", ":8080"]
