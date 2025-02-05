#!/bin/bash
# Generate Go code
protoc --go_out=. --go-grpc_out=. weather.proto

# Generate Python code
# python -m grpc_tools.protoc -I. --python_out=../ml_service/services  --grpc_python_out=../ml_service/services ml.proto