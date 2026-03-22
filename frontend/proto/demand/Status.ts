// Original file: ../specs/demand.proto

export const Status = {
  RECEIVED: 'RECEIVED',
  ACCEPTED: 'ACCEPTED',
  REJECTED: 'REJECTED',
  IN_PROGRESS: 'IN_PROGRESS',
  FULFILLED: 'FULFILLED',
  CONCLUDED: 'CONCLUDED',
} as const;

export type Status =
  | 'RECEIVED'
  | 0
  | 'ACCEPTED'
  | 1
  | 'REJECTED'
  | 2
  | 'IN_PROGRESS'
  | 3
  | 'FULFILLED'
  | 4
  | 'CONCLUDED'
  | 5

export type Status__Output = typeof Status[keyof typeof Status]
