// Original file: ../specs/product.proto

export const UnitType = {
  PIECES: 'PIECES',
  GRAMS: 'GRAMS',
  KILOGRAMS: 'KILOGRAMS',
  LITER: 'LITER',
  MILLILITER: 'MILLILITER',
} as const;

export type UnitType =
  | 'PIECES'
  | 0
  | 'GRAMS'
  | 1
  | 'KILOGRAMS'
  | 2
  | 'LITER'
  | 3
  | 'MILLILITER'
  | 4

export type UnitType__Output = typeof UnitType[keyof typeof UnitType]
