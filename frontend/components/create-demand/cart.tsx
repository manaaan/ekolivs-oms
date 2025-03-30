import { Fragment } from 'react'

import type { AppProduct } from '@/lib/services/product'

import { CartItem } from '@components/create-demand/cart-item'
import { Button } from '@components/ui/button'
import { ScrollArea } from '@components/ui/scroll-area'
import { Separator } from '@components/ui/separator'
import {
  SheetContent,
  SheetDescription,
  SheetFooter,
  SheetHeader,
  SheetTitle,
} from '@components/ui/sheet'

type CartItems = {
  [productId: AppProduct['id']]: {
    amount: number
    product: AppProduct
    position: number
  }
}

interface CartProps {
  cartItems: CartItems
  handleIncreaseAmount: (product: AppProduct) => () => void
  handleDecreaseAmount: (product: AppProduct) => () => void
}

function Cart({
  cartItems,
  handleIncreaseAmount,
  handleDecreaseAmount,
}: CartProps) {
  return (
    <SheetContent side="right" className="max-h-svh">
      <SheetHeader>
        <SheetTitle>Demand</SheetTitle>
        <SheetDescription>Have a look at your cart</SheetDescription>
      </SheetHeader>
      <ScrollArea className="flex-1 overflow-y-auto">
        <div className="px-4">
          {Object.values(cartItems)
            .toSorted((a, b) => a.position - b.position)
            .map(({ amount, product }) => (
              <Fragment key={product.id}>
                <CartItem
                  amount={amount}
                  onIncreaseAmount={handleIncreaseAmount(product)}
                  onDecreaseAmount={handleDecreaseAmount(product)}
                  {...product}
                />
                <Separator className="my-2" />
              </Fragment>
            ))}
        </div>
      </ScrollArea>
      <SheetFooter>
        <Button
          onClick={() => alert(JSON.stringify(cartItems, null, 2))}
          className="cursor-pointer"
        >
          Order
        </Button>
      </SheetFooter>
    </SheetContent>
  )
}

export { Cart }
