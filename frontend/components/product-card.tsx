import Image from 'next/image'

import type { Product } from '@/lib/services/product'
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

interface ProductCardProps extends Product {
  onAddToCart: () => void
}

function ProductCard({
  name,
  costPrice,
  imageUrl,
  onAddToCart,
}: ProductCardProps) {
  const imgUrl = imageUrl ?? 'https://dummyimage.com/329x192'

  return (
    <Card className="w-80 overflow-hidden pt-0">
      <Image
        priority
        src={imgUrl}
        alt={name}
        width={320}
        height={192}
        className="h-48 w-80 object-cover"
      />
      <CardHeader>
        <CardTitle>{name}</CardTitle>
        <CardDescription>
          {costPrice?.amount && formatPrice(parseInt(costPrice.amount, 10))}
        </CardDescription>
      </CardHeader>
      <CardContent>
        <CardAction>
          <Button onClick={onAddToCart}>Add to cart</Button>
        </CardAction>
      </CardContent>
    </Card>
  )
}

export { ProductCard }
