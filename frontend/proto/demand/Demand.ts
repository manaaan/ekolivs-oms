// Original file: ../specs/demand.proto

import type { Item as _Item, Item__Output as _Item__Output } from './Item';

export interface Demand {
  'ID'?: (string);
  'items'?: (_Item)[];
  'creationDate'?: (string);
  'fulfilmentDate'?: (string);
  'member'?: (string);
  '_fulfilmentDate'?: "fulfilmentDate";
}

export interface Demand__Output {
  'ID': (string);
  'items': (_Item__Output)[];
  'creationDate': (string);
  'fulfilmentDate'?: (string);
  'member': (string);
  '_fulfilmentDate': "fulfilmentDate";
}
