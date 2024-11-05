#!/bin/bash
# Make directory for generated Go code
mkdir -p proto/user
mkdir -p proto/message

# Generate Go code for user.proto in proto/user
protoc --go_out=./proto/user --go_opt=paths=source_relative \
       --go-grpc_out=./proto/user --go-grpc_opt=paths=source_relative \
       proto/user.proto

# Generate Go code for message.proto in proto/message
protoc --go_out=./proto/message --go_opt=paths=source_relative \
       --go-grpc_out=./proto/message --go-grpc_opt=paths=source_relative \
       proto/message.proto
