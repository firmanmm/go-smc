package encoder

type IntEncoder struct {
}

func (i *IntEncoder) Encode(data interface{}, tracker *BufferTracker) ([]byte, error) {
	byteArray := tracker.Get()
	intData := data.(int)
	var unsignedData uint
	if intData >= 0 {
		unsignedData = uint(intData)
		byteArray.WriteByte(0)
	} else {
		unsignedData = uint(-intData)
		byteArray.WriteByte(1)
	}
	for unsignedData > 0 {
		byteArray.WriteByte(byte(unsignedData % 256))
		unsignedData /= 256
	}
	return byteArray.Bytes(), nil
}

func (i *IntEncoder) Decode(data []byte) (interface{}, error) {
	intData := 0
	multiplier := 1
	for i := 1; i < len(data); i++ {
		intData += int(data[i]) * multiplier
		multiplier *= 256
	}
	if data[0] == 1 {
		intData = -intData
	}
	return intData, nil
}

func NewIntEncoder() *IntEncoder {
	return &IntEncoder{}
}
