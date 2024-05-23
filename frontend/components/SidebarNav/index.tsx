import { ROUTES } from '@/lib/constants'

import SidebarItem from './SidebarItem'

export const SidebarNav = () => {
  return (
    <aside className="fixed inset-y-0 left-0 z-10 hidden w-14 flex-col border-r bg-background sm:flex">
      <nav className="flex flex-col items-center gap-4 px-2 sm:py-5">
        <SidebarItem href={ROUTES.DASHBOARD} iconName="home" label="Home" />
        <SidebarItem
          href={ROUTES.ORDERS}
          iconName="shopping-cart"
          label="Orders"
        />
        <SidebarItem href={ROUTES.DEMANDS} iconName="package" label="Demands" />
        <SidebarItem href={ROUTES.PRODUCTS} iconName="salad" label="Products" />
        <SidebarItem
          href={ROUTES.CUSTOMERS}
          iconName="users"
          label="Customers"
        />
      </nav>
      <nav className="mt-auto flex flex-col items-center gap-4 px-2 sm:py-5">
        <SidebarItem
          href={ROUTES.SETTINGS}
          iconName="settings"
          label="Settings"
        />
      </nav>
    </aside>
  )
}
