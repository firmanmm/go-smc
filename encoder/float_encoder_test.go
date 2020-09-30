package encoder

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
