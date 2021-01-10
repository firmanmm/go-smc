package encoder

import (
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapEncoder(t *testing.T) {
	testData := []struct {
		Name     string
		Value    map[interface{}]interface{}
		HasError bool
	}{
		{
			"Int",
			map[interface{}]interface{}{
				1:  1,
				2:  2,
				-1: -1,
				-2: -2,
				-3: -9,
			},
			false,
		},
		{
			"Combined",
			map[interface{}]interface{}{
				1:              1123.312,
				"Not A Number": 13123,
				-1:             "11111",
				-2:             -2,
				"ww":           "www",
			},
			false,
		},
		{
			"Nested",
			map[interface{}]interface{}{
				1:              1123.312,
				"Not A Number": 13123,
				-1:             "11111",
				-2:             -2,
				"ww":           "www",
				"nested": map[interface{}]interface{}{
					1:              1123.312,
					"Not A Number": 13123,
					-1:             "11111",
					-2:             -2,
					"ww":           "www",
				},
			},
			false,
		},
	}

	valueEncoder := NewValueEncoder(
		map[ValueEncoderType]IValueEncoderUnit{
			IntValueEncoder:    NewIntEncoder(),
			UintValueEncoder:   NewUintEncoder(),
			FloatValueEncoder:  NewFloatEncoder(),
			StringValueEncoder: NewStringEncoder(),
		},
	)
	listEncoder := NewListEncoder(
		valueEncoder,
	)

	encoder := NewMapEncoder(
		valueEncoder,
	)

	valueEncoder.SetEncoder(ListValueEncoder, listEncoder)
	valueEncoder.SetEncoder(MapValueEncoder, encoder)

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

func TestMapCompabilityB64(t *testing.T) {
	valueEncoder := NewValueEncoder(map[ValueEncoderType]IValueEncoderUnit{
		IntValueEncoder:    NewIntEncoder(),
		StringValueEncoder: NewStringEncoder(),
		FloatValueEncoder:  NewFloatEncoder(),
	})
	data := map[interface{}]interface{}{
		1:              1123.312,
		"Not A Number": 13123,
		-1:             "11111",
		-2:             -2,
		"ww":           "www",
	}
	encoder := NewMapEncoder(valueEncoder)
	valueEncoder.SetEncoder(MapValueEncoder, encoder)
	writer := NewBufferWriter()
	err := encoder.Encode(data, writer)
	assert.Nil(t, err)
	content, err := writer.GetContent()
	assert.Nil(t, err)
	encoded := base64.StdEncoding.EncodeToString(content)
	fmt.Println(encoded)
}
