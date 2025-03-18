import { BananaIcon, LogOutIcon } from 'lucide-react'
import Link from 'next/link'

import { ROUTES } from '@/lib/constants'

import { SidebarMenuToggle } from '@components/sidebar-menu-toggle'
import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarGroup,
  SidebarGroupContent,
  SidebarGroupLabel,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
} from '@components/ui/sidebar'

type NavItem = {
  name: string
  href: string
  icon: React.ComponentType
}

const items: NavItem[] = [
  {
    name: 'Products',
    href: ROUTES.PRODUCTS,
    icon: BananaIcon,
  },
  {
    name: 'Log out',
    href: ROUTES.LOGIN,
    icon: LogOutIcon,
  },
]

function AppSidebar() {
  return (
    <Sidebar>
      <SidebarContent>
        <SidebarGroup>
          <SidebarGroupLabel>Ekolivs</SidebarGroupLabel>
          <SidebarGroupContent>
            <SidebarMenu>
              {items.map((item) => (
                <SidebarMenuItem key={item.name}>
                  <SidebarMenuButton asChild>
                    <Link href={item.href}>
                      <item.icon />
                      <span>{item.name}</span>
                    </Link>
                  </SidebarMenuButton>
                </SidebarMenuItem>
              ))}
            </SidebarMenu>
          </SidebarGroupContent>
        </SidebarGroup>
      </SidebarContent>
      <SidebarFooter>
        <SidebarMenu>
          <SidebarMenuItem>
            <SidebarMenuToggle />
          </SidebarMenuItem>
        </SidebarMenu>
      </SidebarFooter>
    </Sidebar>
  )
}

export { AppSidebar }
