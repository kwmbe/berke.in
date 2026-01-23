FROM oven/bun:1 AS builder

WORKDIR /app

COPY package.json bun.lockb* ./
RUN bun install --frozen-lockfile

COPY . .
RUN bun run build

FROM alpine:latest AS expose

WORKDIR /shared

COPY --from=builder /app/dist ./
