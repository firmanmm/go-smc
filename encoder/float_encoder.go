package encoder

import (
	"math"
)

type FloatEncoder struct {
	intEncoder *IntEncoder
}

func (f *FloatEncoder) Encode(data interface{}, writer IWriter) error {
	floatData := data.(float64)
	headPart, fractionPart := math.Modf(floatData) // Remove fraction
	intHeadPart := int(headPart)
	intFracPart := int((math.MaxInt64) * fractionPart) // Convert to Int part so it can be encoded by IntEncoder
	if err := f.intEncoder.Encode(intHeadPart, writer); err != nil {
		return err
	}
	if err := f.intEncoder.Encode(intFracPart, writer); err != nil {
		return err
	}
	return nil
}

func (f *FloatEncoder) Decode(reader IReader) (interface{}, error) {
	head, err := f.intEncoder.Decode(reader)
	if err != nil {
		return nil, err
	}
	frac, err := f.intEncoder.Decode(reader)
	if err != nil {
		return nil, err
	}

	headPart := float64(head.(int))
	fracPart := float64(frac.(int)) / (math.MaxInt64)
	result := headPart + fracPart
	return result, nil
}

func NewFloatEncoder() *FloatEncoder {
	return &FloatEncoder{
		intEncoder: NewIntEncoder(),
	}
}
