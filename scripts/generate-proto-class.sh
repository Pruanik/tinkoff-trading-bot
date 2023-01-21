#!/usr/bin/env bash

wget https://raw.githubusercontent.com/Tinkoff/investAPI/main/src/docs/contracts/common.proto -O ./internal/infrastructure/grpc/proto/common.proto
wget https://raw.githubusercontent.com/Tinkoff/investAPI/main/src/docs/contracts/instruments.proto -O ./internal/infrastructure/grpc/proto/instruments.proto
wget https://raw.githubusercontent.com/Tinkoff/investAPI/main/src/docs/contracts/marketdata.proto -O ./internal/infrastructure/grpc/proto/marketdata.proto
wget https://raw.githubusercontent.com/Tinkoff/investAPI/main/src/docs/contracts/operations.proto -O ./internal/infrastructure/grpc/proto/operations.proto
wget https://raw.githubusercontent.com/Tinkoff/investAPI/main/src/docs/contracts/orders.proto -O ./internal/infrastructure/grpc/proto/orders.proto
wget https://raw.githubusercontent.com/Tinkoff/investAPI/main/src/docs/contracts/sandbox.proto -O ./internal/infrastructure/grpc/proto/sandbox.proto
wget https://raw.githubusercontent.com/Tinkoff/investAPI/main/src/docs/contracts/stoporders.proto -O ./internal/infrastructure/grpc/proto/stoporders.proto
wget https://raw.githubusercontent.com/Tinkoff/investAPI/main/src/docs/contracts/users.proto -O ./internal/infrastructure/grpc/proto/users.proto

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