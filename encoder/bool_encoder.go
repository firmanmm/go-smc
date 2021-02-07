package encoder

type BoolEncoder struct {
}

func (b *BoolEncoder) Encode(raw interface{}, writer IWriter) error {
	data := raw.(bool)
	if data {
		return writer.WriteByte(1)
	}
	return writer.WriteByte(0)
}

func (b *BoolEncoder) Decode(reader IReader) (interface{}, error) {
	res, err := reader.ReadByte()
	if err != nil {
		return nil, err
	}
	return res > 0, nil
}

func NewBoolEncoder() *BoolEncoder {
	return &BoolEncoder{}
}
