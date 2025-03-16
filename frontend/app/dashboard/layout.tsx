import { AppSidebar } from '@components/app-sidebar'
import { SidebarProvider } from '@components/ui/sidebar'

const DashboardLayout = ({ children }: { children: React.ReactNode }) => {
  return (
    <SidebarProvider>
      <AppSidebar />
      <div className="flex-1">{children}</div>
    </SidebarProvider>
  )
}

export default DashboardLayout
