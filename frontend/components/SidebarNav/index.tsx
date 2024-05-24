import { ROUTES } from '@/lib/constants'

import SidebarItem from './SidebarItem'

export const SidebarNav = () => {
  return (
    <aside className="fixed inset-y-0 left-0 z-10 hidden w-14 flex-col border-r bg-background sm:flex">
      <nav className="flex flex-col items-center gap-4 px-2 sm:py-5">
        <SidebarItem href={ROUTES.DASHBOARD} label="Home" />
        <SidebarItem href={ROUTES.ORDERS} label="Orders" />
        <SidebarItem href={ROUTES.DEMANDS} label="Demands" />
        <SidebarItem href={ROUTES.PRODUCTS} label="Products" />
        <SidebarItem href={ROUTES.CUSTOMERS} label="Customers" />
      </nav>
      <nav className="mt-auto flex flex-col items-center gap-4 px-2 sm:py-5">
        <SidebarItem href={ROUTES.SETTINGS} label="Settings" />
      </nav>
    </aside>
  )
}
