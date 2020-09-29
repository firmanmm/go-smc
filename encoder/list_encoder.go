package encoder

import "reflect"

type ListEncoder struct {
	valueEncoder *ValueEncoder
	uintEncoder  *UintEncoder
}

func (l *ListEncoder) Encode(data interface{}, tracker *BufferTracker) ([]byte, error) {
	reflected := reflect.ValueOf(data)
	reflectedLen := reflected.Len()
	encodedList := make([][]byte, 0, reflectedLen)
	for i := 0; i < reflectedLen; i++ {
		encoded, err := l.valueEncoder.Encode(reflected.Index(i).Interface(), tracker)
		if err != nil {
			return nil, err
		}
		encodedList = append(encodedList, encoded)
	}
	return l.merge(encodedList, tracker)
}

func (l *ListEncoder) merge(byteList [][]byte, tracker *BufferTracker) ([]byte, error) {
	childCount := len(byteList)
	lengthBytes := make([][]byte, childCount)
	lengthCounts := make([]byte, childCount)
	for idx, val := range byteList {
		length, err := l.uintEncoder.Encode(uint(len(val)), tracker)
		if err != nil {
			return nil, err
		}
		lengthBytes[idx] = length
		lengthSize := len(length)
		lengthCounts[idx] = byte(lengthSize)
	}
	childLengthCount, err := l.uintEncoder.Encode(uint(childCount), tracker)
	if err != nil {
		return nil, err
	}
	childLengthCountLength := len(childLengthCount)
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
