import { cache } from 'react'
import 'server-only'

import { GrpcError } from '@/lib/errors/grpc-error'
import type { Demand__Output } from '@/proto/demand/Demand'

import { demandClient } from './grpc'

export const getDemands = cache(
  async (): Promise<Demand__Output[]> =>
    new Promise((resolve, reject) => {
      demandClient.getDemands({}, (error, value) => {
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

        resolve(value.demands)
      })
    })
)
