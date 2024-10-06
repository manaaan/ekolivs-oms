// Original file: ../backend/services/product/api/service.proto
import type * as grpc from '@grpc/grpc-js'
import type { MethodDefinition } from '@grpc/proto-loader'

import type {
  ProductsRes as _ProductsRes,
  ProductsRes__Output as _ProductsRes__Output,
} from './ProductsRes'
import type {
  Empty as _google_protobuf_Empty,
  Empty__Output as _google_protobuf_Empty__Output,
} from './google/protobuf/Empty'

export interface ProductServiceClient extends grpc.Client {
  GetProducts(
    argument: _google_protobuf_Empty,
    metadata: grpc.Metadata,
    options: grpc.CallOptions,
    callback: grpc.requestCallback<_ProductsRes__Output>
  ): grpc.ClientUnaryCall
  GetProducts(
    argument: _google_protobuf_Empty,
    metadata: grpc.Metadata,
    callback: grpc.requestCallback<_ProductsRes__Output>
  ): grpc.ClientUnaryCall
  GetProducts(
    argument: _google_protobuf_Empty,
    options: grpc.CallOptions,
    callback: grpc.requestCallback<_ProductsRes__Output>
  ): grpc.ClientUnaryCall
  GetProducts(
    argument: _google_protobuf_Empty,
    callback: grpc.requestCallback<_ProductsRes__Output>
  ): grpc.ClientUnaryCall
  getProducts(
    argument: _google_protobuf_Empty,
    metadata: grpc.Metadata,
    options: grpc.CallOptions,
    callback: grpc.requestCallback<_ProductsRes__Output>
  ): grpc.ClientUnaryCall
  getProducts(
    argument: _google_protobuf_Empty,
    metadata: grpc.Metadata,
    callback: grpc.requestCallback<_ProductsRes__Output>
  ): grpc.ClientUnaryCall
  getProducts(
    argument: _google_protobuf_Empty,
    options: grpc.CallOptions,
    callback: grpc.requestCallback<_ProductsRes__Output>
  ): grpc.ClientUnaryCall
  getProducts(
    argument: _google_protobuf_Empty,
    callback: grpc.requestCallback<_ProductsRes__Output>
  ): grpc.ClientUnaryCall
}

export interface ProductServiceHandlers
  extends grpc.UntypedServiceImplementation {
  GetProducts: grpc.handleUnaryCall<
    _google_protobuf_Empty__Output,
    _ProductsRes
  >
}

export interface ProductServiceDefinition extends grpc.ServiceDefinition {
  GetProducts: MethodDefinition<
    _google_protobuf_Empty,
    _ProductsRes,
    _google_protobuf_Empty__Output,
    _ProductsRes__Output
  >
}
