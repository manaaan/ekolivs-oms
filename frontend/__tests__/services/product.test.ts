import { Metadata } from '@grpc/grpc-js'

import { GrpcError } from '@/lib/errors/GrpcError'
import { getProducts } from '@/lib/services/product'

jest.mock('@/lib/services/grpc', () => ({
  productClient: {
    getProducts: jest
      .fn()
      .mockImplementationOnce((_, callback) => {
        callback(null, { products: [{ id: '1', name: 'Product 1' }] })
      })
      .mockImplementationOnce((_, callback) => {
        callback(
          {
            code: 1,
            details: 'Mocked error',
            metadata: new Metadata(),
          },
          null
        )
      }),
  },
}))

describe('product service', () => {
  it('should return products', async () => {
    const products = await getProducts()
    expect(Array.isArray(products)).toBe(true)
    expect(products).toHaveLength(1)
    expect(products[0]).toEqual({ id: '1', name: 'Product 1' })
  })

  it('should throw an GrpcError', async () => {
    expect(getProducts()).rejects.toThrow(GrpcError)
  })
})
