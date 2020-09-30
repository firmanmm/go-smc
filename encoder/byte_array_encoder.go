package encoder

type ByteArrayEncoder struct {
	intEncoder *IntEncoder
}

func (b *ByteArrayEncoder) Encode(raw interface{}, writer IWriter) error {
	data := raw.([]byte)
	if err := b.intEncoder.Encode(len(data), writer); err != nil {
		return err
	}
	return writer.Write(data)
}

func (b *ByteArrayEncoder) Decode(reader IReader) (interface{}, error) {
	length, err := b.intEncoder.Decode(reader)
	if err != nil {
		return nil, err
	}
	return reader.Read(length.(int))
}

func NewByteArrayEncoder() *ByteArrayEncoder {
	return &ByteArrayEncoder{
		intEncoder: NewIntEncoder(),
	}
}
