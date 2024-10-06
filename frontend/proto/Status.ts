// Original file: ../backend/services/product/api/service.proto

export const Status = {
  ACTIVE: 'ACTIVE',
  HIDDEN: 'HIDDEN',
} as const

export type Status = 'ACTIVE' | 0 | 'HIDDEN' | 1

export type Status__Output = (typeof Status)[keyof typeof Status]
