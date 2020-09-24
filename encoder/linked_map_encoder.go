package encoder

import "reflect"

type LinkedMapEncoder struct {
	valueEncoder *LinkedValueEncoder
	uintEncoder  *UintEncoder
}

func (l *LinkedMapEncoder) Encode(data interface{}) (*LinkedByte, error) {
	output := NewLinkedByte()
	reflected := reflect.ValueOf(data)
	mapEntry := reflected.MapRange()

	childCount, err := l.uintEncoder.Encode(uint(reflected.Len()))
	if err != nil {
		return nil, err
	}
	output.WriteByte(byte(len(childCount)))
	output.Write(childCount)

	for mapEntry.Next() {
		key := mapEntry.Key().Interface()
		value := mapEntry.Value().Interface()
		err := l.mapVal(key, value, output)
		if err != nil {
			return nil, err
		}
	}
	return output, nil
}

func (l *LinkedMapEncoder) mapVal(key, val interface{}, output *LinkedByte) error {
	if err := l.pack(key, output); err != nil {
		return err
	}
	if err := l.pack(val, output); err != nil {
		return err
	}
	return nil
}

func (l *LinkedMapEncoder) pack(val interface{}, output *LinkedByte) error {
	result, err := l.valueEncoder.Encode(val)
	if err != nil {
		return err
	}
	size, err := l.uintEncoder.Encode(uint(result.GetSize()))
	if err != nil {
		return err
	}
	output.WriteByte(byte(len(size)))
	output.Write(size)
	output.Merge(result)
	return nil
}

func (l *LinkedMapEncoder) Decode(data []byte) (interface{}, error) {

	rawCount := int(data[0])
	dCount, err := l.uintEncoder.Decode(data[1 : 1+rawCount])
	if err != nil {
		return nil, err
	}
	count := int(dCount.(uint))
	data = data[1+rawCount:]
	result := make(map[interface{}]interface{})

	var key interface{}
	var val interface{}
	limit := count * 2
	for i := 0; i < limit; i++ {
		childLength := data[0]
		startIdx := 1 + childLength
		dPayloadSize, err := l.uintEncoder.Decode(data[1:startIdx])
		if err != nil {
			return nil, err
		}
		payloadSize := dPayloadSize.(uint)
		endIdx := 1 + int(childLength) + int(payloadSize)
		dPayload, err := l.valueEncoder.Decode(data[startIdx:endIdx])
		if err != nil {
			return nil, err
		}
		if i%2 == 0 {
			key = dPayload
		} else {
			val = dPayload
			result[key] = val
		}
		data = data[endIdx:]
	}

	return result, nil
}

func (l *LinkedMapEncoder) listDecode(data []byte) (interface{}, error) {
	rawCount := int(data[0])
	dCount, err := l.uintEncoder.Decode(data[1 : 1+rawCount])
	if err != nil {
		return nil, err
	}
	count := int(dCount.(uint))
	result := make([]interface{}, 0, count)
	data = data[1+rawCount:]
	for i := 0; i < count; i++ {
		childLength := data[0]
		startIdx := 1 + childLength
		dPayloadSize, err := l.uintEncoder.Decode(data[1:startIdx])
		if err != nil {
			return nil, err
		}
		payloadSize := dPayloadSize.(uint)
		endIdx := 1 + int(childLength) + int(payloadSize)
		dPayload, err := l.valueEncoder.Decode(data[startIdx:endIdx])
		if err != nil {
			return nil, err
		}
		result = append(result, dPayload)
		data = data[endIdx:]
	}
	return result, nil
}

func NewLinkedMapEncoder(
	valueEncoder *LinkedValueEncoder,
	uintEncoder *UintEncoder) *LinkedMapEncoder {
	return &LinkedMapEncoder{
		valueEncoder: valueEncoder,
		uintEncoder:  uintEncoder,
	}
}
