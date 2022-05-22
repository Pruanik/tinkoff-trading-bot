# syntax=docker/dockerfile:1
#Create GO Lang application
FROM golang:1.17.6-bullseye as build

RUN mkdir -p /tmp/app/src
COPY . /tmp/app/src
WORKDIR /tmp/app/src

RUN go mod download
RUN go build -o /tmp/app/out/tinkoffinvestconnection /tmp/app/src/cmd/tinkoffinvestconnection/main.go

#Build main container
FROM debian:bullseye-slim

RUN apt-get -q update && apt-get install -qy supervisor procps tini wget netcat
RUN wget https://raw.githubusercontent.com/eficode/wait-for/master/wait-for -O /wait-for && chmod +x /wait-for
COPY build/package/tinkoffinvestconnection/supervisor /etc/supervisor/

RUN mkdir /app

COPY --from=build /tmp/app/out/tinkoffinvestconnection /app/tinkoffinvestconnection
RUN chmod +x /app/tinkoffinvestconnection

COPY build/package/tinkoffinvestconnection/entrypoint.sh /app/entrypoint.sh
RUN chmod +x /app/entrypoint.sh

ENTRYPOINT ["tini", "--", "/app/entrypoint.sh"]
