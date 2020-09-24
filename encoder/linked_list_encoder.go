package encoder

import "reflect"

type LinkedListEncoder struct {
	valueEncoder *LinkedValueEncoder
	uintEncoder  *UintEncoder
}

func (l *LinkedListEncoder) Encode(data interface{}) (*LinkedByte, error) {
	output := NewLinkedByte()
	reflected := reflect.ValueOf(data)
	reflectLength := reflected.Len()
	childCount, err := l.uintEncoder.Encode(uint(reflectLength))
	if err != nil {
		return nil, err
	}
	output.WriteByte(byte(len(childCount)))
	output.Write(childCount)

	for i := 0; i < reflectLength; i++ {
		val := reflected.Index(i).Interface()
		result, err := l.valueEncoder.Encode(val)
		if err != nil {
			return nil, err
		}
		size, err := l.uintEncoder.Encode(uint(result.GetSize()))
		if err != nil {
			return nil, err
		}
		output.WriteByte(byte(len(size)))
		output.Write(size)
		output.Merge(result)
	}

	return output, nil
}

func (l *LinkedListEncoder) Decode(data []byte) (interface{}, error) {
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

func NewLinkedListEncoder(
	valueEncoder *LinkedValueEncoder,
	uintEncoder *UintEncoder) *LinkedListEncoder {

	return &LinkedListEncoder{
		valueEncoder: valueEncoder,
		uintEncoder:  uintEncoder,
	}
}
