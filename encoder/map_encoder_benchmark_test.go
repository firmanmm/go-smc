package encoder

import (
	"testing"
)

func BenchmarkMapEncoder(b *testing.B) {

	data := map[interface{}]interface{}{
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
	for i := 0; i < b.N; i++ {
		writer := NewBufferWriter()
		err := encoder.Encode(data, writer)
		if err != nil {
			b.Error(err)
		}
		content, err := writer.GetContent()
		if err != nil {
			b.Error(err)
		}
		reader := NewSliceReader(content)
		_, err = encoder.Decode(reader)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkCommonMapEncoder(b *testing.B) {

	data := map[interface{}]interface{}{
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

	encoder := NewMapCommonEncoder(
		valueEncoder,
	)

	valueEncoder.SetEncoder(ListValueEncoder, listEncoder)
	valueEncoder.SetEncoder(MapValueEncoder, encoder)
	for i := 0; i < b.N; i++ {
		writer := NewBufferWriter()
		err := encoder.Encode(data, writer)
		if err != nil {
			b.Error(err)
		}
		content, err := writer.GetContent()
		if err != nil {
			b.Error(err)
		}
		reader := NewSliceReader(content)
		_, err = encoder.Decode(reader)
		if err != nil {
			b.Error(err)
		}
	}
}
