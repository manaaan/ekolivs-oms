// Original file: ../specs/demand.proto

import type { Item as _Item, Item__Output as _Item__Output } from './Item';

export interface CreateDemandReq {
  'items'?: (_Item)[];
  'member'?: (string);
}

export interface CreateDemandReq__Output {
  'items': (_Item__Output)[];
  'member': (string);
}
