package encoder

import (
	"math"
	"reflect"
	"testing"
)

func TestIntEncoder(t *testing.T) {
	testData := []struct {
		Name     string
		Value    int
		HasError bool
	}{
		{
			"Zero",
			0,
			false,
		},
		{
			"Max",
			int(math.MaxInt64),
			false,
		},
		{
			"Min",
			int(math.MinInt64),
			false,
		},
		{
			"10000",
			10000,
			false,
		},
	}

	encoder := NewIntEncoder()

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
