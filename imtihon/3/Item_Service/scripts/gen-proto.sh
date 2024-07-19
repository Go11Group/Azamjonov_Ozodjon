#!/bin/bash

# Ensure the script exits if any command fails
set -e

# Get the directory of the proto files
PROTO_DIR="$1/grpc-proto"

# Path to the generated Go code
GO_OUT_DIR="$1/generated"

# Ensure the output directory exists
mkdir -p "$GO_OUT_DIR"

# Generate Go code from proto files
protoc --proto_path="$PROTO_DIR" --go_out="$GO_OUT_DIR" --go-grpc_out="$GO_OUT_DIR" "$PROTO_DIR/item/item.proto"

echo "Proto files generated successfully."
