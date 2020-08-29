package encoder

import (
	"reflect"
	"testing"
)

func TestNativeLinkedUnitAdapter(t *testing.T) {
	testData := []struct {
		Name     string
		Value    int
		HasError bool
	}{
		{
			"int",
			-100,
			false,
		},
		{
			"uint",
			100,
			false,
		},
	}

	intEncoder := NewIntEncoder()
	encoder := NewNativeLinkedEncoderUnitAdapter(
		intEncoder,
	)
	for _, val := range testData {
		t.Run(val.Name, func(t *testing.T) {
			encoded, err := encoder.Encode(val.Value)
			if err != nil != val.HasError {
				t.Errorf("Expected error value of %v but got %v", val.HasError, err != nil)
			}
			oldEncoded, err := intEncoder.Encode(val.Value)
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
