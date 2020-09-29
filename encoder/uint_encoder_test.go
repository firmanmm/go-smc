package encoder

import (
	"math"
	"reflect"
	"testing"
)

func TestUintEncoder(t *testing.T) {
	testData := []struct {
		Name     string
		Value    uint
		HasError bool
	}{
		{
			"Zero",
			0,
			false,
		},
		{
			"Max",
			uint(math.MaxUint64),
			false,
		},
		{
			"10000",
			10000,
			false,
		},
		{
			"123456789",
			123456789,
			false,
		},
	}

	encoder := NewUintEncoder()

	for _, val := range testData {
		t.Run(val.Name, func(t *testing.T) {
			tracker := GetBufferTracker()
			encoded, err := encoder.Encode(val.Value, tracker)
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
			PutBufferTracker(tracker)
		})
	}
}
