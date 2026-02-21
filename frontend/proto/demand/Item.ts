// Original file: ../specs/demand.proto

import type { Status as _Status, Status__Output as _Status__Output } from './Status';

export interface Item {
  'ID'?: (string);
  'demandID'?: (string);
  'productId'?: (string);
  'amount'?: (number);
  'position'?: (number);
  'status'?: (_Status);
  'fulfilmentDate'?: (string);
  'creationDate'?: (string);
  '_fulfilmentDate'?: "fulfilmentDate";
}

export interface Item__Output {
  'ID': (string);
  'demandID': (string);
  'productId': (string);
  'amount': (number);
  'position': (number);
  'status': (_Status__Output);
  'fulfilmentDate'?: (string);
  'creationDate': (string);
  '_fulfilmentDate': "fulfilmentDate";
}
