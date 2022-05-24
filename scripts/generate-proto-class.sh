#!/usr/bin/env bash

protoc -I ./internal/infrastructure/grpc/proto/ ./internal/infrastructure/grpc/proto/common.proto --go-grpc_out=./internal/infrastructure/grpc/investapi && \
protoc -I ./internal/infrastructure/grpc/proto/ ./internal/infrastructure/grpc/proto/instruments.proto --go-grpc_out=./internal/infrastructure/grpc/investapi && \
protoc -I ./internal/infrastructure/grpc/proto/ ./internal/infrastructure/grpc/proto/marketdata.proto --go-grpc_out=./internal/infrastructure/grpc/investapi && \
protoc -I ./internal/infrastructure/grpc/proto/ ./internal/infrastructure/grpc/proto/operations.proto --go-grpc_out=./internal/infrastructure/grpc/investapi && \
protoc -I ./internal/infrastructure/grpc/proto/ ./internal/infrastructure/grpc/proto/orders.proto --go-grpc_out=./internal/infrastructure/grpc/investapi && \
protoc -I ./internal/infrastructure/grpc/proto/ ./internal/infrastructure/grpc/proto/sandbox.proto --go-grpc_out=./internal/infrastructure/grpc/investapi && \
protoc -I ./internal/infrastructure/grpc/proto/ ./internal/infrastructure/grpc/proto/stoporders.proto --go-grpc_out=./internal/infrastructure/grpc/investapi && \
protoc -I ./internal/infrastructure/grpc/proto/ ./internal/infrastructure/grpc/proto/users.proto --go-grpc_out=./internal/infrastructure/grpc/investapi && \
protoc -I=./internal/infrastructure/grpc/proto/ --go_out=./internal/infrastructure/grpc/investapi ./internal/infrastructure/grpc/proto/common.proto && \
protoc -I=./internal/infrastructure/grpc/proto/ --go_out=./internal/infrastructure/grpc/investapi ./internal/infrastructure/grpc/proto/instruments.proto && \
protoc -I=./internal/infrastructure/grpc/proto/ --go_out=./internal/infrastructure/grpc/investapi ./internal/infrastructure/grpc/proto/marketdata.proto  && \
protoc -I=./internal/infrastructure/grpc/proto/ --go_out=./internal/infrastructure/grpc/investapi ./internal/infrastructure/grpc/proto/operations.proto && \
protoc -I=./internal/infrastructure/grpc/proto/ --go_out=./internal/infrastructure/grpc/investapi ./internal/infrastructure/grpc/proto/orders.proto     && \
protoc -I=./internal/infrastructure/grpc/proto/ --go_out=./internal/infrastructure/grpc/investapi ./internal/infrastructure/grpc/proto/sandbox.proto && \
protoc -I=./internal/infrastructure/grpc/proto/ --go_out=./internal/infrastructure/grpc/investapi ./internal/infrastructure/grpc/proto/stoporders.proto && \
protoc -I=./internal/infrastructure/grpc/proto/ --go_out=./internal/infrastructure/grpc/investapi ./internal/infrastructure/grpc/proto/users.proto     