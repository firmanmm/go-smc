package encoder

import "reflect"

type MapEncoder struct {
	uintEncoder  *UintEncoder
	valueEncoder *ValueEncoder
}

func (l *MapEncoder) Encode(data interface{}, tracker *BufferTracker) ([]byte, error) {
	reflected := reflect.ValueOf(data)
	pairData := reflected.MapRange()
	byteList := make([][]byte, 0, reflected.Len()*2)
	for pairData.Next() {
		key := pairData.Key()
		value := pairData.Value()
		encodedKey, err := l.valueEncoder.Encode(key.Interface(), tracker)
		if err != nil {
			return nil, err
		}
		byteList = append(byteList, encodedKey)
		encodedVal, err := l.valueEncoder.Encode(value.Interface(), tracker)
		if err != nil {
			return nil, err
		}
		byteList = append(byteList, encodedVal)
	}
	return l.merge(byteList, tracker)
}

func (l *MapEncoder) merge(byteList [][]byte, tracker *BufferTracker) ([]byte, error) {
	childCount := len(byteList)
	lengthBytes := make([][]byte, childCount)
	lengthCounts := make([]byte, childCount)
	payloadSize := childCount
	for idx, val := range byteList {
		length, err := l.uintEncoder.Encode(uint(len(val)), tracker)
		if err != nil {
			return nil, err
		}
		lengthBytes[idx] = length
		lengthSize := len(length)
		lengthCounts[idx] = byte(lengthSize)
		payloadSize += len(val)
		payloadSize += lengthSize
	}
	childLengthCount, err := l.uintEncoder.Encode(uint(childCount), tracker)
	if err != nil {
		return nil, err
	}
	childLengthCountLength := len(childLengthCount)
	payloadSize += childLengthCountLength
	payload := tracker.Get()
	payload.WriteByte(byte(childLengthCountLength))
	payload.Write(childLengthCount)
	for idx, val := range byteList {
		payload.WriteByte(lengthCounts[idx])
		payload.Write(lengthBytes[idx])
		payload.Write(val)
	}
	return payload.Bytes(), nil
}

func (l *MapEncoder) Decode(data []byte) (interface{}, error) {

	rawCount := int(data[0])
	dCount, err := l.uintEncoder.Decode(data[1 : 1+rawCount])
	if err != nil {
		return nil, err
	}
	count := int(dCount.(uint))
	result := make(map[interface{}]interface{})
	data = data[1+rawCount:]

	var key interface{}
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
		if i%2 == 0 {
			key = dPayload
		} else {
			result[key] = dPayload
			key = nil
		}
		data = data[endIdx:]
	}
	return result, nil
}

func NewMapEncoder(valueEncoder *ValueEncoder) *MapEncoder {
	return &MapEncoder{
		valueEncoder: valueEncoder,
	}
}
