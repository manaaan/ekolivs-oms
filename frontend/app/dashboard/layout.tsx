import { SidebarNav } from '@/components/SidebarNav'

const DashboardLayout = ({ children }: { children: React.ReactNode }) => {
  return (
    <>
      <SidebarNav />
      {children}
    </>
  )
}

export default DashboardLayout
