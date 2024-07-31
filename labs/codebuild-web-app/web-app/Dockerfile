FROM node:20-alpine AS build
WORKDIR /app

COPY . .

RUN npm ci
RUN npm run build


FROM node:20-alpine AS production
WORKDIR /app

COPY --from=build /app/.next/standalone /app
COPY --from=build /app/public /app/public
COPY --from=build /app/.next/static /app/.next/static

EXPOSE 3000

CMD ["node", "server.js"]