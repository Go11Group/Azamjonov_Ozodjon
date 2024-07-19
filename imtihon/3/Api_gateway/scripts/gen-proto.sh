#!/bin/bash

# Ensure the script exits if any command fails
set -e

# Get the base directory containing the proto directories
BASE_DIR="$1/grpc-proto"

# Path to the generated Go code
GO_OUT_DIR="$1/generated"

# Ensure the output directory exists
mkdir -p "$GO_OUT_DIR"

# Array of directories containing proto files
PROTO_DIRS=("auth" "item")

# Loop through each proto directory and generate Go code
for DIR in "${PROTO_DIRS[@]}"; do
  PROTO_PATH="$BASE_DIR/$DIR"
  protoc --proto_path="$PROTO_PATH" --go_out="$GO_OUT_DIR" --go-grpc_out="$GO_OUT_DIR" "$PROTO_PATH"/*.proto
done

echo "Proto files generated successfully."
