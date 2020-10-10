package encoder

import (
	"reflect"
	"sync"
)

type _StructCache struct {
	Reflected  reflect.Type
	FieldNames []string
}

type StructEncoder struct {
	intEncoder    *IntEncoder
	stringEncoder *StringEncoder
	valueEncoder  *ValueEncoder
}

var structCacheMap map[reflect.Type]*_StructCache
var structCacheMutex *sync.Mutex

func init() {
	structCacheMap = make(map[reflect.Type]*_StructCache)
	structCacheMutex = &sync.Mutex{}
}

func obtainStructCache(value reflect.Value) *_StructCache {
	reflectedType := value.Type()
	structCacheMutex.Lock()
	defer structCacheMutex.Unlock()
	if res, ok := structCacheMap[reflectedType]; !ok {
		fields := make([]string, value.NumField())
		for i := 0; i < len(fields); i++ {
			fields[i] = reflectedType.Field(i).Name
		}
		res = &_StructCache{
			Reflected:  reflectedType,
			FieldNames: fields,
		}
		structCacheMap[reflectedType] = res
		return res
	} else {
		return res
	}
}

func (s *StructEncoder) Encode(data interface{}, writer IWriter) error {
	reflected := reflect.ValueOf(data)
	structCache := obtainStructCache(reflected)
	fields := structCache.FieldNames
	numField := len(fields)
	if err := s.intEncoder.Encode(numField, writer); err != nil {
		return err
	}
	for i := 0; i < numField; i++ {
		if err := s.stringEncoder.Encode(fields[i], writer); err != nil {
			return err
		}
		value := reflected.Field(i)
		if err := s.valueEncoder.Encode(value.Interface(), writer); err != nil {
			return err
		}
	}
	return nil
}

func (s *StructEncoder) Encode_Old(data interface{}, writer IWriter) error {
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
