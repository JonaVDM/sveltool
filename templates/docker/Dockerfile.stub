FROM node:19-alpine as builder
WORKDIR /app
RUN npm i -g pnpm
COPY package.json pnpm-lock.yaml ./
RUN pnpm i --frozen-lockfile
COPY . .
RUN pnpm build

FROM node:19-alpine AS runner
WORKDIR /app
COPY --from=builder /app/build/ .
CMD ["node", "."]
