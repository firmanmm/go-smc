package encoder

import "reflect"

type ListEncoder struct {
	valueEncoder *ValueEncoder
	uintEncoder  *UintEncoder
}

func (l *ListEncoder) Encode(data interface{}) ([]byte, error) {
	reflected := data.(reflect.Value)
	reflectedLen := reflected.Len()
	encodedList := make([][]byte, 0, reflectedLen)
	for i := 0; i < reflectedLen; i++ {
		encoded, err := l.valueEncoder.Encode(reflected.Index(i).Interface())
		if err != nil {
			return nil, err
		}
		encodedList = append(encodedList, encoded)
	}
	return l.merge(encodedList)
}

func (l *ListEncoder) merge(byteList [][]byte) ([]byte, error) {
	childCount := len(byteList)
	lengthBytes := make([][]byte, childCount)
	lengthCounts := make([]byte, childCount)
	payloadSize := childCount
	for idx, val := range byteList {
		length, err := l.uintEncoder.Encode(uint(len(val)))
		if err != nil {
			return nil, err
		}
		lengthBytes[idx] = length
		lengthSize := len(length)
		lengthCounts[idx] = byte(lengthSize)
		payloadSize += len(val)
		payloadSize += lengthSize
	}
	childLengthCount, err := l.uintEncoder.Encode(uint(childCount))
	if err != nil {
		return nil, err
	}
	childLengthCountLength := len(childLengthCount)
	payloadSize += childLengthCountLength
	payload := make([]byte, 1, 1+childCount+payloadSize)
	payload[0] = byte(childLengthCountLength)
	payload = append(payload, childLengthCount...)
	for idx, val := range byteList {
		payload = append(payload, lengthCounts[idx])
		payload = append(payload, lengthBytes[idx]...)
		payload = append(payload, val...)
	}
	return payload, nil
}

func (l *ListEncoder) Decode(data []byte) (interface{}, error) {
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

func NewListEncoder(valueEncoder *ValueEncoder, uintEncoder *UintEncoder) *ListEncoder {
	return &ListEncoder{
		valueEncoder: valueEncoder,
		uintEncoder:  uintEncoder,
	}
}
