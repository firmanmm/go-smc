package encoder

type StringEncoder struct {
	intEncoder *IntEncoder
}

func (s *StringEncoder) Encode(data interface{}, writer IWriter) error {
	converted := data.(string)
	if err := s.intEncoder.Encode(len(converted), writer); err != nil {
		return err
	}
	return writer.WriteString(converted)
}

func (s *StringEncoder) Decode(reader IReader) (interface{}, error) {
	length, err := s.intEncoder.Decode(reader)
	if err != nil {
		return nil, err
	}
	byteData, err := reader.Read(length.(int))
	if err != nil {
		return nil, err
	}
	return string(byteData), nil
}

func NewStringEncoder() *StringEncoder {
	return &StringEncoder{
		intEncoder: NewIntEncoder(),
	}
}
