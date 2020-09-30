package encoder

import "reflect"

type ListEncoder struct {
	valueEncoder *ValueEncoder
	intEncoder   *IntEncoder
}

func (l *ListEncoder) Encode(data interface{}, writer IWriter) error {
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
