import {
  Home,
  Package,
  Salad,
  Settings,
  ShoppingBag,
  Users,
} from 'lucide-react'

export enum ROUTES {
  DASHBOARD = '/dashboard',
  DEMANDS = '/dashboard/demands',
  ORDERS = '/dashboard/orders',
  PRODUCTS = '/dashboard/products',
  CUSTOMERS = '/dashboard/customers',
  SETTINGS = '/dashboard/settings',
}

export const ROUTES_ICON = {
  [ROUTES.DASHBOARD]: Home,
  [ROUTES.ORDERS]: ShoppingBag,
  [ROUTES.DEMANDS]: Package,
  [ROUTES.PRODUCTS]: Salad,
  [ROUTES.CUSTOMERS]: Users,
  [ROUTES.SETTINGS]: Settings,
}
