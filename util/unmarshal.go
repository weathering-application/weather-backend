package util

import (
	"encoding/json"
	"errors"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func UnmarshalData[T any](data []byte) (T, error) {
	var result T
	if msg, ok := any(&result).(proto.Message); ok {
		err := protojson.Unmarshal(data, msg)
		return result, err
	}

	err := json.Unmarshal(data, &result)
	if err != nil {
		return result, errors.New("failed to unmarshal data as both proto and JSON")
	}

	return result, nil
}
