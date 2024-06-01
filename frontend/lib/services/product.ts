import * as grpc from '@grpc/grpc-js'
import * as protoLoader from '@grpc/proto-loader'
import path from 'path'

import { ProtoGrpcType } from '@/proto/service'

const PROTO_PATH = path.join(
  process.cwd(),
  '../backend/services/product/api/service.proto'
)

// suggested options for similarity to loading grpc.load behavior
const packageDefinition = protoLoader.loadSync(PROTO_PATH, {
  keepCase: true,
  defaults: true,
  oneofs: true,
})

const productService = grpc.loadPackageDefinition(
  packageDefinition
) as unknown as ProtoGrpcType

export const productClient = new productService.ProductService(
  'localhost:8080',
  grpc.credentials.createInsecure()
)
