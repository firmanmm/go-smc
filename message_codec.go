package gosmc

import "github.com/firmanmm/gosmc/encoder"

type IMessageCodec interface {
	Encode(interface{}) ([]byte, error)
	Decode([]byte) (interface{}, error)
}

type SimpleMessageCodec struct {
	valueEncoder *encoder.ValueEncoder
}

func (s *SimpleMessageCodec) Encode(value interface{}) ([]byte, error) {
	return nil, nil
}

func (s *SimpleMessageCodec) Decode([]byte) (interface{}, error) {
	return nil, nil
}

func NewSimpleMessageCodec() *SimpleMessageCodec {
	return &SimpleMessageCodec{}
}
