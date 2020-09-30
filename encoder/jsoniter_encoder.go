package encoder

import (
	jsoniter "github.com/json-iterator/go"
)

type JsoniterEncoder struct {
	intEncoder *IntEncoder
}

func (j *JsoniterEncoder) Encode(data interface{}, writer IWriter) error {
	marshalled, err := jsoniter.Marshal(data)
	if err != nil {
		return err
	}
	if err := j.intEncoder.Encode(len(marshalled), writer); err != nil {
		return err
	}
	return writer.Write(marshalled)
}

func (j *JsoniterEncoder) Decode(reader IReader) (interface{}, error) {
	length, err := j.intEncoder.Decode(reader)
	if err != nil {
		return nil, err
	}
	var result interface{}
	data, err := reader.Read(length.(int))
	if err := jsoniter.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func NewJsoniterEncoder() *JsoniterEncoder {
	return &JsoniterEncoder{
		intEncoder: NewIntEncoder(),
	}
}
