package encoder

import (
	"errors"
	"reflect"
)

type IValueEncoderLinkedUnit interface {
	Encode(interface{}) (*LinkedByte, error)
	Decode([]byte) (interface{}, error)
}

type LinkedValueEncoder struct {
	encoders map[ValueEncoderType]IValueEncoderLinkedUnit
}

func (s *LinkedValueEncoder) Encode(data interface{}) (*LinkedByte, error) {
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

func (v *LinkedValueEncoder) encode(dataType ValueEncoderType, data interface{}) (*LinkedByte, error) {
	output := NewLinkedByte()
	output.WriteByte(byte(dataType))
	packet, err := v.encoders[dataType].Encode(data)
	if err != nil {
		return nil, err
	}
	sizePacket, err := v.encoders[UintValueEncoder].Encode(uint(packet.GetSize()))
	if err != nil {
		return nil, err
	}
	output.WriteByte(byte(sizePacket.GetSize()))
	output.Merge(sizePacket)
	output.Merge(packet)
	return output, nil
}

func (v *LinkedValueEncoder) Decode(data []byte) (interface{}, error) {
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

func (v *LinkedValueEncoder) SetEncoder(dataType ValueEncoderType, encoder IValueEncoderLinkedUnit) {
	v.encoders[dataType] = encoder
}

func NewLinkedValueEncoder(encoders map[ValueEncoderType]IValueEncoderLinkedUnit) *LinkedValueEncoder {
	return &LinkedValueEncoder{
		encoders: encoders,
	}
}
