import Link from 'next/link'

import { ROUTES } from '@/lib/constants'
import { getProducts } from '@/lib/services/product'

import { ProductMasonry } from '@components/create-demand/product-masonry'
import {
  Breadcrumb,
  BreadcrumbItem,
  BreadcrumbLink,
  BreadcrumbList,
  BreadcrumbPage,
  BreadcrumbSeparator,
} from '@components/ui/breadcrumb'
import { SidebarTrigger } from '@components/ui/sidebar'

async function CreateDemandPage() {
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
              <BreadcrumbLink asChild>
                <Link href={ROUTES.DEMANDS}>Demands</Link>
              </BreadcrumbLink>
            </BreadcrumbItem>
            <BreadcrumbSeparator />
            <BreadcrumbItem>
              <BreadcrumbPage>Create</BreadcrumbPage>
            </BreadcrumbItem>
          </BreadcrumbList>
        </Breadcrumb>
      </div>

      <div className="container mx-auto">
        <ProductMasonry products={products} />
      </div>
    </div>
  )
}

export default CreateDemandPage
