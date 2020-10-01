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
	)
	return oldEncoder
}

func _GetListInterfaceEncoder() *ListEncoder {
	valueEncoder := NewValueEncoder(
		map[ValueEncoderType]IValueEncoderUnit{
			IntValueEncoder:    NewIntEncoder(),
			UintValueEncoder:   NewUintEncoder(),
			FloatValueEncoder:  NewFloatEncoder(),
			StringValueEncoder: NewStringEncoder(),
		},
	)

	oldEncoder := NewListInterfaceEncoder(
		valueEncoder,
	)
	return oldEncoder
}

func BenchmarkListEncoder(b *testing.B) {

	testData := []string{
		"abcdefghijklmnopqrstuvwxyz",
		"1234567890",
		"this is not button mashing",
	}

	encoder := _GetListEncoder()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		writer := NewBufferWriter()
		err := encoder.Encode(testData, writer)
		if err != nil {
			b.Error(err)
		}
		content, err := writer.GetContent()
		if err != nil {
			b.Error(err)

		}
		reader := NewSliceReader(content)
		_, err = encoder.Decode(reader)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkListInterfaceEncoder(b *testing.B) {

	testData := []interface{}{
		"abcdefghijklmnopqrstuvwxyz",
		"1234567890",
		"this is not button mashing",
	}

	encoder := _GetListInterfaceEncoder()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		writer := NewBufferWriter()
		err := encoder.Encode(testData, writer)
		if err != nil {
			b.Error(err)
		}
		content, err := writer.GetContent()
		if err != nil {
			b.Error(err)

		}
		reader := NewSliceReader(content)
		_, err = encoder.Decode(reader)
		if err != nil {
			b.Error(err)
		}
	}
}
