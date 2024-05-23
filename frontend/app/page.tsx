import { Metadata } from 'next'
import Image from 'next/image'
import Link from 'next/link'

import { ModeToggle } from '@/components/ModeToggle'
import { UserAuthForm } from '@/components/UserAuthForm'
import { buttonVariants } from '@/components/ui/button'

import logoTextWhitePic from '@public/images/logo_text_white.png'

export const metadata: Metadata = {
  title: 'Login',
  description: 'Login to Ekolivs OMS',
}

const LoginPage = () => {
  return (
    <div className="relative grid min-h-screen items-center lg:grid-cols-2">
      <div className="absolute right-4 top-4 flex md:right-8 md:top-8">
        <Link
          href="/dashboard"
          className={buttonVariants({ variant: 'ghost' })}
        >
          Login
        </Link>
        <ModeToggle />
      </div>
      <div className="relative hidden h-full flex-col justify-end bg-zinc-900 p-10 text-white dark:border-r lg:flex">
        <div className="m-auto max-w-md xl:max-w-2xl">
          <Image priority src={logoTextWhitePic} alt="Logo in white" />
        </div>
        <blockquote className="space-y-2">
          <p className="text-lg">
            &ldquo;This service has saved me countless hours of work and helped
            me deliver stunning orders to my people faster than ever
            before.&rdquo;
          </p>
          <footer className="text-sm">Sofia Davis</footer>
        </blockquote>
      </div>
      <div className="p-8">
        <div className="mx-auto flex w-full flex-col justify-center space-y-6 sm:w-[350px]">
          <div className="flex flex-col space-y-2 text-center">
            <h1 className="text-2xl font-semibold tracking-tight">
              Create an account
            </h1>
            <p className="text-sm text-muted-foreground">
              Enter your email below to create your account
            </p>
          </div>
          <UserAuthForm />
          <p className="px-8 text-center text-sm text-muted-foreground">
            By clicking continue, you agree to our{' '}
            <Link
              href="/terms"
              className="underline underline-offset-4 hover:text-primary"
            >
              Terms of Service
            </Link>{' '}
            and{' '}
            <Link
              href="/privacy"
              className="underline underline-offset-4 hover:text-primary"
            >
              Privacy Policy
            </Link>
            .
          </p>
        </div>
      </div>
    </div>
  )
}

export default LoginPage
