package encoder

type StringEncoder struct {
}

func (s *StringEncoder) Encode(data interface{}, tracker *BufferTracker) ([]byte, error) {
	return []byte(data.(string)), nil
}

func (s *StringEncoder) Decode(data []byte) (interface{}, error) {
	return string(data), nil
}

func NewStringEncoder() *StringEncoder {
	return &StringEncoder{}
}
