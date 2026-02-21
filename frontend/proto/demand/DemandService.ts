// Original file: ../specs/demand.proto

import type * as grpc from '@grpc/grpc-js'
import type { MethodDefinition } from '@grpc/proto-loader'
import type { CreateDemandReq as _CreateDemandReq, CreateDemandReq__Output as _CreateDemandReq__Output } from './CreateDemandReq';
import type { Demand as _Demand, Demand__Output as _Demand__Output } from './Demand';
import type { DemandsReq as _DemandsReq, DemandsReq__Output as _DemandsReq__Output } from './DemandsReq';
import type { DemandsRes as _DemandsRes, DemandsRes__Output as _DemandsRes__Output } from './DemandsRes';
import type { Empty as _google_protobuf_Empty, Empty__Output as _google_protobuf_Empty__Output } from './google/protobuf/Empty';
import type { IdReq as _IdReq, IdReq__Output as _IdReq__Output } from './IdReq';

export interface DemandServiceClient extends grpc.Client {
  CreateDemand(argument: _CreateDemandReq, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_Demand__Output>): grpc.ClientUnaryCall;
  CreateDemand(argument: _CreateDemandReq, metadata: grpc.Metadata, callback: grpc.requestCallback<_Demand__Output>): grpc.ClientUnaryCall;
  CreateDemand(argument: _CreateDemandReq, options: grpc.CallOptions, callback: grpc.requestCallback<_Demand__Output>): grpc.ClientUnaryCall;
  CreateDemand(argument: _CreateDemandReq, callback: grpc.requestCallback<_Demand__Output>): grpc.ClientUnaryCall;
  createDemand(argument: _CreateDemandReq, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_Demand__Output>): grpc.ClientUnaryCall;
  createDemand(argument: _CreateDemandReq, metadata: grpc.Metadata, callback: grpc.requestCallback<_Demand__Output>): grpc.ClientUnaryCall;
  createDemand(argument: _CreateDemandReq, options: grpc.CallOptions, callback: grpc.requestCallback<_Demand__Output>): grpc.ClientUnaryCall;
  createDemand(argument: _CreateDemandReq, callback: grpc.requestCallback<_Demand__Output>): grpc.ClientUnaryCall;
  
  DeleteDemand(argument: _IdReq, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  DeleteDemand(argument: _IdReq, metadata: grpc.Metadata, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  DeleteDemand(argument: _IdReq, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  DeleteDemand(argument: _IdReq, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  deleteDemand(argument: _IdReq, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  deleteDemand(argument: _IdReq, metadata: grpc.Metadata, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  deleteDemand(argument: _IdReq, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  deleteDemand(argument: _IdReq, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  
  GetDemand(argument: _IdReq, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_Demand__Output>): grpc.ClientUnaryCall;
  GetDemand(argument: _IdReq, metadata: grpc.Metadata, callback: grpc.requestCallback<_Demand__Output>): grpc.ClientUnaryCall;
  GetDemand(argument: _IdReq, options: grpc.CallOptions, callback: grpc.requestCallback<_Demand__Output>): grpc.ClientUnaryCall;
  GetDemand(argument: _IdReq, callback: grpc.requestCallback<_Demand__Output>): grpc.ClientUnaryCall;
  getDemand(argument: _IdReq, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_Demand__Output>): grpc.ClientUnaryCall;
  getDemand(argument: _IdReq, metadata: grpc.Metadata, callback: grpc.requestCallback<_Demand__Output>): grpc.ClientUnaryCall;
  getDemand(argument: _IdReq, options: grpc.CallOptions, callback: grpc.requestCallback<_Demand__Output>): grpc.ClientUnaryCall;
  getDemand(argument: _IdReq, callback: grpc.requestCallback<_Demand__Output>): grpc.ClientUnaryCall;
  
  GetDemands(argument: _DemandsReq, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_DemandsRes__Output>): grpc.ClientUnaryCall;
  GetDemands(argument: _DemandsReq, metadata: grpc.Metadata, callback: grpc.requestCallback<_DemandsRes__Output>): grpc.ClientUnaryCall;
  GetDemands(argument: _DemandsReq, options: grpc.CallOptions, callback: grpc.requestCallback<_DemandsRes__Output>): grpc.ClientUnaryCall;
  GetDemands(argument: _DemandsReq, callback: grpc.requestCallback<_DemandsRes__Output>): grpc.ClientUnaryCall;
  getDemands(argument: _DemandsReq, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_DemandsRes__Output>): grpc.ClientUnaryCall;
  getDemands(argument: _DemandsReq, metadata: grpc.Metadata, callback: grpc.requestCallback<_DemandsRes__Output>): grpc.ClientUnaryCall;
  getDemands(argument: _DemandsReq, options: grpc.CallOptions, callback: grpc.requestCallback<_DemandsRes__Output>): grpc.ClientUnaryCall;
  getDemands(argument: _DemandsReq, callback: grpc.requestCallback<_DemandsRes__Output>): grpc.ClientUnaryCall;
  
}

export interface DemandServiceHandlers extends grpc.UntypedServiceImplementation {
  CreateDemand: grpc.handleUnaryCall<_CreateDemandReq__Output, _Demand>;
  
  DeleteDemand: grpc.handleUnaryCall<_IdReq__Output, _google_protobuf_Empty>;
  
  GetDemand: grpc.handleUnaryCall<_IdReq__Output, _Demand>;
  
  GetDemands: grpc.handleUnaryCall<_DemandsReq__Output, _DemandsRes>;
  
}

export interface DemandServiceDefinition extends grpc.ServiceDefinition {
  CreateDemand: MethodDefinition<_CreateDemandReq, _Demand, _CreateDemandReq__Output, _Demand__Output>
  DeleteDemand: MethodDefinition<_IdReq, _google_protobuf_Empty, _IdReq__Output, _google_protobuf_Empty__Output>
  GetDemand: MethodDefinition<_IdReq, _Demand, _IdReq__Output, _Demand__Output>
  GetDemands: MethodDefinition<_DemandsReq, _DemandsRes, _DemandsReq__Output, _DemandsRes__Output>
}
