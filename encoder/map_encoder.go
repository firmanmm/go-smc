package encoder

import "reflect"

type MapEncoder struct {
	valueEncoder *ValueEncoder
}

func (l *MapEncoder) Encode(data interface{}) ([]byte, error) {
	reflected := data.(reflect.Value)
	keys := reflected.MapKeys()
	interfaceList := make([]interface{}, 0, len(keys)*2)
	for _, val := range keys {
		interfaceList = append(interfaceList,
			val.Interface(), reflected.MapIndex(val).Interface())
	}
	return l.valueEncoder.Encode(interfaceList)
}

func (l *MapEncoder) Decode(data []byte) (interface{}, error) {
	dListData, err := l.valueEncoder.Decode(data)
	if err != nil {
		return nil, err
	}
	listData := dListData.([]interface{})
	dataLength := len(listData) / 2
	result := make(map[interface{}]interface{})
	for i := 0; i < dataLength; i++ {
		key := listData[i*2]
		val := listData[i*2+1]
		result[key] = val
	}
	return result, nil
}

func NewMapEncoder(valueEncoder *ValueEncoder) *MapEncoder {
	return &MapEncoder{
		valueEncoder: valueEncoder,
	}
}
