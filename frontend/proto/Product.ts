// Original file: ../backend/services/product/api/service.proto

import type { Price as _Price, Price__Output as _Price__Output } from './Price';
import type { Status as _Status, Status__Output as _Status__Output } from './Status';
import type { UnitType as _UnitType, UnitType__Output as _UnitType__Output } from './UnitType';

export interface Product {
  'ID'?: (string);
  'name'?: (string);
  'sku'?: (string);
  'barcode'?: (string);
  'price'?: (_Price | null);
  'costPrice'?: (_Price | null);
  'imageUrl'?: (string);
  'vatPercentage'?: (string);
  'status'?: (_Status);
  'unitType'?: (_UnitType);
  'createdAt'?: (string);
  'updatedAt'?: (string);
  '_sku'?: "sku";
  '_barcode'?: "barcode";
  '_imageUrl'?: "imageUrl";
  '_vatPercentage'?: "vatPercentage";
  '_createdAt'?: "createdAt";
  '_updatedAt'?: "updatedAt";
}

export interface Product__Output {
  'ID': (string);
  'name': (string);
  'sku'?: (string);
  'barcode'?: (string);
  'price': (_Price__Output | null);
  'costPrice': (_Price__Output | null);
  'imageUrl'?: (string);
  'vatPercentage'?: (string);
  'status': (_Status__Output);
  'unitType': (_UnitType__Output);
  'createdAt'?: (string);
  'updatedAt'?: (string);
  '_sku': "sku";
  '_barcode': "barcode";
  '_imageUrl': "imageUrl";
  '_vatPercentage': "vatPercentage";
  '_createdAt': "createdAt";
  '_updatedAt': "updatedAt";
}
