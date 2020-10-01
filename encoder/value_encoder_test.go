package encoder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockValueEncoderData struct {
	Value int
}

type MockValueEncoderUnit struct{}

func (m *MockValueEncoderUnit) Encode(data interface{}, writer IWriter) error {
	return writer.WriteByte(byte(data.(MockValueEncoderData).Value))
}

func (m *MockValueEncoderUnit) Decode(reader IReader) (interface{}, error) {
	value, err := reader.ReadByte()
	if err != nil {
		return nil, err
	}
	return MockValueEncoderData{
		Value: int(value),
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
		StructValueEncoder: &MockValueEncoderUnit{},
	})
	for _, val := range testData {
		t.Run(val.Name, func(t *testing.T) {
			writer := NewBufferWriter()
			err := encoder.Encode(val.Value, writer)
			assert.Nil(t, err)
			content, err := writer.GetContent()
			assert.Nil(t, err)
			reader := NewSliceReader(content)
			decoded, err := encoder.Decode(reader)
			if err != nil != val.HasError {
				t.Errorf("Expected error value of %v but got %v", val.HasError, err != nil)
			}
			assert.EqualValues(t, val.Value, decoded)
		})
	}

}
