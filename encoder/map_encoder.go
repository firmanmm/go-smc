package encoder

import "reflect"

type MapEncoder struct {
	speedUpCommon bool
	intEncoder    *IntEncoder
	valueEncoder  *ValueEncoder
}

func (l *MapEncoder) Encode(data interface{}, writer IWriter) error {
	if l.speedUpCommon {
		switch data.(type) {
		case map[string]interface{}:
			return l.EncodeStringInterface(data.(map[string]interface{}), writer)
		case map[interface{}]interface{}:
			return l.EncodeInterfaceInterface(data.(map[interface{}]interface{}), writer)
		}
	}
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

func (l *MapEncoder) EncodeStringInterface(data map[string]interface{}, writer IWriter) error {
	if err := l.intEncoder.Encode(len(data), writer); err != nil {
		return err
	}
	for key, value := range data {
		if err := l.valueEncoder.Encode(key, writer); err != nil {
			return err
		}
		if err := l.valueEncoder.Encode(value, writer); err != nil {
			return err
		}
	}
	return nil
}

func (l *MapEncoder) EncodeInterfaceInterface(data map[interface{}]interface{}, writer IWriter) error {
	if err := l.intEncoder.Encode(len(data), writer); err != nil {
		return err
	}
	for key, value := range data {
		if err := l.valueEncoder.Encode(key, writer); err != nil {
			return err
		}
		if err := l.valueEncoder.Encode(value, writer); err != nil {
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

func NewMapCommonEncoder(valueEncoder *ValueEncoder) *MapEncoder {
	return &MapEncoder{
		speedUpCommon: true,
		valueEncoder:  valueEncoder,
		intEncoder:    NewIntEncoder(),
	}
}
