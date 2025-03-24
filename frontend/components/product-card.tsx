import Image from 'next/image'

import type { AppProduct } from '@/lib/services/product'
import { formatPrice } from '@/lib/utils'

import { Button } from '@components/ui/button'
import {
  Card,
  CardAction,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from '@components/ui/card'

interface ProductCardProps extends AppProduct {
  onAddToCart: () => void
}

function ProductCard({ name, price, imageUrl, onAddToCart }: ProductCardProps) {
  const imgUrl = imageUrl ?? 'https://dummyimage.com/329x192'

  return (
    <Card className="w-full overflow-hidden pt-0 md:w-80">
      <div className="relative h-48 w-full">
        <Image fill src={imgUrl} alt={name} className="object-cover" />
      </div>
      <CardHeader>
        <CardTitle>{name}</CardTitle>
        <CardDescription>
          {price?.amount && formatPrice(parseInt(price.amount, 10))}
        </CardDescription>
      </CardHeader>
      <CardContent className="mt-auto">
        <CardAction>
          <Button onClick={onAddToCart} className="cursor-pointer">
            Add to cart
          </Button>
        </CardAction>
      </CardContent>
    </Card>
  )
}

export { ProductCard }
