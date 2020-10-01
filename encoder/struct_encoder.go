package encoder

import (
	"reflect"
)

type StructEncoder struct {
	intEncoder    *IntEncoder
	stringEncoder *StringEncoder
	valueEncoder  *ValueEncoder
}

func (s *StructEncoder) Encode(data interface{}, writer IWriter) error {
	reflected := reflect.ValueOf(data)
	reflectedType := reflected.Type()
	numField := reflected.NumField()
	if err := s.intEncoder.Encode(numField, writer); err != nil {
		return err
	}
	for i := 0; i < numField; i++ {
		field := reflectedType.Field(i)
		if err := s.stringEncoder.Encode(field.Name, writer); err != nil {
			return err
		}
		value := reflected.Field(i)
		if err := s.valueEncoder.Encode(value.Interface(), writer); err != nil {
			return err
		}
	}
	return nil
}

func (s *StructEncoder) Decode(reader IReader) (interface{}, error) {
	rawLength, err := s.intEncoder.Decode(reader)
	if err != nil {
		return nil, err
	}
	length := rawLength.(int)
	result := make(map[string]interface{})

	for i := 0; i < length; i++ {
		key, err := s.stringEncoder.Decode(reader)
		if err != nil {
			return nil, err
		}
		value, err := s.valueEncoder.Decode(reader)
		if err != nil {
			return nil, err
		}
		result[key.(string)] = value
	}
	return result, nil
}

func NewStructEncoder(valueEncoder *ValueEncoder) *StructEncoder {
	return &StructEncoder{
		intEncoder:    NewIntEncoder(),
		stringEncoder: NewStringEncoder(),
		valueEncoder:  valueEncoder,
	}
}
