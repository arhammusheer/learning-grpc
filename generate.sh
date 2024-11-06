#!/bin/bash

# Clean the existing generated code by removing only .pb.go files
find proto -type f -name "*.pb.go" -delete

# Loop through each .proto file in the proto directory
for proto_file in proto/*.proto; do
  # Get the base name of the file, e.g., "user" for "proto/user.proto"
  base_name=$(basename "$proto_file" .proto)
  
  # Create the output directory for the specific proto file
  mkdir -p "proto/$base_name"

  # Generate the gRPC code with protoc for each .proto file
  protoc --go_out=. --go_opt=M"$proto_file"="proto/$base_name" \
         --go-grpc_out=. --go-grpc_opt=M"$proto_file"="proto/$base_name" \
         "$proto_file"
done
