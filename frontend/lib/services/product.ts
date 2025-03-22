import { cache } from 'react'
import 'server-only'

import { GrpcError } from '@/lib/errors/grpc-error'
import type { Product__Output } from '@/proto/Product'

import { productClient } from './grpc'

export type AppProduct = ReturnType<typeof toAppProduct>

function toAppProduct({
  ID,
  name,
  sku,
  imageUrl,
  price,
  createdAt,
  updatedAt,
  status,
  vatPercentage,
}: Product__Output) {
  return {
    id: ID,
    name,
    sku,
    imageUrl,
    price,
    createdAt,
    updatedAt,
    status,
    vatPercentage,
  }
}

export const getProducts = cache(
  async (): Promise<AppProduct[]> =>
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

        resolve(value.products.map(toAppProduct))
      })
    })
)
