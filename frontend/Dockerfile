FROM node:22-slim AS base

ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable
WORKDIR /app

FROM base AS deps
COPY /frontend/package.json /frontend/pnpm-lock.yaml* ./
RUN --mount=type=cache,id=pnpm,target=/pnpm/store pnpm install --frozen-lockfile

FROM base AS build
ENV NEXT_TELEMETRY_DISABLED=1
COPY --from=deps /app/node_modules ./node_modules
COPY /frontend ./frontend
COPY /specs ./specs
ARG PRODUCT_SERVICE_HOST
ENV PRODUCT_SERVICE_HOST $PRODUCT_SERVICE_HOST
RUN pnpm run build

FROM base AS run
ENV NEXT_TELEMETRY_DISABLED=1
ENV HOSTNAME=0.0.0.0
ENV PORT=3000
EXPOSE 3000
RUN addgroup --system --gid 1001 nodejs
RUN adduser --system --uid 1001 nextjs
COPY --from=build --chown=nextjs:nodejs /app/frontend/.next/standalone ./
COPY --from=build --chown=nextjs:nodejs /app/frontend/.next/static ./.next/static
COPY --from=build /app/frontend/public ./public
USER nextjs
CMD [ "node", "server.js" ]

FROM base AS dev
EXPOSE 3000
CMD [ "pnpm", "dev" ]