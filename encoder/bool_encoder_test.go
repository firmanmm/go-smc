package encoder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoolEncoder(t *testing.T) {

	testData := []struct {
		Name     string
		Value    bool
		HasError bool
	}{
		{
			"True",
			true,
			false,
		},
		{
			"False",
			false,
			false,
		},
	}

	encoder := NewBoolEncoder()

	for _, val := range testData {
		t.Run(val.Name, func(t *testing.T) {
			writer := NewBufferWriter()
			err := encoder.Encode(val.Value, writer)
			assert.Nil(t, err)
			content, err := writer.GetContent()
			assert.Nil(t, err)
			reader := NewSliceReader(content)
			decoded, err := encoder.Decode(reader)
			assert.Nil(t, err)
			assert.EqualValues(t, val.Value, decoded)
		})
	}
}
