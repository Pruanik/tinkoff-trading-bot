# syntax=docker/dockerfile:1
FROM golang:1.19.5-bullseye as build

ARG MIGRATE_PACKAGE_TAG=v4.15.2

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@$MIGRATE_PACKAGE_TAG

FROM debian:bullseye-slim

RUN apt-get -q update && apt-get install -qy make openssl ca-certificates wget netcat
RUN wget https://raw.githubusercontent.com/eficode/wait-for/master/wait-for -O /wait-for && chmod +x /wait-for

COPY --from=build /go/bin/migrate /migrate
RUN chmod +x /migrate
COPY /build/package/migrator/Makefile Makefile
COPY /migrations /migrations

ENTRYPOINT ["make"]
CMD ["help"]
