FROM node:17-alpine
RUN apk add yarn
WORKDIR /usr/src/app
COPY ./web ./
COPY build/package/frontend/Makefile /usr/src/app/Makefile
RUN yarn install && yarn run build && yarn cache clean && rm -rf node_modules
CMD ["yarn"]