package encoder

import (
	"testing"
)

func _GetLinkedListEncoder() *LinkedListEncoder {
	byteArrayEncoder := NewNativeLinkedEncoderUnitAdapter(
		NewByteArrayEncoder(),
	)

	floatEncoder := NewNativeLinkedEncoderUnitAdapter(
		NewFloatEncoder(),
	)

	intEncoder := NewNativeLinkedEncoderUnitAdapter(
		NewIntEncoder(),
	)

	stringEncoder := NewNativeLinkedEncoderUnitAdapter(
		NewStringEncoder(),
	)

	uintEncoder := NewNativeLinkedEncoderUnitAdapter(
		NewUintEncoder(),
	)

	linkedValueEncoder := NewLinkedValueEncoder(
		map[ValueEncoderType]IValueEncoderLinkedUnit{
			ByteArrayValueEncoder: byteArrayEncoder,
			FloatValueEncoder:     floatEncoder,
			IntValueEncoder:       intEncoder,
			StringValueEncoder:    stringEncoder,
			UintValueEncoder:      uintEncoder,
		},
	)

	return NewLinkedListEncoder(
		linkedValueEncoder,
		NewUintEncoder(),
	)
}

func _GetListEncoder() *ListEncoder {
	valueEncoder := NewValueEncoder(
		map[ValueEncoderType]IValueEncoderUnit{
			IntValueEncoder:    NewIntEncoder(),
			UintValueEncoder:   NewUintEncoder(),
			FloatValueEncoder:  NewFloatEncoder(),
			StringValueEncoder: NewStringEncoder(),
		},
	)

	oldEncoder := NewListEncoder(
		valueEncoder,
		NewUintEncoder(),
	)
	return oldEncoder
}

func BenchmarkLinkedListEncoder(b *testing.B) {

	testData := []interface{}{
		"abcdefghijklmnopqrstuvwxyz",
		"1234567890",
		"this is not button mashing",
		1, 2, 3, 4, 5, 6,
	}

	linkedEncoder := _GetLinkedListEncoder()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		encoded, err := linkedEncoder.Encode(testData)
		if err != nil {
			b.Error(err)
		}
		_, err = linkedEncoder.Decode(encoded.GetResult())
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkListEncoder(b *testing.B) {

	testData := []interface{}{
		"abcdefghijklmnopqrstuvwxyz",
		"1234567890",
		"this is not button mashing",
		1, 2, 3, 4, 5, 6,
	}

	encoder := _GetListEncoder()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		encoded, err := encoder.Encode(testData)
		if err != nil {
			b.Error(err)
		}
		_, err = encoder.Decode(encoded)
		if err != nil {
			b.Error(err)
		}
	}
}
