package encoder

import "reflect"

type MapEncoder struct {
	intEncoder   *IntEncoder
	valueEncoder *ValueEncoder
}

func (l *MapEncoder) Encode(data interface{}, writer IWriter) error {
	reflected := reflect.ValueOf(data)
	pairData := reflected.MapRange()
	length := reflected.Len()
	if err := l.intEncoder.Encode(length, writer); err != nil {
		return err
	}
	for pairData.Next() {
		key := pairData.Key()
		value := pairData.Value()
		if err := l.valueEncoder.Encode(key.Interface(), writer); err != nil {
			return err
		}
		if err := l.valueEncoder.Encode(value.Interface(), writer); err != nil {
			return err
		}
	}
	return nil
}

func (l *MapEncoder) Decode(reader IReader) (interface{}, error) {

	rawLength, err := l.intEncoder.Decode(reader)
	if err != nil {
		return nil, err
	}
	length := rawLength.(int)
	result := make(map[interface{}]interface{})

	for i := 0; i < length; i++ {
		key, err := l.valueEncoder.Decode(reader)
		if err != nil {
			return nil, err
		}
		value, err := l.valueEncoder.Decode(reader)
		if err != nil {
			return nil, err
		}
		result[key] = value
	}
	return result, nil
}

func NewMapEncoder(valueEncoder *ValueEncoder) *MapEncoder {
	return &MapEncoder{
		valueEncoder: valueEncoder,
		intEncoder:   NewIntEncoder(),
	}
}
