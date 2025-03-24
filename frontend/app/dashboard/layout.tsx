import { AppSidebar } from '@components/app-sidebar'
import { SidebarProvider } from '@components/ui/sidebar'

const DashboardLayout = ({ children }: { children: React.ReactNode }) => {
  return (
    <SidebarProvider>
      <AppSidebar />
      {children}
    </SidebarProvider>
  )
}

export default DashboardLayout
