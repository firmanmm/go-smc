package encoder

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

func TestJsoniterEncoder(t *testing.T) {
	testData := []struct {
		Name     string
		Value    interface{}
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
	}

	encoder := NewJsoniterEncoder()

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
			originalJSON, err := jsoniter.MarshalToString(val.Value)
			assert.Nil(t, err)
			decodedJSON, err := jsoniter.MarshalToString(decoded)
			assert.Nil(t, err)
			assert.JSONEq(t, originalJSON, decodedJSON)
		})
	}
}
