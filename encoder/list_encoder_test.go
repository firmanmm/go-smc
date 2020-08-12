package encoder

import (
	"reflect"
	"testing"
)

func TestListEncoder(t *testing.T) {
	testData := []struct {
		Name     string
		Value    []interface{}
		HasError bool
	}{
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

	valueEncoder := NewValueEncoder(
		map[ValueEncoderType]IValueEncoderUnit{
			IntValueEncoder:   NewIntEncoder(),
			UintValueEncoder:  NewUintEncoder(),
			FloatValueEncoder: NewFloatEncoder(),
		},
	)

	encoder := NewListEncoder(
		valueEncoder,
		NewUintEncoder(),
	)
	valueEncoder.SetEncoder(ListValueEncoder, encoder)

	for _, val := range testData {
		t.Run(val.Name, func(t *testing.T) {
			encoded, err := encoder.Encode(val.Value)
			if err != nil != val.HasError {
				t.Errorf("Expected error value of %v but got %v", val.HasError, err != nil)
			}
			if reflect.DeepEqual(encoded, val.Value) {
				t.Errorf("Expected data to be transformed but nothing happens, %v", encoded)
			}
			decoded, err := encoder.Decode(encoded)
			if err != nil != val.HasError {
				t.Errorf("Expected error value of %v but got %v", val.HasError, err != nil)
			}
			if !reflect.DeepEqual(val.Value, decoded) {
				t.Errorf("Expected %v but got %v", val.Value, decoded)
			}
		})
	}
}
