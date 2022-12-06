FROM golang:1.16 as builder

# Set workdir for build to /app
WORKDIR /build
# Copy go mod files to current directory
COPY go.mod go.sum ./
# Download all dependency files
RUN go mod download -x
# Copy rest of the application source code
COPY . .
# Build application
RUN CGO_ENABLED=0 GOOS=linux go build -o cmd/http/app ./cmd/http

# Create separated image to reduce image size
FROM alpine
# Set workdir to /app
WORKDIR /app
# Copy binary file from builder to alpine image
COPY --from=builder /build/cmd/http/app .
# Copy config files
COPY --from=builder /build/config ./config
# App Environment
ENV APP_ENV docker
# Set entrypoint to binary files
ENTRYPOINT ["./app"]