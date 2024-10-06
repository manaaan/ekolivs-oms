import { getProducts } from '@/lib/services/product'

import ProductsTableDefinition from './ProductsTableDefinition'

const ProductsTable = async () => {
  const products = await getProducts()

  return <ProductsTableDefinition products={products} />
}

export default ProductsTable
