# syntax=docker/dockerfile:1
#Create GO Lang application
FROM golang:1.17.6-bullseye as build

RUN mkdir -p /tmp/app/src
COPY . /tmp/app/src
WORKDIR /tmp/app/src

RUN go mod download
RUN go build -o /tmp/app/out/webapplication /tmp/app/src/cmd/webapplication/main.go

#Create VueJs frontend
FROM node:17-alpine AS frontend

RUN apk add yarn
WORKDIR /usr/src/app
COPY ./web ./

RUN yarn install && yarn run build && yarn cache clean && rm -rf node_modules

#Build main container
FROM debian:bullseye-slim

RUN apt-get -q update && apt-get install -qy supervisor procps tini wget netcat
RUN wget https://raw.githubusercontent.com/eficode/wait-for/master/wait-for -O /wait-for && chmod +x /wait-for
COPY build/package/webapplication/supervisor /etc/supervisor/

RUN mkdir /app

COPY --from=build /tmp/app/out/webapplication /app/webapplication
RUN chmod +x /app/webapplication

COPY --from=frontend /usr/src/app/dist /app/web/dist

COPY web/dist /app/web/dist
COPY web/static /app/web/static

COPY build/package/webapplication/entrypoint.sh /app/entrypoint.sh
RUN chmod +x /app/entrypoint.sh

ENTRYPOINT ["tini", "--", "/app/entrypoint.sh"]
