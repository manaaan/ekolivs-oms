# Dockerfile needs to be executed from /backend folder
# docker build . -f ./services/product/Dockerfile -t product
# Run with env vars
# docker run --env-file ./services/product/.env product:latest

# Builder stage
# =============================================================================
FROM golang:1.23-bookworm AS builder

# Move to working directory /build
WORKDIR /build

# Copy the entire source code into the container
COPY ./services/product ./services/product

# Copy internal dependencies into container (replace statements)
COPY ./pkg ./pkg

WORKDIR /build/services/product
# Install dependencies (external and internal)
RUN go mod download

# Build the application
# Turn off CGO to ensure static binaries
RUN CGO_ENABLED=0 go build -C ./cmd -o product

# Production stage
# =============================================================================
# Create a production stage to run the application binary
FROM scratch AS production

# Move to working directory /prod
WORKDIR /prod

# Copy binary from builder stage
COPY --from=builder /build/services/product/cmd/product ./

# Document the port that may need to be published
EXPOSE 8080

# Start the application
CMD ["/prod/product"]
