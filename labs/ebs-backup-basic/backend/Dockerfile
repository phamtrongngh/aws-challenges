FROM node:20-alpine AS build

WORKDIR /app

COPY --chown=node:node . .
RUN npm ci

RUN npm run build
RUN npm prune --production

USER node


FROM node:20-alpine AS production

ENV NODE_ENV=production
WORKDIR /app

COPY --from=build /app/dist ./dist
COPY --from=build /app/node_modules ./node_modules

USER node

CMD ["node", "dist/main.js"]