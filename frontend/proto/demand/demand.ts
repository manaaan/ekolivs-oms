import type * as grpc from '@grpc/grpc-js';
import type { EnumTypeDefinition, MessageTypeDefinition } from '@grpc/proto-loader';

import type { DemandServiceClient as _DemandServiceClient, DemandServiceDefinition as _DemandServiceDefinition } from './DemandService';

type SubtypeConstructor<Constructor extends new (...args: any) => any, Subtype> = {
  new(...args: ConstructorParameters<Constructor>): Subtype;
};

export interface ProtoGrpcType {
  CreateDemandReq: MessageTypeDefinition
  Demand: MessageTypeDefinition
  DemandService: SubtypeConstructor<typeof grpc.Client, _DemandServiceClient> & { service: _DemandServiceDefinition }
  DemandsReq: MessageTypeDefinition
  DemandsRes: MessageTypeDefinition
  IdReq: MessageTypeDefinition
  Item: MessageTypeDefinition
  Status: EnumTypeDefinition
  google: {
    protobuf: {
      Empty: MessageTypeDefinition
    }
  }
}

