# ekolivs-oms
Ekolivs order management system

## Setup

```bash
# install protobuf generator binaries
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.33.0
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0
# install binary to generate structs from openapi spec
go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@v2.1.0
```