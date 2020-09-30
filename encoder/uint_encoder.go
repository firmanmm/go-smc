package encoder

type UintEncoder struct {
}

const _UINT_ENCODER_MAX_ARRAY_LIMIT = 8

func (i *UintEncoder) Encode(data interface{}, writer IWriter) error {
	byteArray := make([]byte, _UINT_ENCODER_MAX_ARRAY_LIMIT)
	unsignedData := data.(uint)
	spaceUsed := 0
	for unsignedData > 0 {
		byteArray[spaceUsed] = byte(unsignedData % 256)
		unsignedData /= 256
		spaceUsed++
	}
	if err := writer.WriteByte(byte(spaceUsed)); err != nil {
		return err
	}
	return writer.Write(byteArray[:spaceUsed])
}

func (i *UintEncoder) Decode(reader IReader) (interface{}, error) {
	uIntData := uint(0)
	multiplier := uint(1)
	length, err := reader.ReadByte()
	if err != nil {
		return nil, err
	}
	rawByte, err := reader.Read(int(length))
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(rawByte); i++ {
		nextByte := rawByte[i]
		uIntData += uint(nextByte) * multiplier
		multiplier *= 256
	}
	return uIntData, nil
}

func NewUintEncoder() *UintEncoder {
	return &UintEncoder{}
}
