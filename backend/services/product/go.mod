module github.com/manaaan/ekolivs-oms/product

go 1.22.2

require (
	github.com/manaaan/ekolivs-oms/pkg v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.63.2
	google.golang.org/protobuf v1.33.0
)

require (
	github.com/golang-jwt/jwt/v5 v5.2.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/oapi-codegen/runtime v1.1.1 // indirect
	golang.org/x/net v0.21.0 // indirect
	golang.org/x/sys v0.17.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240227224415-6ceb2ff114de // indirect
)

replace github.com/manaaan/ekolivs-oms/pkg => ../../pkg
