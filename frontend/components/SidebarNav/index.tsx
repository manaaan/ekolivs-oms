import { ROUTES } from '@/lib/constants'

import { SheetContent } from '@components/ui/sheet'

import SheetItem from './SheetItem'
import SidebarItem from './SidebarItem'

const navItems = [
  { href: ROUTES.DASHBOARD, label: 'Home' },
  { href: ROUTES.ORDERS, label: 'Orders' },
  { href: ROUTES.DEMANDS, label: 'Demands' },
  { href: ROUTES.PRODUCTS, label: 'Products' },
  { href: ROUTES.CUSTOMERS, label: 'Customers' },
]

export const SidebarNav = () => {
  return (
    <>
      <aside className="fixed inset-y-0 left-0 z-10 hidden w-14 flex-col border-r bg-background sm:flex">
        <nav className="flex flex-col items-center gap-4 px-2 sm:py-5">
          {navItems.map((props) => (
            <SidebarItem key={props.label} {...props} />
          ))}
        </nav>
        <nav className="mt-auto flex flex-col items-center gap-4 px-2 sm:py-5">
          <SidebarItem href={ROUTES.SETTINGS} label="Settings" />
        </nav>
      </aside>
      <SheetContent side="left" className="sm:max-w-xs">
        <nav className="grid gap-6 text-lg font-medium">
          {navItems.map((props) => (
            <SheetItem key={props.label} {...props} />
          ))}
          <SheetItem href={ROUTES.SETTINGS} label="Settings" />
        </nav>
      </SheetContent>
    </>
  )
}
