package encoder

import (
	jsoniter "github.com/json-iterator/go"
)

type JsoniterEncoder struct {
}

func (j *JsoniterEncoder) Encode(data interface{}, tracker *BufferTracker) ([]byte, error) {
	return jsoniter.Marshal(data)
}

func (j *JsoniterEncoder) Decode(data []byte) (interface{}, error) {
	var result interface{}
	if err := jsoniter.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func NewJsoniterEncoder() *JsoniterEncoder {
	return &JsoniterEncoder{}
}
