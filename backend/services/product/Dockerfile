# Dockerfile needs to be executed from /backend folder
# docker build . -f ./services/product/Dockerfile -t product
# Run with env vars
# docker run --env-file ./secrets/.env product:latest

# Builder stage
# =============================================================================
FROM golang:1.24-bookworm AS builder

# Move to working directory /build
WORKDIR /build

RUN mkdir -p /build/.config/gcloud && \
    apt install -y ca-certificates && \
    update-ca-certificates

# Copy the entire source code into the container
COPY ./backend/jobs/sync_products/. ./jobs/sync_products

# Copy internal dependencies into container (replace statements)
COPY ./backend/pkg/. ./pkg
COPY ./backend/services/product/. ./services/product

# Change to the working directory where go.mod resides.
WORKDIR /build/services/product

# Install dependencies (external and internal)
RUN go mod download

# Build the application
# Turn off CGO to ensure static binaries
RUN CGO_ENABLED=0 go build -C cmd -o product

# Production stage
# =============================================================================
# Create a production stage to run the application binary
FROM scratch AS production

# Move to working directory /prod
WORKDIR /prod

# Copy binary from builder stage
COPY --from=builder /build/services/product/cmd/product ./
# Set a HOME directory
ENV HOME=/prod

# copy ca certs
COPY --from=builder --chown=root:root --chmod=0644 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

# Document the port that may need to be published
EXPOSE 8080

# Start the application
ENTRYPOINT ["/prod/product"]
