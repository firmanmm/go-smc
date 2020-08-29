package encoder

import (
	"reflect"
	"testing"
)

func TestLinkedListEncoder(t *testing.T) {
	testData := []struct {
		Name     string
		Value    []interface{}
		HasError bool
	}{
		{
			"Single Int",
			[]interface{}{
				1,
			},
			false,
		},
		{
			"Double Int",
			[]interface{}{
				1, 2,
			},
			false,
		},
		{
			"Int",
			[]interface{}{
				1, 2, 3, 4, 5, 6, 7, 8, 9,
			},
			false,
		},
		{
			"Float",
			[]interface{}{
				10001.222, 131.31, -131.3123, 124.14, 52352.3333, 123.123, 412.22, -123123.13, -0.123321,
			},
			false,
		},

		/// Can't use reflect deep equal because returned type is now interface{}
		// {
		// 	"Another List",
		// 	[]interface{}{
		// 		[]int{
		// 			1, 2, 3, 4, 5, 6, 8,
		// 		},
		// 		[]uint{
		// 			14, 2, 3, 4, 5, 6, 111111111118,
		// 		},
		// 	},
		// 	false,
		// },
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
			IntValueEncoder:   NewIntEncoder(),
			UintValueEncoder:  NewUintEncoder(),
			FloatValueEncoder: NewFloatEncoder(),
		},
	)

	oldEncoder := NewListEncoder(
		valueEncoder,
		NewUintEncoder(),
	)

	encoder := NewLinkedListEncoder(
		linkedValueEncoder,
		NewUintEncoder(),
	)
	for _, val := range testData {
		t.Run(val.Name, func(t *testing.T) {
			encoded, err := encoder.Encode(val.Value)
			if err != nil != val.HasError {
				t.Errorf("Expected error value of %v but got %v", val.HasError, err != nil)
			}
			oldEncoded, err := oldEncoder.Encode(val.Value)
			if err != nil != val.HasError {
				t.Errorf("Expected error value of %v but got %v", val.HasError, err != nil)
			}
			if reflect.DeepEqual(encoded, val.Value) {
				t.Errorf("Expected data to be transformed but nothing happens, %v", encoded)
			}
			if !reflect.DeepEqual(encoded.GetResult(), oldEncoded) {
				t.Errorf("New data is not the same with old data, New : %v, Old : %v", encoded.GetResult(), oldEncoded)
			}

			decoded, err := encoder.Decode(encoded.GetResult())
			if err != nil != val.HasError {
				t.Errorf("Expected error value of %v but got %v", val.HasError, err != nil)
			}
			if !reflect.DeepEqual(val.Value, decoded) {
				t.Errorf("Expected %v but got %v", val.Value, decoded)
			}
		})
	}
}
