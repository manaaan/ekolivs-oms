import Link from 'next/link'

import { ROUTES_ICON } from '@/lib/constants'

import { NavItemProps } from './types'

const SheetItem = ({ label, href }: NavItemProps) => {
  const Icon = ROUTES_ICON[href]

  return (
    <Link
      href={href}
      scroll={false}
      className="flex items-center gap-4 px-2.5 text-muted-foreground hover:text-foreground"
    >
      <Icon className={'h-5 w-5'} />
      {label}
    </Link>
  )
}

export default SheetItem
