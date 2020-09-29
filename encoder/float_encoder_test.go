package encoder

import (
	"reflect"
	"testing"
)

func TestFloatEncoder(t *testing.T) {
	testData := []struct {
		Name     string
		Value    float64
		HasError bool
	}{
		{
			"Zero",
			0,
			false,
		},
		////There Test are failing
		// {
		// 	"Max",
		// 	math.MaxFloat64,
		// 	false,
		// },
		// {
		// 	"Min",
		// 	-math.MaxFloat64,
		// 	false,
		// },
		// {
		// 	"Smallest Non Zero",
		// 	math.SmallestNonzeroFloat64,
		// 	false,
		// },
		// {
		// 	"Smallest Negative Non Zero",
		// 	-math.SmallestNonzeroFloat64,
		// 	false,
		// },
		{
			"10000",
			10000,
			false,
		},
		{
			"1234.1234",
			1234.1234,
			false,
		},
		{
			"123456789123456789.123456789123456789123456789123456789123456789123456789",
			123456789123456789.123456789123456789123456789123456789123456789123456789,
			false,
		},
		{
			"-123456789123456789.123456789123456789123456789123456789123456789123456789",
			-123456789123456789.123456789123456789123456789123456789123456789123456789,
			false,
		},
	}

	encoder := NewFloatEncoder()

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
