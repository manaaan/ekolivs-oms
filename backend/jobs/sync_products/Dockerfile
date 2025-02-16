# Dockerfile needs to be executed from /backend folder
# docker build . -f ./jobs/sync_products/Dockerfile -t sync-product

# Builder stage
# =============================================================================
FROM golang:1.23-bookworm AS builder

# Move to working directory /build
WORKDIR /build

# Copy the go.mod and go.sum files to the /build directory
# COPY jobs/sync_products/go.mod jobs/sync_products/go.sum ./jobs/sync_products
# Copy the entire source code into the container
COPY ./jobs/sync_products ./jobs/sync_products

# Copy internal dependencies into container (replace statements)
COPY ./pkg ./pkg
COPY ./services/product ./services/product

WORKDIR /build/jobs/sync_products
# Install dependencies (external and internal)
RUN go mod download

# Build the application
# Turn off CGO to ensure static binaries
RUN CGO_ENABLED=0 go build -C ./cmd -o sync-products

# Production stage
# =============================================================================
# Create a production stage to run the application binary
FROM scratch AS production

# Move to working directory /prod
WORKDIR /prod

# Copy binary from builder stage
COPY --from=builder /build/jobs/sync_products/cmd/sync-products ./

# Document the port that may need to be published
EXPOSE 8080

# Start the application
CMD ["/prod/sync-products"]
