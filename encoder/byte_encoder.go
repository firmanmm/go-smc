package encoder

type ByteEncoder struct {
}

func (b *ByteEncoder) Encode(data interface{}) ([]byte, error) {
	return data.([]byte), nil
}

func (b *ByteEncoder) Decode(data []byte) (interface{}, error) {
	return data, nil
}

func NewByteEncoder() *ByteEncoder {
	return &ByteEncoder{}
}
