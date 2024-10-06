import { Metadata, MetadataValue } from '@grpc/grpc-js'
import { Status } from '@grpc/grpc-js/build/src/constants'

export class GrpcError extends Error {
  code: string
  details: string
  metadata: Record<string, MetadataValue>

  constructor(
    message: string,
    code: Status,
    details: string,
    metadata: Metadata,
    cause?: unknown
  ) {
    super(message, { cause })
    this.code = Status[code]
    this.details = details
    this.metadata = metadata.getMap()
  }
}
