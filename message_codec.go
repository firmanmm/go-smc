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
	return s.valueEncoder.Encode(value)
}

func (s *SimpleMessageCodec) Decode(data []byte) (interface{}, error) {
	return s.valueEncoder.Decode(data)
}

func NewSimpleMessageCodec() *SimpleMessageCodec {

	valueEncoder := encoder.NewValueEncoder(
		map[encoder.ValueEncoderType]encoder.IValueEncoderUnit{
			encoder.ByteArrayValueEncoder: encoder.NewByteArrayEncoder(),
			encoder.FloatValueEncoder:     encoder.NewFloatEncoder(),
			encoder.IntValueEncoder:       encoder.NewIntEncoder(),
			encoder.StringValueEncoder:    encoder.NewStringEncoder(),
			encoder.UintValueEncoder:      encoder.NewUintEncoder(),
		},
	)

	listEncoder := encoder.NewListEncoder(valueEncoder, encoder.NewUintEncoder())
	mapEncoder := encoder.NewMapEncoder(valueEncoder)

	valueEncoder.SetEncoder(encoder.ListValueEncoder, listEncoder)
	valueEncoder.SetEncoder(encoder.MapValueEncoder, mapEncoder)

	return &SimpleMessageCodec{
		valueEncoder: valueEncoder,
	}
}
