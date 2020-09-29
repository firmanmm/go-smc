package encoder

import (
	"bytes"
	"reflect"
)

type ListEncoder struct {
	valueEncoder *ValueEncoder
	uintEncoder  *UintEncoder
}

func (l *ListEncoder) Encode(data interface{}, tracker *BufferTracker) ([]byte, error) {
	reflected := reflect.ValueOf(data)
	reflectedLen := reflected.Len()
	buffer := tracker.Get()
	listLength, err := l.uintEncoder.Encode(uint(reflectedLen), tracker)
	if err != nil {
		return nil, err
	}
	buffer.WriteByte(byte(len(listLength)))
	buffer.Write(listLength)
	for i := 0; i < reflectedLen; i++ {
		encoded, err := l.valueEncoder.Encode(reflected.Index(i).Interface(), tracker)
		if err != nil {
			return nil, err
		}
		if err := l.writePacket(encoded, buffer, tracker); err != nil {
			return nil, err
		}
	}
	return buffer.Bytes(), nil
}

func (l *ListEncoder) writePacket(data []byte, buffer *bytes.Buffer, tracker *BufferTracker) error {
	length, err := l.uintEncoder.Encode(uint(len(data)), tracker)
	if err != nil {
		return err
	}
	buffer.WriteByte(byte(len(length)))
	buffer.Write(length)
	buffer.Write(data)
	return nil
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
