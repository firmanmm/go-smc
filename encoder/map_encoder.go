package encoder

import (
	"bytes"
	"reflect"
)

type MapEncoder struct {
	uintEncoder  *UintEncoder
	valueEncoder *ValueEncoder
}

func (m *MapEncoder) Encode(data interface{}, tracker *BufferTracker) ([]byte, error) {
	reflected := reflect.ValueOf(data)
	pairData := reflected.MapRange()
	reflectedLen := reflected.Len()
	buffer := tracker.Get()
	listLength, err := m.uintEncoder.Encode(uint(reflectedLen*2), tracker)
	if err != nil {
		return nil, err
	}
	buffer.WriteByte(byte(len(listLength)))
	buffer.Write(listLength)
	for pairData.Next() {
		key := pairData.Key()
		value := pairData.Value()
		encodedKey, err := m.valueEncoder.Encode(key.Interface(), tracker)
		if err != nil {
			return nil, err
		}
		if err := m.writePacket(encodedKey, buffer, tracker); err != nil {
			return nil, err
		}
		encodedVal, err := m.valueEncoder.Encode(value.Interface(), tracker)
		if err != nil {
			return nil, err
		}
		if err := m.writePacket(encodedVal, buffer, tracker); err != nil {
			return nil, err
		}
	}
	return buffer.Bytes(), nil
}

func (l *MapEncoder) writePacket(data []byte, buffer *bytes.Buffer, tracker *BufferTracker) error {
	length, err := l.uintEncoder.Encode(uint(len(data)), tracker)
	if err != nil {
		return err
	}
	buffer.WriteByte(byte(len(length)))
	buffer.Write(length)
	buffer.Write(data)
	return nil
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
