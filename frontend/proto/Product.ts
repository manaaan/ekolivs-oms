// Original file: ../backend/services/product/api/service.proto
import type { Price as _Price, Price__Output as _Price__Output } from './Price'
import type {
  Status as _Status,
  Status__Output as _Status__Output,
} from './Status'
import type {
  UnitType as _UnitType,
  UnitType__Output as _UnitType__Output,
} from './UnitType'
import type {
  Timestamp as _google_protobuf_Timestamp,
  Timestamp__Output as _google_protobuf_Timestamp__Output,
} from './google/protobuf/Timestamp'

export interface Product {
  id?: string
  name?: string
  sku?: string
  barcode?: string
  price?: _Price | null
  costPrice?: _Price | null
  imageUrl?: string
  vatPercentage?: string
  status?: _Status
  unitType?: _UnitType
  createdAt?: _google_protobuf_Timestamp | null
  updatedAt?: _google_protobuf_Timestamp | null
  _sku?: 'sku'
  _barcode?: 'barcode'
  _imageUrl?: 'imageUrl'
  _vatPercentage?: 'vatPercentage'
}

export interface Product__Output {
  id: string
  name: string
  sku?: string
  barcode?: string
  price: _Price__Output | null
  costPrice: _Price__Output | null
  imageUrl?: string
  vatPercentage?: string
  status: _Status__Output
  unitType: _UnitType__Output
  createdAt: _google_protobuf_Timestamp__Output | null
  updatedAt: _google_protobuf_Timestamp__Output | null
  _sku: 'sku'
  _barcode: 'barcode'
  _imageUrl: 'imageUrl'
  _vatPercentage: 'vatPercentage'
}
