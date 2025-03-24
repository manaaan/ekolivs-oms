'use client'

import { ShoppingCartIcon } from 'lucide-react'
import { useState } from 'react'

import type { AppProduct } from '@/lib/services/product'

import { Cart } from '@components/create-demand/cart'
import { ProductCard } from '@components/product-card'
import { Input } from '@components/ui/input'
import { Sheet, SheetTrigger } from '@components/ui/sheet'

interface ProductMasonryProps {
  products: AppProduct[]
}

type ProductCart = {
  [productId: AppProduct['id']]: {
    amount: number
    product: AppProduct
    position: number
  }
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

      <div className="flex flex-wrap gap-4">
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
      <Sheet open={isCartOpen} onOpenChange={setIsCartOpen}>
        <SheetTrigger className="bg-accent fixed right-4 bottom-4 cursor-pointer rounded-full border p-5 shadow transition-shadow hover:shadow-lg">
          <ShoppingCartIcon className="size-6" />
        </SheetTrigger>
        <Cart
          cartItems={cartItems}
          handleIncreaseAmount={handleIncreaseAmount}
          handleDecreaseAmount={handleDecreaseAmount}
        />
      </Sheet>
    </div>
  )
}

export { ProductMasonry }
