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

To test our gRPC services, we use Bruno: https://docs.usebruno.com/send-requests/grpc/overview
The testfiles are in this repo with the `*.bru` file ending.

### Frontend

1. Install node and pnpm

Start the frontend. In `frontend` run

```bash
pnpm install # installs dependencies
pnpm dev # starts dev server
```

More details in [frontend README](./frontend/README.md).
