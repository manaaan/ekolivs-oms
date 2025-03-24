'use client'

import { MoonIcon, SunIcon } from 'lucide-react'
import { useTheme } from 'next-themes'

import { SidebarMenuButton } from '@components/ui/sidebar'

function SidebarMenuToggle() {
  const { setTheme } = useTheme()

  return (
    <SidebarMenuButton
      onClick={() => setTheme((theme) => (theme === 'dark' ? 'light' : 'dark'))}
    >
      <SunIcon className="scale-100 rotate-0 transition-all dark:scale-0 dark:-rotate-90" />
      <MoonIcon className="absolute scale-0 rotate-90 transition-all dark:scale-100 dark:rotate-0" />
      Toggle theme
    </SidebarMenuButton>
  )
}

export { SidebarMenuToggle }
