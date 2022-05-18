# syntax=docker/dockerfile:1
FROM golang:1.17.6-bullseye as build

ARG MIGRATE_PACKAGE_TAG=v4.15.2

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@$MIGRATE_PACKAGE_TAG

FROM debian:bullseye-slim

RUN apt-get update && apt-get install -q make openssl ca-certificates

COPY --from=build /go/bin/migrate /migrate
RUN chmod +x /migrate
COPY /build/package/migrator/Makefile Makefile
COPY /migrations /migrations

ENTRYPOINT ["make"]
CMD ["help"]
