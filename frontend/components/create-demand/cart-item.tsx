import type { AppProduct } from '@/lib/services/product'

import { Button } from '@components/ui/button'

interface CartItemProps extends AppProduct {
  amount: number
  onIncreaseAmount: () => void
  onDecreaseAmount: () => void
}

function CartItem({
  name,
  amount,
  onIncreaseAmount,
  onDecreaseAmount,
}: CartItemProps) {
  return (
    <div className="flex items-center gap-4">
      <div className="grow">{name}</div>
      <div className="ring-muted-foreground flex flex-col rounded-md ring">
        <Button
          onClick={onIncreaseAmount}
          className="cursor-pointer rounded-b-none"
        >
          +
        </Button>
        <span className="text-center font-bold">{amount}x</span>
        <Button
          onClick={onDecreaseAmount}
          className="cursor-pointer rounded-t-none"
        >
          -
        </Button>
      </div>
    </div>
  )
}

export { CartItem }
