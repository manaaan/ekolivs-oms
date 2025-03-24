import Link from 'next/link'

import { ROUTES } from '@/lib/constants'

import {
  Breadcrumb,
  BreadcrumbItem,
  BreadcrumbList,
  BreadcrumbPage,
  BreadcrumbSeparator,
} from '@components/ui/breadcrumb'
import { Button } from '@components/ui/button'
import { SidebarTrigger } from '@components/ui/sidebar'

function DemandsPage() {
  return (
    <div className="flex-1 p-4">
      <div className="flex items-center gap-2 pb-4">
        <SidebarTrigger />
        <Breadcrumb>
          <BreadcrumbList>
            <BreadcrumbItem>Dashboard</BreadcrumbItem>
            <BreadcrumbSeparator />
            <BreadcrumbItem>
              <BreadcrumbPage>Demands</BreadcrumbPage>
            </BreadcrumbItem>
          </BreadcrumbList>
        </Breadcrumb>
      </div>

      <div className="container mx-auto">
        <Button asChild>
          <Link href={ROUTES.DEMANDS_CREATE}>Create new</Link>
        </Button>
      </div>
    </div>
  )
}

export default DemandsPage
