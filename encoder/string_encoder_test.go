package encoder

import (
	"reflect"
	"testing"
)

func TestStringEncoder(t *testing.T) {
	testData := []struct {
		Name     string
		Value    string
		HasError bool
	}{
		{
			"Normal",
			"Normal",
			false,
		},
		{
			"White Space",
			"White Space",
			false,
		},
		{
			"With Enter",
			"With\nEnter",
			false,
		},
		{
			"With Weird Symbol",
			"With Weird\nSymbol!!!!!",
			false,
		},
	}

	encoder := NewStringEncoder()

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
