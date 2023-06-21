#!/bin/sh
PATH="$PATH:$(go env GOPATH)/bin"

PROTO_PATH="/Users/igoryarkov/projects/microserver/examples/admin/api:/Users/igoryarkov/projects/microserver/foundation/protobuf"
TARGET=generated

rm -f ${TARGET}/*.pb.go

protoc --go_out=. \
    --go_opt=Mgroup.proto=./${TARGET} \
    --go_opt=Mstandard.proto=./${TARGET} \
    --proto_path=${PROTO_PATH} \
    standard.proto

protoc --go_out=. \
    --go_opt=Mgroup.proto=./${TARGET} \
    --go_opt=Mstandard.proto=./${TARGET} \
    --proto_path=${PROTO_PATH} \
    --go-grpc_out=${TARGET} \
    --go-grpc_opt=paths=source_relative \
    --go-grpc_opt=Mgroup.proto=./${TARGET} \
    --go-grpc_opt=Mstandard.proto=./${TARGET} \
    group.proto
