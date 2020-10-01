package encoder

import (
	"errors"
	"fmt"
	"reflect"
)

type IValueEncoderUnit interface {
	Encode(interface{}, IWriter) error
	Decode(IReader) (interface{}, error)
}

type ValueEncoderType int

const (
	ByteValueEncoder      ValueEncoderType = 1 //Not Implemented
	IntValueEncoder       ValueEncoderType = 2
	UintValueEncoder      ValueEncoderType = 3
	FloatValueEncoder     ValueEncoderType = 4
	ComplexValueEncoder   ValueEncoderType = 5 //Not Implemented
	ByteArrayValueEncoder ValueEncoderType = 6
	StringValueEncoder    ValueEncoderType = 7
	ListValueEncoder      ValueEncoderType = 8
	MapValueEncoder       ValueEncoderType = 9
	BoolValueEncoder      ValueEncoderType = 10
	StructValueEncoder    ValueEncoderType = 11
	GeneralValueEncoder   ValueEncoderType = 255
)

type ValueEncoder struct {
	encoders map[ValueEncoderType]IValueEncoderUnit
}

func (s *ValueEncoder) Encode(data interface{}, writer IWriter) error {
	encoderUsed := ValueEncoderType(0)
	switch data.(type) {
	case bool:
		encoderUsed = BoolValueEncoder
	case int8:
		data = int(data.(int8))
		encoderUsed = IntValueEncoder
	case int16:
		data = int(data.(int16))
		encoderUsed = IntValueEncoder
	case int32:
		data = int(data.(int32))
		encoderUsed = IntValueEncoder
	case int64:
		data = int(data.(int64))
		encoderUsed = IntValueEncoder
	case int:
		encoderUsed = IntValueEncoder
	case uint8:
		data = uint(data.(uint8))
		encoderUsed = UintValueEncoder
	case uint16:
		data = uint(data.(uint16))
		encoderUsed = UintValueEncoder
	case uint32:
		data = uint(data.(uint32))
		encoderUsed = UintValueEncoder
	case uint64:
		data = uint(data.(uint64))
		encoderUsed = UintValueEncoder
	case uint:
		encoderUsed = UintValueEncoder
	case float32:
		data = float64(data.(float32))
		encoderUsed = FloatValueEncoder
	case float64:
		encoderUsed = FloatValueEncoder
	case []byte:
		encoderUsed = ByteArrayValueEncoder
	case string:
		encoderUsed = StringValueEncoder
	default:
		reflected := reflect.ValueOf(data)
		switch reflected.Kind() {
		case reflect.Slice:
			encoderUsed = ListValueEncoder
		case reflect.Map:
			encoderUsed = MapValueEncoder
		case reflect.Struct:
			encoderUsed = StructValueEncoder
		case reflect.Ptr:
			return s.Encode(reflected.Elem().Interface(), writer)
		default:
			_, ok := s.encoders[GeneralValueEncoder]
			if !ok {
				return errors.New("Unsupported type, try to register fallback encoder")
			}
			encoderUsed = GeneralValueEncoder
		}

	}
	return s.encode(encoderUsed, data, writer)
}

func (v *ValueEncoder) encode(dataType ValueEncoderType, data interface{}, writer IWriter) error {
	dataEncoder, ok := v.encoders[dataType]
	if !ok {
		return fmt.Errorf("Data Type encoder not registered for data %v", data)
	}
	if err := writer.WriteByte(byte(dataType)); err != nil {
		return err
	}
	if err := dataEncoder.Encode(data, writer); err != nil {
		return err
	}
	return nil
}

func (v *ValueEncoder) Decode(reader IReader) (interface{}, error) {
	decoderType, err := reader.ReadByte()
	if err != nil {
		return nil, err
	}
	decoder, ok := v.encoders[ValueEncoderType(decoderType)]
	if !ok {
		return nil, fmt.Errorf("Decoder Not Found for data type %d", decoderType)
	}
	return decoder.Decode(reader)
}

func (v *ValueEncoder) SetEncoder(dataType ValueEncoderType, encoder IValueEncoderUnit) {
	v.encoders[dataType] = encoder
}

func NewValueEncoder(encoders map[ValueEncoderType]IValueEncoderUnit) *ValueEncoder {
	return &ValueEncoder{
		encoders: encoders,
	}
}
