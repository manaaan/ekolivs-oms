import * as grpc from '@grpc/grpc-js'
import * as protoLoader from '@grpc/proto-loader'
import path from 'path'

import { ProtoGrpcType } from '@/proto/service'

const PRODUCT_PROTO_PATH = path.join(
  process.cwd(),
  '../backend/services/product/api/service.proto'
)

const productDefinition = protoLoader.loadSync(PRODUCT_PROTO_PATH, {
  longs: String,
  enums: String,
  defaults: true,
  oneofs: true,
})

const productGrpc = grpc.loadPackageDefinition(
  productDefinition
) as unknown as ProtoGrpcType

// Add more clients here
const productClient = new productGrpc.ProductService(
  'localhost:8080',
  grpc.credentials.createInsecure()
)

export { productClient }
