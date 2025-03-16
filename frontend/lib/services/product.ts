import 'server-only'

import { GrpcError } from '@/lib/errors/grpc-error'
import type { Product__Output } from '@/proto/Product'

import { productClient } from './grpc'

export type Product = Product__Output

export const getProducts = async (): Promise<Product[]> =>
  new Promise((resolve, reject) => {
    productClient.getProducts({}, (error, value) => {
      if (error) {
        reject(
          new GrpcError(
            error.message,
            error.code,
            error.details,
            error.metadata,
            { cause: error }
          )
        )
        return
      }

      if (!value) {
        resolve([])
        return
      }

      resolve(value.products)
    })
  })
