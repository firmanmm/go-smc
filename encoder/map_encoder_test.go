package encoder

import (
	"reflect"
	"testing"
)

func TestMapEncoder(t *testing.T) {
	testData := []struct {
		Name     string
		Value    map[interface{}]interface{}
		HasError bool
	}{
		{
			"Int",
			map[interface{}]interface{}{
				1:  1,
				2:  2,
				-1: -1,
				-2: -2,
				-3: -9,
			},
			false,
		},
		{
			"Combined",
			map[interface{}]interface{}{
				1:              1123.312,
				"Not A Number": 13123,
				-1:             "11111",
				-2:             -2,
				"ww":           "www",
			},
			false,
		},
	}

	valueEncoder := NewValueEncoder(
		map[ValueEncoderType]IValueEncoderUnit{
			IntValueEncoder:    NewIntEncoder(),
			UintValueEncoder:   NewUintEncoder(),
			FloatValueEncoder:  NewFloatEncoder(),
			StringValueEncoder: NewStringEncoder(),
		},
	)
	listEncoder := NewListEncoder(
		valueEncoder, NewUintEncoder(),
	)

	encoder := NewMapEncoder(
		valueEncoder,
	)

	valueEncoder.SetEncoder(ListValueEncoder, listEncoder)
	valueEncoder.SetEncoder(MapValueEncoder, encoder)

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
