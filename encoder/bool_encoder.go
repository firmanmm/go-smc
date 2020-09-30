package encoder

type BoolEncoder struct {
}

func (b *BoolEncoder) Encode(raw interface{}, writer IWriter) error {
	data := raw.(bool)
	if data {
		writer.WriteByte(1)
	} else {
		writer.WriteByte(0)
	}
	return nil
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
