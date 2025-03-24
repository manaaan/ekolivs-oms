import type * as grpc from '@grpc/grpc-js'
import type {
  EnumTypeDefinition,
  MessageTypeDefinition,
} from '@grpc/proto-loader'

import type {
  ProductServiceClient as _ProductServiceClient,
  ProductServiceDefinition as _ProductServiceDefinition,
} from './ProductService'

type SubtypeConstructor<
  Constructor extends new (...args: any) => any,
  Subtype,
> = {
  new (...args: ConstructorParameters<Constructor>): Subtype
}

export interface ProtoGrpcType {
  Price: MessageTypeDefinition
  Product: MessageTypeDefinition
  ProductService: SubtypeConstructor<
    typeof grpc.Client,
    _ProductServiceClient
  > & { service: _ProductServiceDefinition }
  ProductsRes: MessageTypeDefinition
  Status: EnumTypeDefinition
  UnitType: EnumTypeDefinition
  google: {
    protobuf: {
      Empty: MessageTypeDefinition
    }
  }
}
