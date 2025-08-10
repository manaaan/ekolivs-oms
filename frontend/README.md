# ekolivs-oms-frontend

This is the frontend application for the ekolivs-oms application. It's built using Next.js with React server components,
using grpc on the backend to fetch data from the microservice stack.

Before the startup, ensure that the backend is running and the environment variable "PRODUCT_SERVICE_HOST" is set to
proper address of the backend service such as "dns-name:port/ip address:port". i.e: "localhost:8080/127.0.0.1:8080"

To run it locally, do the following:

1. install the pnpm by using [the official documentation](https://pnpm.io/installation)
2. change directory to frontend `cd frontend`
3. Execute once `pnpm` to install dependencies
4. Lastly, always run `pnpm dev` for dev environment

To run it in docker compose:

1. At the root directory, ensure that specs directory exist
2. Execute `docker compose up`
