// Original file: ../specs/product.proto

import type { Long } from '@grpc/proto-loader';

export interface Price {
  'amount'?: (number | string | Long);
  'currencyID'?: (string);
}

export interface Price__Output {
  'amount': (string);
  'currencyID': (string);
}
