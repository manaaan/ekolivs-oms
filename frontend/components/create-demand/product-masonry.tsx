'use client'

import { ShoppingCartIcon } from 'lucide-react'
import { useState } from 'react'

import type { AppProduct } from '@/lib/services/product'

import { ProductCard } from '@components/product-card'
import { Button } from '@components/ui/button'
import { Input } from '@components/ui/input'
import { Separator } from '@components/ui/separator'
import {
  Sheet,
  SheetContent,
  SheetFooter,
  SheetHeader,
  SheetTitle,
  SheetTrigger,
} from '@components/ui/sheet'

interface ProductMasonryProps {
  products: AppProduct[]
}

interface ProductSheetProps extends AppProduct {
  amount: number
  onIncreaseAmount: () => void
  onDecreaseAmount: () => void
}

type ProductCart = {
  [productId: AppProduct['id']]: {
    amount: number
    product: AppProduct
    position: number
  }
}

function CartItem({
  name,
  amount,
  onIncreaseAmount,
  onDecreaseAmount,
}: ProductSheetProps) {
  return (
    <>
      <div className="flex flex-wrap items-center gap-2 md:flex-nowrap">
        <p className="grow">{name}</p>
        <div className="flex items-center justify-evenly space-x-2">
          <Button onClick={onIncreaseAmount} className="cursor-pointer">
            +
          </Button>
          <p className="font-bold">{amount}x</p>
          <Button onClick={onDecreaseAmount} className="cursor-pointer">
            -
          </Button>
        </div>
      </div>
      <Separator />
    </>
  )
}

function getHighestPosition(productsInCart: ProductCart) {
  return Math.max(...Object.values(productsInCart).map((p) => p.position))
}

function removeFromCart(cart: ProductCart, product: AppProduct) {
  const cartItem = cart[product.id]
  if (!cartItem) {
    return cart
  }

  if (cartItem.amount === 1) {
    delete cart[product.id]
    return { ...cart }
  }

  return {
    ...cart,
    [product.id]: {
      ...cartItem,
      amount: cartItem.amount - 1,
    },
  }
}

function addToCart(cart: ProductCart, product: AppProduct) {
  if (Object.keys(cart).length === 0) {
    return {
      [product.id]: {
        amount: 1,
        position: 0,
        product,
      },
    }
  }

  const cartItem = cart[product.id]
  if (!cartItem) {
    return {
      ...cart,
      [product.id]: {
        product,
        amount: 1,
        position: getHighestPosition(cart) + 1,
      },
    }
  }

  return {
    ...cart,
    [product.id]: {
      product,
      amount: cartItem.amount + 1,
      position: cartItem.position,
    },
  }
}

function ProductMasonry({ products }: ProductMasonryProps) {
  const [isCartOpen, setIsCartOpen] = useState(false)
  const [cartItems, setCartItems] = useState<ProductCart>({})
  const [filter, setFilter] = useState('')

  const handleAddToCart = (product: AppProduct) => () => {
    setCartItems((productsInCart) => addToCart(productsInCart, product))
    setIsCartOpen(true)
  }

  const handleIncreaseAmount = (product: AppProduct) => () => {
    setCartItems((productsInCart) => addToCart(productsInCart, product))
  }
  const handleDecreaseAmount = (product: AppProduct) => () => {
    setCartItems((productsInCart) => removeFromCart(productsInCart, product))
  }

  return (
    <div className="flex flex-col gap-4">
      <Input
        placeholder="Filter products..."
        value={filter}
        onChange={(event) => setFilter(event.target.value)}
        className="max-w-sm"
      />

      <Sheet open={isCartOpen} onOpenChange={setIsCartOpen}>
        <div className="flex flex-wrap justify-center gap-4">
          {products
            .filter(({ name }) =>
              name.toLowerCase().includes(filter.toLowerCase())
            )
            .slice(0, 12)
            .map((product) => {
              return (
                <ProductCard
                  key={product.id}
                  onAddToCart={handleAddToCart(product)}
                  {...product}
                />
              )
            })}
        </div>
        <SheetTrigger className="bg-accent fixed right-4 bottom-4 cursor-pointer rounded-full border p-5 shadow transition-shadow hover:shadow-lg">
          <ShoppingCartIcon className="size-6" />
        </SheetTrigger>

        <SheetContent side="right">
          <SheetHeader>
            <SheetTitle>Demand</SheetTitle>
            {Object.values(cartItems)
              .toSorted((a, b) => a.position - b.position)
              .map(({ amount, product }) => (
                <CartItem
                  key={product.id}
                  amount={amount}
                  onIncreaseAmount={handleIncreaseAmount(product)}
                  onDecreaseAmount={handleDecreaseAmount(product)}
                  {...product}
                />
              ))}
          </SheetHeader>
          <SheetFooter>
            <Button
              onClick={() => alert(JSON.stringify(cartItems, null, 2))}
              className="cursor-pointer"
            >
              Order
            </Button>
          </SheetFooter>
        </SheetContent>
      </Sheet>
    </div>
  )
}

export { ProductMasonry }
