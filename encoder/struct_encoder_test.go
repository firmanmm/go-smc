package encoder

import (
	"crypto/sha512"
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

type Organism struct {
	Name    string
	Age     uint
	Species string
}

type ParentOrganism struct {
	Name         string
	Age          uint
	Species      string
	Active       bool
	Passive      bool
	Weight       float64
	Fingerprint  []byte
	Child        Organism
	PointerChild *Organism
}

func _GetSource() ParentOrganism {

	fingerprint := sha512.Sum512([]byte("A Fingerprint"))

	return ParentOrganism{
		Name:        "Rendoru",
		Age:         22,
		Species:     "Human",
		Active:      true,
		Passive:     false,
		Weight:      172.2,
		Fingerprint: fingerprint[:],
		Child: Organism{
			Name:    "Doru",
			Age:     1,
			Species: "Digital Or Unknown",
		},
		PointerChild: &Organism{
			Name:    "Ren",
			Age:     1,
			Species: "Digital",
		},
	}
}

func TestStructEncoder(t *testing.T) {
	source := _GetSource()
	valueEncoder := NewValueEncoder(
		map[ValueEncoderType]IValueEncoderUnit{
			BoolValueEncoder:      NewBoolEncoder(),
			ByteArrayValueEncoder: NewByteArrayEncoder(),
			IntValueEncoder:       NewIntEncoder(),
			UintValueEncoder:      NewUintEncoder(),
			FloatValueEncoder:     NewFloatEncoder(),
			StringValueEncoder:    NewStringEncoder(),
		},
	)
	listEncoder := NewListEncoder(
		valueEncoder,
	)

	encoder := NewStructEncoder(valueEncoder)
	valueEncoder.SetEncoder(ListValueEncoder, listEncoder)
	valueEncoder.SetEncoder(StructValueEncoder, encoder)

	writer := NewBufferWriter()
	err := encoder.Encode(source, writer)
	assert.Nil(t, err)
	content, err := writer.GetContent()
	assert.Nil(t, err)
	reader := NewSliceReader(content)
	decoded, err := encoder.Decode(reader)
	assert.Nil(t, err)

	originalJSON, err := jsoniter.MarshalToString(source)
	assert.Nil(t, err)
	decodedJSON, err := jsoniter.MarshalToString(decoded)
	assert.Nil(t, err)
	assert.JSONEq(t, originalJSON, decodedJSON)

}

func BenchmarkStructEncoderCached(t *testing.B) {
	source := _GetSource()
	valueEncoder := NewValueEncoder(
		map[ValueEncoderType]IValueEncoderUnit{
			BoolValueEncoder:      NewBoolEncoder(),
			ByteArrayValueEncoder: NewByteArrayEncoder(),
			IntValueEncoder:       NewIntEncoder(),
			UintValueEncoder:      NewUintEncoder(),
			FloatValueEncoder:     NewFloatEncoder(),
			StringValueEncoder:    NewStringEncoder(),
		},
	)
	listEncoder := NewListEncoder(
		valueEncoder,
	)

	encoder := NewStructEncoder(valueEncoder)
	valueEncoder.SetEncoder(ListValueEncoder, listEncoder)
	valueEncoder.SetEncoder(StructValueEncoder, encoder)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		writer := NewBufferWriter()
		err := encoder.Encode(source, writer)
		if err != nil {
			t.Error(err)
		}
		content, err := writer.GetContent()
		if err != nil {
			t.Error(err)
		}
		reader := NewSliceReader(content)
		_, err = encoder.Decode(reader)
		if err != nil {
			t.Error(err)
		}

	}

}

func BenchmarkStructEncoderOld(t *testing.B) {
	source := _GetSource()
	valueEncoder := NewValueEncoder(
		map[ValueEncoderType]IValueEncoderUnit{
			BoolValueEncoder:      NewBoolEncoder(),
			ByteArrayValueEncoder: NewByteArrayEncoder(),
			IntValueEncoder:       NewIntEncoder(),
			UintValueEncoder:      NewUintEncoder(),
			FloatValueEncoder:     NewFloatEncoder(),
			StringValueEncoder:    NewStringEncoder(),
		},
	)
	listEncoder := NewListEncoder(
		valueEncoder,
	)

	encoder := NewStructEncoder(valueEncoder)
	valueEncoder.SetEncoder(ListValueEncoder, listEncoder)
	valueEncoder.SetEncoder(StructValueEncoder, encoder)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		writer := NewBufferWriter()
		err := encoder.Encode_Old(source, writer)
		if err != nil {
			t.Error(err)
		}
		content, err := writer.GetContent()
		if err != nil {
			t.Error(err)
		}
		reader := NewSliceReader(content)
		_, err = encoder.Decode(reader)
		if err != nil {
			t.Error(err)
		}
	}
}
