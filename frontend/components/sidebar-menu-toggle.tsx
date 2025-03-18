'use client'

import { MoonIcon, SunIcon } from 'lucide-react'
import { useTheme } from 'next-themes'

import { SidebarMenuButton } from '@components/ui/sidebar'

function SidebarMenuToggle() {
  const { theme, setTheme } = useTheme()

  return (
    <SidebarMenuButton
      onClick={() => setTheme(theme === 'dark' ? 'light' : 'dark')}
    >
      {theme === 'dark' ? <MoonIcon /> : <SunIcon />}
      Toggle theme
    </SidebarMenuButton>
  )
}

export { SidebarMenuToggle }
