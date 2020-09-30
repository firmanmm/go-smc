package encoder

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
			"With Weird\nSymbol!!!!!@#$%^&*(#$%^&*()_#$%^&*()",
			false,
		},
	}

	encoder := NewStringEncoder()

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
