'use client'

import type { ColumnDef } from '@tanstack/react-table'
import { ArrowUpDown, ExternalLinkIcon } from 'lucide-react'
import Link from 'next/link'

import type { Product } from '@/lib/services/product'

import { Button } from '@components/ui/button'

export const columns: ColumnDef<Product>[] = [
  {
    id: 'name',
    accessorKey: 'name',
    header: ({ column }) => (
      <Button
        variant="ghost"
        onClick={() => column.toggleSorting(column.getIsSorted() === 'asc')}
      >
        Name
        <ArrowUpDown className="ml-2 h-4 w-4" />
      </Button>
    ),
  },
  {
    id: 'price',
    accessorKey: 'price.amount',
    sortUndefined: -1,
    header: ({ column }) => (
      <Button
        variant="ghost"
        onClick={() => column.toggleSorting(column.getIsSorted() === 'asc')}
      >
        Price
        <ArrowUpDown className="ml-2 h-4 w-4" />
      </Button>
    ),
    cell: ({ row }) => {
      const price = parseFloat(row.getValue('price') ?? 0)

      const formatted = new Intl.NumberFormat('sv-SE', {
        style: 'currency',
        currency: 'SEK',
      }).format(price / 1000)

      return <div className="font-medium">{formatted}</div>
    },
  },
  {
    id: 'link',
    enableSorting: false,
    header: () => <span className="float-end">Link</span>,
    cell: ({ row }) => {
      const href = `/products/${row.original.ID}`
      return (
        <Link href={href} className="float-end">
          <ExternalLinkIcon />
        </Link>
      )
    },
  },
]
