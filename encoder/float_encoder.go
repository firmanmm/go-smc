package encoder

import (
	"bytes"
	"math"
)

type FloatEncoder struct {
	intEncoder *IntEncoder
}

func (f *FloatEncoder) Encode(data interface{}, tracker *BufferTracker) ([]byte, error) {
	buffer := tracker.Get()
	floatData := data.(float64)
	headPart, fractionPart := math.Modf(floatData) // Remove fraction
	intHeadPart := int(headPart)
	intFracPart := int((math.MaxInt64) * fractionPart) // Convert to Int part so it can be encoded by IntEncoder
	headByte, err := f.intEncoder.Encode(intHeadPart, tracker)
	if err != nil {
		return nil, err
	}
	fracByte, err := f.intEncoder.Encode(intFracPart, tracker)
	if err != nil {
		return nil, err
	}
	return f.merge(headByte, fracByte, buffer)
}

func (f *FloatEncoder) merge(headPart, fracPart []byte, results *bytes.Buffer) ([]byte, error) {
	results.WriteByte(byte(len(headPart)))
	results.Write(headPart)
	results.Write(fracPart)
	return results.Bytes(), nil
}

func (f *FloatEncoder) Decode(data []byte) (interface{}, error) {
	headByte, fracByte, err := f.split(data)
	if err != nil {
		return nil, err
	}
	head, err := f.intEncoder.Decode(headByte)
	if err != nil {
		return nil, err
	}
	frac, err := f.intEncoder.Decode(fracByte)
	if err != nil {
		return nil, err
	}

	headPart := float64(head.(int))
	fracPart := float64(frac.(int)) / (math.MaxInt64)
	result := headPart + fracPart
	return result, nil
}

func (f *FloatEncoder) split(data []byte) ([]byte, []byte, error) {
	headCount := int(data[0])
	head := data[1 : headCount+1]
	frac := data[1+headCount:]
	return head, frac, nil
}

func NewFloatEncoder() *FloatEncoder {
	return &FloatEncoder{
		intEncoder: NewIntEncoder(),
	}
}
