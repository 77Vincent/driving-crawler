FROM golang:1.24.3-alpine AS builder

# Set the working directory
WORKDIR /app

ENV CGO_ENABLED=0
ENV GOARCH=arm64

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./
# Download the dependencies
RUN go mod download && go mod tidy
# Copy the source code
COPY . .
# Build the Go application
RUN go build -o main .

# Use a minimal base image for the final image
FROM alpine:latest AS final

RUN apk add --no-cache chromium

# Set the working directory
WORKDIR /app
# Copy the binary from the builder stage
COPY --from=builder /app/main .

ENV CHROME_PATH=/usr/bin/chromium-browser

# Command to run the application
CMD ["./main"]
