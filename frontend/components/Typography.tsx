import { Slot } from '@radix-ui/react-slot'
import { cva } from 'class-variance-authority'
import * as React from 'react'

import { cn } from '@/lib/utils'

type VariantMap = typeof variantMap

type TypographyProps<T extends keyof VariantMap = keyof VariantMap> =
  React.ComponentProps<VariantMap[T]> & {
    variant?: T
    asChild?: boolean
  }

const typographyVariants = cva('', {
  variants: {
    variant: {
      h1: 'scroll-m-20 text-4xl font-extrabold tracking-tight lg:text-5xl',
      h2: 'scroll-m-20 border-b text-3xl font-semibold tracking-tight',
      h3: 'scroll-m-20 text-2xl font-semibold tracking-tight',
      h4: 'scroll-m-20 text-xl font-semibold tracking-tight',
      p: 'leading-7',
      blockquote: 'mt-6 border-l-2 pl-6 italic',
      table:
        'text-left[&[align=center]]:text-center [&[align=right]]:text-right',
    },
  },
})

const variantMap = {
  h1: 'h1',
  h2: 'h2',
  h3: 'h3',
  h4: 'h4',
  p: 'p',
  blockquote: 'blockquote',
  table: 'p',
} as const

function Typography({
  variant = 'p',
  asChild,
  className,
  ...props
}: TypographyProps) {
  const Comp = asChild ? Slot : variantMap[variant]

  return (
    <Comp
      data-slot="typography"
      className={cn(typographyVariants({ variant, className }))}
      {...props}
    />
  )
}

export { Typography }
