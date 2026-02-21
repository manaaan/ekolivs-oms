// Original file: ../specs/product.proto

import type * as grpc from '@grpc/grpc-js'
import type { MethodDefinition } from '@grpc/proto-loader'
import type { Empty as _google_protobuf_Empty, Empty__Output as _google_protobuf_Empty__Output } from './google/protobuf/Empty';
import type { Product as _Product, Product__Output as _Product__Output } from './Product';
import type { ProductIDReq as _ProductIDReq, ProductIDReq__Output as _ProductIDReq__Output } from './ProductIDReq';
import type { ProductsRes as _ProductsRes, ProductsRes__Output as _ProductsRes__Output } from './ProductsRes';

export interface ProductServiceClient extends grpc.Client {
  GetProductByID(argument: _ProductIDReq, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_Product__Output>): grpc.ClientUnaryCall;
  GetProductByID(argument: _ProductIDReq, metadata: grpc.Metadata, callback: grpc.requestCallback<_Product__Output>): grpc.ClientUnaryCall;
  GetProductByID(argument: _ProductIDReq, options: grpc.CallOptions, callback: grpc.requestCallback<_Product__Output>): grpc.ClientUnaryCall;
  GetProductByID(argument: _ProductIDReq, callback: grpc.requestCallback<_Product__Output>): grpc.ClientUnaryCall;
  getProductById(argument: _ProductIDReq, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_Product__Output>): grpc.ClientUnaryCall;
  getProductById(argument: _ProductIDReq, metadata: grpc.Metadata, callback: grpc.requestCallback<_Product__Output>): grpc.ClientUnaryCall;
  getProductById(argument: _ProductIDReq, options: grpc.CallOptions, callback: grpc.requestCallback<_Product__Output>): grpc.ClientUnaryCall;
  getProductById(argument: _ProductIDReq, callback: grpc.requestCallback<_Product__Output>): grpc.ClientUnaryCall;
  
  GetProducts(argument: _google_protobuf_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_ProductsRes__Output>): grpc.ClientUnaryCall;
  GetProducts(argument: _google_protobuf_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_ProductsRes__Output>): grpc.ClientUnaryCall;
  GetProducts(argument: _google_protobuf_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_ProductsRes__Output>): grpc.ClientUnaryCall;
  GetProducts(argument: _google_protobuf_Empty, callback: grpc.requestCallback<_ProductsRes__Output>): grpc.ClientUnaryCall;
  getProducts(argument: _google_protobuf_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_ProductsRes__Output>): grpc.ClientUnaryCall;
  getProducts(argument: _google_protobuf_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_ProductsRes__Output>): grpc.ClientUnaryCall;
  getProducts(argument: _google_protobuf_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_ProductsRes__Output>): grpc.ClientUnaryCall;
  getProducts(argument: _google_protobuf_Empty, callback: grpc.requestCallback<_ProductsRes__Output>): grpc.ClientUnaryCall;
  
  UpdateProduct(argument: _Product, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_Product__Output>): grpc.ClientUnaryCall;
  UpdateProduct(argument: _Product, metadata: grpc.Metadata, callback: grpc.requestCallback<_Product__Output>): grpc.ClientUnaryCall;
  UpdateProduct(argument: _Product, options: grpc.CallOptions, callback: grpc.requestCallback<_Product__Output>): grpc.ClientUnaryCall;
  UpdateProduct(argument: _Product, callback: grpc.requestCallback<_Product__Output>): grpc.ClientUnaryCall;
  updateProduct(argument: _Product, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_Product__Output>): grpc.ClientUnaryCall;
  updateProduct(argument: _Product, metadata: grpc.Metadata, callback: grpc.requestCallback<_Product__Output>): grpc.ClientUnaryCall;
  updateProduct(argument: _Product, options: grpc.CallOptions, callback: grpc.requestCallback<_Product__Output>): grpc.ClientUnaryCall;
  updateProduct(argument: _Product, callback: grpc.requestCallback<_Product__Output>): grpc.ClientUnaryCall;
  
}

export interface ProductServiceHandlers extends grpc.UntypedServiceImplementation {
  GetProductByID: grpc.handleUnaryCall<_ProductIDReq__Output, _Product>;
  
  GetProducts: grpc.handleUnaryCall<_google_protobuf_Empty__Output, _ProductsRes>;
  
  UpdateProduct: grpc.handleUnaryCall<_Product__Output, _Product>;
  
}

export interface ProductServiceDefinition extends grpc.ServiceDefinition {
  GetProductByID: MethodDefinition<_ProductIDReq, _Product, _ProductIDReq__Output, _Product__Output>
  GetProducts: MethodDefinition<_google_protobuf_Empty, _ProductsRes, _google_protobuf_Empty__Output, _ProductsRes__Output>
  UpdateProduct: MethodDefinition<_Product, _Product, _Product__Output, _Product__Output>
}
