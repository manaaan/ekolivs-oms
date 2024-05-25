import { SidebarNav } from '@/components/SidebarNav'
import { Sheet } from '@/components/ui/sheet'

const DashboardLayout = ({ children }: { children: React.ReactNode }) => {
  return (
    <>
      <Sheet>
        <SidebarNav />
        {children}
      </Sheet>
    </>
  )
}

export default DashboardLayout
