package util

import (
	"encoding/json"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func MarshalData(data any) ([]byte, error) {
	if msg, ok := data.(proto.Message); ok {
		return protojson.Marshal(msg)
	}

	return json.Marshal(data)
}
