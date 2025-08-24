# ekolivs-oms

Ekolivs order management system

## Onboarding

### Backend

1. Install Go
  https://grpc.io/docs/languages/go/quickstart/
2. Setup tooling
  Run make command:

  ```bash
  make setup
  ```
  
  https://grpc.io/docs/protoc-installation/

#### Testing endpoints

As we're using gRPC services, we need tools to test our endpoints.

1. Postman
   Support for gRPC requests by importing the proto specs. Can be challenging to get in updates on the gRPC requests.
2. gRPCurl: https://github.com/fullstorydev/grpcurl
   curl-like CLI to run against gRPC services.

   `go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest`

   A UI version can be found here: https://github.com/fullstorydev/grpcui

   `go install github.com/fullstorydev/grpcui/cmd/grpcui@latest`

   Example to target against local product service:

    ```
    grpcui -proto ./specs/product.proto -plaintext localhost:8080
    ```


### Frontend

1. Install node and pnpm

Start the frontend. In `frontend` run

```bash
pnpm # installs dependencies
pnpm dev # starts dev server
```