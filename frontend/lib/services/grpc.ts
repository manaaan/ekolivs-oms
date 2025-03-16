import * as grpc from '@grpc/grpc-js'
import * as protoLoader from '@grpc/proto-loader'
import path from 'path'
import 'server-only'

import type { ProtoGrpcType } from '@/proto/service'

const PRODUCT_PROTO_PATH = path.join(
  process.cwd(),
  'proto-definitions/product.proto'
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
  process.env.PRODUCT_SERVICE_HOST as string,
  grpc.credentials.createInsecure()
)

export { productClient }
