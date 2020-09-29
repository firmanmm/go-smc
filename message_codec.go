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
	tracker := encoder.GetBufferTracker()
	defer encoder.PutBufferTracker(tracker)
	return s.valueEncoder.Encode(value, tracker)
}

func (s *SimpleMessageCodec) Decode(data []byte) (interface{}, error) {
	return s.valueEncoder.Decode(data)
}

//Creates new message codec with default encoder
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

//Creates new message codec but uses Jsoniter to handle struct and map
func NewSimpleMessageCodecWithJsoniter() *SimpleMessageCodec {
	current := NewSimpleMessageCodec()
	jsoniterEncoder := encoder.NewJsoniterEncoder()
	current.valueEncoder.SetEncoder(encoder.MapValueEncoder, jsoniterEncoder)
	current.valueEncoder.SetEncoder(encoder.GeneralValueEncoder, jsoniterEncoder)
	return current
}
