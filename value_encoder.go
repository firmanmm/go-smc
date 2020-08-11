package gosmc

type IValueEncoderUnit interface {
	Encode(interface{}) ([]byte, error)
	Decode([]byte) (interface{}, error)
}

type SimpleValueEncoder struct {
	encoders map[ValueEncoderType]IValueEncoderUnit
}

func (s *SimpleValueEncoder) Encode(dataType ValueEncoderType, data interface{}) ([]byte, error) {
	return nil, nil
}

func (s *SimpleValueEncoder) Decode([]byte) (interface{}, error) {
	return nil, nil
}

func NewSimpleValueEncoder() *SimpleValueEncoder {
	return &SimpleValueEncoder{
		encoders: map[ValueEncoderType]IValueEncoderUnit{},
	}
}
