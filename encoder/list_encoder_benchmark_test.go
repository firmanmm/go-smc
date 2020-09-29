package encoder

import "testing"

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
		tracker := GetBufferTracker()
		encoded, err := encoder.Encode(testData, tracker)
		if err != nil {
			b.Error(err)
		}
		_, err = encoder.Decode(encoded)
		if err != nil {
			b.Error(err)
		}
		PutBufferTracker(tracker)
	}
}
