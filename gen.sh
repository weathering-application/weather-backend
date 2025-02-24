PROTO_DIR=proto
# Generate Go code into the generated directory
protoc --go_out=${PROTO_DIR} --go-grpc_out=${PROTO_DIR} ${PROTO_DIR}/weather.proto
