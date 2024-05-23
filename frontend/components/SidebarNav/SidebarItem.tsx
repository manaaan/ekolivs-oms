'use client'

import dynamicIconImports from 'lucide-react/dynamicIconImports'
import dynamic from 'next/dynamic'
import Link from 'next/link'
import { usePathname } from 'next/navigation'

import { ROUTES } from '@/lib/constants'
import { cn } from '@/lib/utils'

import { Tooltip, TooltipContent, TooltipTrigger } from '@components/ui/tooltip'

type SidebarItemProps = {
  label: string
  href: ROUTES
  iconName: keyof typeof dynamicIconImports
}

const SidebarItem = ({ label, href, iconName }: SidebarItemProps) => {
  const pathname = usePathname()
  const isActive = pathname === href

  const LucideIcon = dynamic(dynamicIconImports[iconName])

  return (
    <Tooltip>
      <TooltipTrigger asChild>
        <Link
          href={href}
          className={cn(
            'flex h-9 w-9 items-center justify-center rounded-lg text-muted-foreground transition-colors hover:text-foreground md:h-8 md:w-8',
            isActive && 'bg-primary'
          )}
        >
          <LucideIcon className={cn('h-5 w-5', isActive && 'text-white')} />
          <span className="sr-only">{label}</span>
        </Link>
      </TooltipTrigger>
      <TooltipContent side="right">{label}</TooltipContent>
    </Tooltip>
  )
}

export default SidebarItem
