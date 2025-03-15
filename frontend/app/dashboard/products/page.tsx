import { getProducts } from '@/lib/services/product'

import { columns } from '@components/product-table/columns'
import { DataTable } from '@components/product-table/data-table'

async function ProductsPage() {
  const products = await getProducts()

  return (
    <div className="container mx-auto p-4">
      <h1 className="mb-4 text-2xl font-bold">Products</h1>
      <DataTable columns={columns} data={products} />
    </div>
  )
}

export default ProductsPage
