FROM golang:1.20-alpine

# Install bash (or sh) for interactive shell
RUN apk add --no-cache bash

# Set up application directory
WORKDIR /app

# Copy the site code and assets into the image
COPY go ./go
COPY templates ./templates
COPY static ./static

# Work inside the go folder so relative paths resolve correctly
WORKDIR /app/go

# Initialize a Go module to allow `go run`
RUN go mod init site

# Expose port 8080 for the HTTP server
EXPOSE 8080

# Drop into shell by default so you can run code manually
CMD ["go", "run", "main.go"]
