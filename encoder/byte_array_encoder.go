package encoder

type ByteArrayEncoder struct {
}

func (b *ByteArrayEncoder) Encode(data interface{}) ([]byte, error) {
	return data.([]byte), nil
}

func (b *ByteArrayEncoder) Decode(data []byte) (interface{}, error) {
	return data, nil
}

func NewByteArrayEncoder() *ByteArrayEncoder {
	return &ByteArrayEncoder{}
}
