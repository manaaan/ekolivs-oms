import { ROUTES } from '@/lib/constants'

import SidebarItem from './SidebarItem'

export const SidebarNav = () => {
  return (
    <aside className="fixed inset-x-0 z-10 flex min-h-14 items-center border-b bg-background sm:inset-y-0 sm:h-full sm:w-14 sm:flex-col sm:border-b-0 sm:border-r">
      <nav className="flex items-center gap-4 px-2 sm:flex-col sm:py-5">
        <SidebarItem href={ROUTES.DASHBOARD} label="Home" />
        <SidebarItem href={ROUTES.ORDERS} label="Orders" />
        <SidebarItem href={ROUTES.DEMANDS} label="Demands" />
        <SidebarItem href={ROUTES.PRODUCTS} label="Products" />
        <SidebarItem href={ROUTES.CUSTOMERS} label="Customers" />
      </nav>
      <nav className="ml-auto gap-4 px-2 sm:ml-0 sm:mt-auto sm:items-center sm:py-5">
        <SidebarItem href={ROUTES.SETTINGS} label="Settings" />
      </nav>
    </aside>
  )
}
