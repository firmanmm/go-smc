package encoder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListEncoder(t *testing.T) {
	testData := []struct {
		Name     string
		Value    []interface{}
		HasError bool
	}{
		{
			"Int",
			[]interface{}{
				1, 2, 3, 4, 5, 6, 7, 8, 9,
			},
			false,
		},
		{
			"Float",
			[]interface{}{
				10001.222, 131.31, -131.3123, 124.14, 52352.3333, 123.123, 412.22, -123123.13, -0.123321,
			},
			false,
		},
	}

	valueEncoder := NewValueEncoder(
		map[ValueEncoderType]IValueEncoderUnit{
			IntValueEncoder:   NewIntEncoder(),
			UintValueEncoder:  NewUintEncoder(),
			FloatValueEncoder: NewFloatEncoder(),
		},
	)

	encoder := NewListEncoder(
		valueEncoder,
	)
	valueEncoder.SetEncoder(ListValueEncoder, encoder)

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
