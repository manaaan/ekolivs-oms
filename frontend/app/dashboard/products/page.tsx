import { getProducts } from '@/lib/services/product'

import { columns } from '@components/product-table/columns'
import { DataTable } from '@components/product-table/data-table'
import {
  Breadcrumb,
  BreadcrumbItem,
  BreadcrumbList,
  BreadcrumbPage,
  BreadcrumbSeparator,
} from '@components/ui/breadcrumb'
import { SidebarTrigger } from '@components/ui/sidebar'

async function ProductsPage() {
  const products = await getProducts()

  return (
    <div className="flex-1 p-4">
      <div className="flex items-center gap-2 pb-4">
        <SidebarTrigger />
        <Breadcrumb>
          <BreadcrumbList>
            <BreadcrumbItem>Dashboard</BreadcrumbItem>
            <BreadcrumbSeparator />
            <BreadcrumbItem>
              <BreadcrumbPage>Products</BreadcrumbPage>
            </BreadcrumbItem>
          </BreadcrumbList>
        </Breadcrumb>
      </div>

      <div className="container mx-auto">
        <DataTable columns={columns} data={products} />
      </div>
    </div>
  )
}

export default ProductsPage
