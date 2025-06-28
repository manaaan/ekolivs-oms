# Ekolivs - microservices - Product

To run the service as a container in a docker compose, Here are the requirements:

1. Create a sample.env with the following key/value pairs:

ZETTLE_CLIENT_ID="Replace me with proper zettle client id"
ZETTLE_API_KEY="Replace me with with proper zettle api key"
GOOGLE_CLOUD_PROJECT="Replace me with the google cloud project"
GOOGLE_APPLICATION_CREDENTIALS="Replace me with the google service-account authentication json file"

> [!note]
> For GOOGLE_APPLICATION_CREDENTIALS, its value is the path of the service-account authentication file.
> For instance, GOOGLE_APPLICATION_CREDENTIALS="/prod/filename.json"
> On the host, the above file is placed as the same directory level as your docker-compose.yaml in a directory called "secrets"

2. Make sure that all the path exist with the proper configuration, here is the list:

- secrets/sample.env
- secrets/<filename-firestore-sa.json>

To run the service as container outside a docker compose, execute the following command:

```sh

docker run --mount type=bind,source=$(pwd)/secrets/sample.env,target=/prod/.env \
-v $(pwd)/secrets/ekolivs-firestore-sa.json:/prod/firestore-sa.json -v $(pwd)/backend/services/product:/prod/product -it \
ekolivs:0.0.1-test /prod/sync-products
```

> [!note]
> The above command needs the replacements of the base scratch container image by a container with bash/sh/zsh binaries
> otherwise scratch doesn't have a shell so you won't be able to login

i.e:

```sh
# Dockerfile
# Create a production stage to run the application binary
FROM golang:1.23.3-alpine AS production
```

To run the service on your local machine, proceed as following:

Execute `make run` under the backend/services/product

> [!note]

> Before the `make run`, ensure you have installed Golang binary and you have executed the
> `make setup` from the root of the directory.
> [read me first](../../../README.md)
