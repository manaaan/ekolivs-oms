import Image from 'next/image'

import { getProducts } from '@/lib/services/product'
import { formatPrice } from '@/lib/utils'

import { SidebarTrigger } from '@components/ui/sidebar'

async function ProductPage({ params }: { params: Promise<{ id: string }> }) {
  const { id } = await params
  const products = await getProducts()
  const product = products.find((product) => product.ID === id)

  if (!product)
    return (
      <div className="container mx-auto p-4">
        <div className="flex items-center gap-2 pb-4">
          <SidebarTrigger />
          <span className="text-red-500">Product not found</span>
        </div>
      </div>
    )

  const imgUrl = product.imageUrl ?? 'https://dummyimage.com/400'

  return (
    <div className="container mx-auto p-4">
      <div className="flex items-center gap-2 pb-4">
        <SidebarTrigger />
        {product.name}
      </div>
      <div className="grid items-start gap-4 lg:w-fit lg:grid-cols-3">
        <Image
          width={400}
          height={400}
          src={imgUrl}
          priority
          alt="Product image"
          className="bg-card h-96 w-96 rounded-xl border object-cover shadow"
        />
        <div className="bg-card text-card-foreground flex flex-col gap-2 rounded-xl border p-4 shadow lg:col-span-2">
          <p>
            <span className="font-bold">Name:</span> {product.name}
          </p>
          <p>
            <span className="font-bold">Price:</span>{' '}
            {product.price?.amount &&
              formatPrice(parseInt(product.price.amount, 10))}
          </p>
          <p>
            <span className="font-bold">VAT:</span> {product.vatPercentage ?? 0}
            %
          </p>
          <p>
            <span className="font-bold">Status:</span> {product.status}
          </p>
          <p>
            <span className="font-bold">Created At:</span>{' '}
            {product.createdAt && new Date(product.createdAt).toLocaleString()}
          </p>
          <p>
            <span className="font-bold">Updated At:</span>{' '}
            {product.updatedAt && new Date(product.updatedAt).toLocaleString()}
          </p>
        </div>
      </div>
    </div>
  )
}

async function generateStaticParams() {
  const products = await getProducts()

  return products.map((product) => ({
    id: product.ID,
  }))
}

export { generateStaticParams }

export default ProductPage
