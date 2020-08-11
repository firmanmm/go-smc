package gosmc

import (
	"reflect"
	"testing"
)

type MockValueEncoderUnit struct{}

func (m *MockValueEncoderUnit) Encode(data interface{}) ([]byte, error) {
	return []byte{byte(data.(int))}, nil
}

func (m *MockValueEncoderUnit) Decode(data []byte) (interface{}, error) {
	return byte(data[0]), nil
}

func TestValueEncoderBehaviour(t *testing.T) {

	testData := []struct {
		Name     string
		DataType ValueEncoderType
		Value    interface{}
		HasError bool
	}{
		{
			"Accepted",
			9999,
			1,
			false,
		},
		{
			"Fail Case",
			0,
			10000,
			true,
		},
	}

	encoder := NewSimpleValueEncoder()
	encoder.encoders[9999] = &MockValueEncoderUnit{}
	for _, val := range testData {
		t.Run(val.Name, func(t *testing.T) {
			encoded, err := encoder.Encode(val.DataType, val.Value)
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
