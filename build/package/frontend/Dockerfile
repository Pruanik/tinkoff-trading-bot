FROM node:18.2.0-alpine
RUN apk add yarn
WORKDIR /usr/src/app
COPY ./web ./
RUN yarn install && yarn run build && yarn cache clean && rm -rf node_modules
CMD ["yarn"]