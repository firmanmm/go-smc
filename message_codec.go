package gosmc

import (
	"sync"

	"github.com/firmanmm/gosmc/encoder"
)

var writerPool *sync.Pool

func init() {
	writerPool = &sync.Pool{
		New: func() interface{} {
			return encoder.NewBufferWriter()
		},
	}
}

type IMessageCodec interface {
	Encode(interface{}) ([]byte, error)
	Decode([]byte) (interface{}, error)
}

type SimpleMessageCodec struct {
	valueEncoder *encoder.ValueEncoder
}

func (s *SimpleMessageCodec) Encode(value interface{}) ([]byte, error) {
	writer := writerPool.Get().(encoder.IWriter)
	if err := s.valueEncoder.Encode(value, writer); err != nil {
		return nil, err
	}
	result, err := writer.GetContent()
	if err != nil {
		return nil, err
	}
	writerPool.Put(writer)
	return result, nil
}

func (s *SimpleMessageCodec) Decode(data []byte) (interface{}, error) {
	reader := encoder.NewSliceReader(data)
	return s.valueEncoder.Decode(reader)
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

	listEncoder := encoder.NewListEncoder(valueEncoder)
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
