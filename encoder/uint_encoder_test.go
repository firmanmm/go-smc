package encoder

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
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
