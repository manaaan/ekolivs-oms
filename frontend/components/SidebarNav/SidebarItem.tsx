'use client'

import {
  Home,
  Package,
  Salad,
  Settings,
  ShoppingBag,
  Users,
} from 'lucide-react'
import Link from 'next/link'
import { usePathname } from 'next/navigation'

import { ROUTES } from '@/lib/constants'
import { cn } from '@/lib/utils'

import { Tooltip, TooltipContent, TooltipTrigger } from '@components/ui/tooltip'

type SidebarItemProps = {
  label: string
  href: ROUTES
}

const SIDEBAR_ICON = {
  [ROUTES.DASHBOARD]: Home,
  [ROUTES.ORDERS]: ShoppingBag,
  [ROUTES.DEMANDS]: Package,
  [ROUTES.PRODUCTS]: Salad,
  [ROUTES.CUSTOMERS]: Users,
  [ROUTES.SETTINGS]: Settings,
}

const SidebarItem = ({ label, href }: SidebarItemProps) => {
  const pathname = usePathname()
  const isActive = pathname === href
  const Icon = SIDEBAR_ICON[href]

  return (
    <Tooltip>
      <TooltipTrigger asChild>
        <Link
          href={href}
          scroll={false}
          className={cn(
            'flex h-9 w-9 items-center justify-center rounded-lg text-muted-foreground transition-colors hover:text-foreground md:h-8 md:w-8',
            isActive && 'bg-primary'
          )}
        >
          <Icon
            className={cn('h-5 w-5', isActive && 'text-primary-foreground')}
          />
          <span className="sr-only">{label}</span>
        </Link>
      </TooltipTrigger>
      <TooltipContent side="right">{label}</TooltipContent>
    </Tooltip>
  )
}

export default SidebarItem
