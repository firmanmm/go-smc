package encoder

type UintEncoder struct {
}

func (i *UintEncoder) Encode(data interface{}, tracker *BufferTracker) ([]byte, error) {
	byteArray := tracker.Get()
	unsignedData := data.(uint)
	spaceUsed := 0
	for unsignedData > 0 {
		byteArray.WriteByte(byte(unsignedData % 256))
		unsignedData /= 256
		spaceUsed++
	}
	return byteArray.Bytes(), nil
}

func (i *UintEncoder) Decode(data []byte) (interface{}, error) {
	uintData := uint(0)
	multiplier := uint(1)
	for i := 0; i < len(data); i++ {
		uintData += uint(data[i]) * multiplier
		multiplier *= 256
	}
	return uintData, nil
}

func NewUintEncoder() *UintEncoder {
	return &UintEncoder{}
}
