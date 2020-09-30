package encoder

type IntEncoder struct {
}

const _INT_ENCODER_MAX_ARRAY_LIMIT = 8

func (i *IntEncoder) Encode(raw interface{}, writer IWriter) error {
	byteArray := make([]byte, _INT_ENCODER_MAX_ARRAY_LIMIT)

	data := raw.(int)
	spaceUsed := 0
	isPositive := data > 0
	if isPositive {
		data = -data
	}
	for data < 0 {
		byteArray[spaceUsed] = byte(-data % 256)
		data /= 256
		spaceUsed++
	}
	length := spaceUsed
	if isPositive {
		length |= 128
	}
	if err := writer.WriteByte(byte(length)); err != nil {
		return err
	}
	return writer.Write(byteArray[:spaceUsed])
}

func (i *IntEncoder) Decode(reader IReader) (interface{}, error) {
	data := 0
	multiplier := 1
	length, err := reader.ReadByte()
	if err != nil {
		return nil, err
	}
	isPositive := length&128 > 0
	if isPositive {
		length = length ^ 128
	}

	rawByte, err := reader.Read(int(length))
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(rawByte); i++ {
		nextByte := rawByte[i]
		data += int(nextByte) * multiplier
		multiplier *= 256
	}
	if !isPositive {
		data = -data
	}
	return data, nil
}

func NewIntEncoder() *IntEncoder {
	return &IntEncoder{}
}
