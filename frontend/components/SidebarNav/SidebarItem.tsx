'use client'

import Link from 'next/link'
import { usePathname } from 'next/navigation'

import { ROUTES_ICON } from '@/lib/constants'
import { cn } from '@/lib/utils'

import { Tooltip, TooltipContent, TooltipTrigger } from '@components/ui/tooltip'

import { NavItemProps } from './types'

const SidebarItem = ({ label, href }: NavItemProps) => {
  const pathname = usePathname()
  const isActive = pathname === href
  const Icon = ROUTES_ICON[href]

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
