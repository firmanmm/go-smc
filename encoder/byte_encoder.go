package encoder

type ByteEncoder struct {
}

func (s *ByteEncoder) Encode(data interface{}) ([]byte, error) {
	return data.([]byte), nil
}

func (s *ByteEncoder) Decode(data []byte) (interface{}, error) {
	return data, nil
}

func NewByteEncoder() *ByteEncoder {
	return &ByteEncoder{}
}
