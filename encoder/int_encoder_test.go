package encoder

import (
	"encoding/base64"
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
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
			math.MaxInt64,
			false,
		},
		{
			"Min",
			math.MinInt64,
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

	encoder := NewIntEncoder()

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

func TestIntCompabilityB64(t *testing.T) {
	encoder := NewIntEncoder()
	writer := NewBufferWriter()
	err := encoder.Encode(123456789, writer)
	assert.Nil(t, err)
	content, err := writer.GetContent()
	assert.Nil(t, err)
	encoded := base64.StdEncoding.EncodeToString(content)
	fmt.Println(encoded)
}
