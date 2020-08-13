package encoder

import (
	"errors"
	"reflect"
)

type IValueEncoderUnit interface {
	Encode(interface{}) ([]byte, error)
	Decode([]byte) (interface{}, error)
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
	GeneralValueEncoder   ValueEncoderType = 255
)

type ValueEncoder struct {
	encoders map[ValueEncoderType]IValueEncoderUnit
}

func (s *ValueEncoder) Encode(data interface{}) ([]byte, error) {
	encoderUsed := ValueEncoderType(0)
	switch data.(type) {
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
			reflectedLength := reflected.Len()
			newData := make([]interface{}, reflectedLength)
			for i := 0; i < reflectedLength; i++ {
				newData[i] = reflected.Index(i).Interface()
			}
			data = newData
			encoderUsed = ListValueEncoder
		case reflect.Map:
			encoderUsed = MapValueEncoder
		default:
			_, ok := s.encoders[GeneralValueEncoder]
			if !ok {
				return nil, errors.New("Unsupported type, try to register fallback encoder")
			}
			encoderUsed = GeneralValueEncoder
		}

	}
	return s.encode(encoderUsed, data)
}

func (v *ValueEncoder) encode(dataType ValueEncoderType, data interface{}) ([]byte, error) {
	result, err := v.encoders[dataType].Encode(data)
	if err != nil {
		return nil, err
	}
	sizeLength, err := v.encoders[UintValueEncoder].Encode(uint(len(result)))
	if err != nil {
		return nil, err
	}
	dataPack := make([]byte, 2, 2+len(sizeLength)+len(result))
	dataPack[0] = byte(dataType)
	dataPack[1] = byte(len(sizeLength))
	dataPack = append(dataPack, sizeLength...)
	dataPack = append(dataPack, result...)
	return dataPack, nil
}

func (v *ValueEncoder) Decode(data []byte) (interface{}, error) {
	decoderUsed := ValueEncoderType(data[0])
	sizeLength := uint(data[1])
	endSizeIdx := 2 + sizeLength
	dataLength, err := v.encoders[UintValueEncoder].Decode(data[2:endSizeIdx])
	if err != nil {
		return nil, err
	}
	decoded, err := v.encoders[decoderUsed].Decode(data[endSizeIdx : endSizeIdx+dataLength.(uint)])
	return decoded, err
}

func (v *ValueEncoder) SetEncoder(dataType ValueEncoderType, encoder IValueEncoderUnit) {
	v.encoders[dataType] = encoder
}

func NewValueEncoder(encoders map[ValueEncoderType]IValueEncoderUnit) *ValueEncoder {
	return &ValueEncoder{
		encoders: encoders,
	}
}
