package encoder

import (
	"reflect"
	"testing"
)

type MockValueEncoderData struct {
	Value int
}

type MockValueEncoderUnit struct{}

func (m *MockValueEncoderUnit) Encode(data interface{}) ([]byte, error) {
	return []byte{byte(data.(MockValueEncoderData).Value)}, nil
}

func (m *MockValueEncoderUnit) Decode(data []byte) (interface{}, error) {
	value := int(data[0])
	return MockValueEncoderData{
		Value: value,
	}, nil
}

func TestValueEncoderBehaviour(t *testing.T) {

	testData := []struct {
		Name     string
		Value    interface{}
		HasError bool
	}{
		{
			"Accepted",
			MockValueEncoderData{
				Value: 1,
			},
			false,
		},
	}

	encoder := NewValueEncoder(map[ValueEncoderType]IValueEncoderUnit{
		GeneralValueEncoder: &MockValueEncoderUnit{},
		UintValueEncoder:    NewUintEncoder(),
	})
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
