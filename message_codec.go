package gosmc

import (
	"sync"

	"github.com/firmanmm/gosmc/encoder"
)

var defaultEncoder *SimpleMessageCodec
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

//Automatically encode value into array of bytes.
//The order are determined by the encoder.
//Can only be decoded using [Decode] function.
func (s *SimpleMessageCodec) Encode(value interface{}) ([]byte, error) {
	writer := writerPool.Get().(encoder.IWriter)
	writer.Reset()
	defer writerPool.Put(writer)
	if err := s.valueEncoder.Encode(value, writer); err != nil {
		return nil, err
	}
	result, err := writer.GetContent()
	if err != nil {
		return nil, err
	}
	return result, nil
}

//Automatically decode value that is encoded using [Encode] function
func (s *SimpleMessageCodec) Decode(data []byte) (interface{}, error) {
	reader := encoder.NewSliceReader(data)
	return s.valueEncoder.Decode(reader)
}

type WriterFunc func(data ...interface{}) error

type ManualEncodeFunc func(data interface{}, write WriterFunc) error

//Manually encode a data
//You are responsible to determine the data order
//May gain performance increase if all encoded data are natively supported
//Useful if you want to encode arbitrary data
func (s *SimpleMessageCodec) ManualEncode(data interface{}, executor ManualEncodeFunc) ([]byte, error) {
	writer := writerPool.Get().(encoder.IWriter)
	writer.Reset()
	defer writerPool.Put(writer)
	if err := executor(data, func(inData ...interface{}) error {
		for _, val := range inData {
			if err := s.valueEncoder.Encode(val, writer); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}
	result, err := writer.GetContent()
	if err != nil {
		return nil, err
	}
	writerPool.Put(writer)
	return result, nil
}

type ManualReader struct {
	reader  encoder.IReader
	decoder *encoder.ValueEncoder
}

//Read one stored data
//Will return the data directly
func (m *ManualReader) Read() (interface{}, error) {
	return m.decoder.Decode(m.reader)
}

//Read N number of stored data
//Will return the data as array
func (m *ManualReader) ReadN(count int) ([]interface{}, error) {
	result := make([]interface{}, 0, count)
	for i := 0; i < count; i++ {
		res, err := m.decoder.Decode(m.reader)
		if err != nil {
			return nil, err
		}
		result = append(result, res)
	}
	return result, nil
}

type ManualDecodeFunc func(reader *ManualReader) (interface{}, error)

//Manually decode a data
//You are responsible to determine the data order
//Data order is the same as in the [ManualEncode]
//Useful if you want to decode arbitrary data
func (s *SimpleMessageCodec) ManualDecode(data []byte, executor ManualDecodeFunc) (interface{}, error) {
	reader := &ManualReader{
		reader:  encoder.NewSliceReader(data),
		decoder: s.valueEncoder,
	}
	return executor(reader)
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
			encoder.BoolValueEncoder:      encoder.NewBoolEncoder(),
		},
	)

	listEncoder := encoder.NewListInterfaceEncoder(valueEncoder)
	mapEncoder := encoder.NewMapCommonEncoder(valueEncoder)
	structEncoder := encoder.NewStructEncoder(valueEncoder)

	valueEncoder.SetEncoder(encoder.ListValueEncoder, listEncoder)
	valueEncoder.SetEncoder(encoder.MapValueEncoder, mapEncoder)
	valueEncoder.SetEncoder(encoder.StructValueEncoder, structEncoder)

	return &SimpleMessageCodec{
		valueEncoder: valueEncoder,
	}
}

//Creates new message codec but uses Jsoniter to handle struct and map
func NewSimpleMessageCodecWithJsoniter() *SimpleMessageCodec {
	current := NewSimpleMessageCodec()
	jsoniterEncoder := encoder.NewJsoniterEncoder()
	current.valueEncoder.SetEncoder(encoder.MapValueEncoder, jsoniterEncoder)
	current.valueEncoder.SetEncoder(encoder.StructValueEncoder, jsoniterEncoder)
	current.valueEncoder.SetEncoder(encoder.GeneralValueEncoder, jsoniterEncoder)
	return current
}

//Creates new message codec but uses universally known data type.
//This avoid inconsistency when dealing with langguage that doesn't have certain data type.
//Example scenario of this is that there is no `uint` data type in `Java`, `Javascript`, or even `Python`.
func NewUniversalSimpleMessageCodec() *SimpleMessageCodec {
	current := NewSimpleMessageCodec()
	current.valueEncoder.SetEncoder(encoder.UintValueEncoder, encoder.NewUintUniversalEncoder())
	return current
}

func init() {
	defaultEncoder = NewSimpleMessageCodec()
}

//Encode value with pure implementation encoder
func Encode(value interface{}) ([]byte, error) {
	return defaultEncoder.Encode(value)
}

//Decode value with pure implementation decoder
func Decode(data []byte) (interface{}, error) {
	return defaultEncoder.Decode(data)
}

//Perform manual encoding with pure implementation manual encoder
func ManualEncode(value interface{}, encoder ManualEncodeFunc) ([]byte, error) {
	return defaultEncoder.ManualEncode(value, encoder)
}

//Perform manual decoding with pure implementation manual decoder
func ManualDecode(data []byte, decoder ManualDecodeFunc) (interface{}, error) {
	return defaultEncoder.ManualDecode(data, decoder)
}
