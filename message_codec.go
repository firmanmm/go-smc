package gosmc

import (
	"github.com/firmanmm/gosmc/encoder"
)

type IMessageCodec interface {
	Encode(interface{}) ([]byte, error)
	Decode([]byte) (interface{}, error)
}

type SimpleMessageCodec struct {
	valueEncoder *encoder.LinkedValueEncoder
}

func (s *SimpleMessageCodec) Encode(value interface{}) ([]byte, error) {
	result, err := s.valueEncoder.Encode(value)
	if err != nil {
		return nil, err
	}
	return result.GetResult(), nil

}

func (s *SimpleMessageCodec) Decode(data []byte) (interface{}, error) {
	return s.valueEncoder.Decode(data)
}

func NewSimpleMessageCodec() *SimpleMessageCodec {

	byteArrayEncoder := encoder.NewNativeLinkedEncoderUnitAdapter(
		encoder.NewByteArrayEncoder(),
	)

	floatEncoder := encoder.NewNativeLinkedEncoderUnitAdapter(
		encoder.NewFloatEncoder(),
	)

	intEncoder := encoder.NewNativeLinkedEncoderUnitAdapter(
		encoder.NewIntEncoder(),
	)

	stringEncoder := encoder.NewNativeLinkedEncoderUnitAdapter(
		encoder.NewStringEncoder(),
	)

	uintEncoder := encoder.NewNativeLinkedEncoderUnitAdapter(
		encoder.NewUintEncoder(),
	)

	linkedValueEncoder := encoder.NewLinkedValueEncoder(
		map[encoder.ValueEncoderType]encoder.IValueEncoderLinkedUnit{
			encoder.ByteArrayValueEncoder: byteArrayEncoder,
			encoder.FloatValueEncoder:     floatEncoder,
			encoder.IntValueEncoder:       intEncoder,
			encoder.StringValueEncoder:    stringEncoder,
			encoder.UintValueEncoder:      uintEncoder,
		})

	listLinkedEncoder := encoder.NewLinkedListEncoder(linkedValueEncoder, encoder.NewUintEncoder()) //TODO
	mapLinkedEncoder := encoder.NewLinkedMapEncoder(linkedValueEncoder, encoder.NewUintEncoder())   //TODO

	linkedValueEncoder.SetEncoder(encoder.ListValueEncoder, listLinkedEncoder)
	linkedValueEncoder.SetEncoder(encoder.MapValueEncoder, mapLinkedEncoder)

	return &SimpleMessageCodec{
		valueEncoder: linkedValueEncoder,
	}
}

func NewSimpleMessageCodecWithJsoniter() *SimpleMessageCodec {
	current := NewSimpleMessageCodec()
	jsoniterEncoder := encoder.NewNativeLinkedEncoderUnitAdapter(
		encoder.NewJsoniterEncoder(),
	)
	current.valueEncoder.SetEncoder(encoder.MapValueEncoder, jsoniterEncoder)
	current.valueEncoder.SetEncoder(encoder.GeneralValueEncoder, jsoniterEncoder)
	return current
}
