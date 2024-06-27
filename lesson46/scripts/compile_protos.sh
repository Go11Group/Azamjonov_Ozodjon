#!/bin/bash

CURRENT_DIR=$1

# Remove old generated proto files
rm -rf ${CURRENT_DIR}/genproto

# Create the output directory for generated proto files
mkdir -p ${CURRENT_DIR}/genproto

# Compile proto files
for x in $(find ${CURRENT_DIR}/protos -type d); do
  protoc -I=${x} -I=${CURRENT_DIR}/protos -I/usr/local/go \
    --go_out=${CURRENT_DIR}/genproto \
    --go-grpc_out=${CURRENT_DIR}/genproto \
    ${x}/*.proto
done

echo "Proto files compiled successfully."
