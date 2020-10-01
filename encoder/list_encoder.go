package encoder

import "reflect"

type ListEncoder struct {
	speedUpInterface bool
	valueEncoder     *ValueEncoder
	intEncoder       *IntEncoder
}

func (l *ListEncoder) Encode(data interface{}, writer IWriter) error {
	if l.speedUpInterface {
		//Start of auto generated switch for List Encoder
		switch data.(type) {

		case []int:
			return l.EncodeInt(data.([]int), writer)
		case []int8:
			return l.EncodeInt8(data.([]int8), writer)
		case []int16:
			return l.EncodeInt16(data.([]int16), writer)
		case []int32:
			return l.EncodeInt32(data.([]int32), writer)
		case []int64:
			return l.EncodeInt64(data.([]int64), writer)
		case []uint:
			return l.EncodeUint(data.([]uint), writer)
		case []uint8:
			return l.EncodeUint8(data.([]uint8), writer)
		case []uint16:
			return l.EncodeUint16(data.([]uint16), writer)
		case []uint32:
			return l.EncodeUint32(data.([]uint32), writer)
		case []uint64:
			return l.EncodeUint64(data.([]uint64), writer)
		case []bool:
			return l.EncodeBool(data.([]bool), writer)
		case []string:
			return l.EncodeString(data.([]string), writer)
		case []float32:
			return l.EncodeFloat32(data.([]float32), writer)
		case []float64:
			return l.EncodeFloat64(data.([]float64), writer)
		case []interface{}:
			return l.EncodeInterface(data.([]interface{}), writer)
		default:
			break
		}
		//End of auto generated switch for List Encoder
	}
	reflected := reflect.ValueOf(data)
	reflectedLen := reflected.Len()
	if err := l.intEncoder.Encode(reflectedLen, writer); err != nil {
		return err
	}
	for i := 0; i < reflectedLen; i++ {
		if err := l.valueEncoder.Encode(reflected.Index(i).Interface(), writer); err != nil {
			return err
		}
	}
	return nil
}

func (l *ListEncoder) Decode(reader IReader) (interface{}, error) {
	length, err := l.intEncoder.Decode(reader)
	if err != nil {
		return nil, err
	}
	result := make([]interface{}, length.(int))
	for i := 0; i < len(result); i++ {
		res, err := l.valueEncoder.Decode(reader)
		if err != nil {
			return nil, err
		}
		result[i] = res
	}
	return result, nil
}

func NewListEncoder(valueEncoder *ValueEncoder) *ListEncoder {
	return &ListEncoder{
		valueEncoder: valueEncoder,
		intEncoder:   NewIntEncoder(),
	}
}

func NewListInterfaceEncoder(valueEncoder *ValueEncoder) *ListEncoder {
	return &ListEncoder{
		speedUpInterface: true,
		valueEncoder:     valueEncoder,
		intEncoder:       NewIntEncoder(),
	}
}
