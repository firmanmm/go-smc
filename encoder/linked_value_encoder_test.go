package encoder

import (
	"reflect"
	"testing"
)

func TestLinkedValueEncoder(t *testing.T) {
	testData := []struct {
		Name     string
		Value    interface{}
		HasError bool
	}{
		{
			"int",
			-100,
			false,
		},
		{
			"uint",
			uint(100),
			false,
		},
		{
			"float64",
			float64(12214.1313),
			false,
		},
		{
			"string",
			"This is string",
			false,
		},
		{
			"bytes",
			[]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
			false,
		},
	}

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

	valueEncoder := NewValueEncoder(
		map[ValueEncoderType]IValueEncoderUnit{
			ByteArrayValueEncoder: NewByteArrayEncoder(),
			FloatValueEncoder:     NewFloatEncoder(),
			IntValueEncoder:       NewIntEncoder(),
			StringValueEncoder:    NewStringEncoder(),
			UintValueEncoder:      NewUintEncoder(),
		},
	)

	for _, val := range testData {
		t.Run(val.Name, func(t *testing.T) {
			encoded, err := linkedValueEncoder.Encode(val.Value)
			if err != nil != val.HasError {
				t.Errorf("Expected error value of %v but got %v", val.HasError, err != nil)
			}
			oldEncoded, err := valueEncoder.Encode(val.Value)
			if err != nil != val.HasError {
				t.Errorf("Expected error value of %v but got %v", val.HasError, err != nil)
			}
			if reflect.DeepEqual(encoded, val.Value) {
				t.Errorf("Expected data to be transformed but nothing happens, %v", encoded)
			}
			if !reflect.DeepEqual(encoded.GetResult(), oldEncoded) {
				t.Errorf("New data is not the same with old data, New : %v, Old : %v", encoded.GetResult(), oldEncoded)
			}
		})
	}
}
