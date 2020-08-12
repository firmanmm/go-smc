package encoder

type UintEncoder struct {
}

const _UINT_ENCODER_MAX_ARRAY_LIMIT = 8

func (i *UintEncoder) Encode(data interface{}) ([]byte, error) {
	byteArray := make([]byte, _UINT_ENCODER_MAX_ARRAY_LIMIT)
	unsignedData := data.(uint)
	spaceUsed := 0
	for i := _UINT_ENCODER_MAX_ARRAY_LIMIT - 1; unsignedData > 0; i-- {
		byteArray[i] = byte(unsignedData % 256)
		unsignedData /= 256
		spaceUsed++
	}
	return byteArray[_UINT_ENCODER_MAX_ARRAY_LIMIT-spaceUsed:], nil
}

func (i *UintEncoder) Decode(data []byte) (interface{}, error) {
	uIntData := uint(0)
	multiplier := uint(1)
	for i := len(data) - 1; i >= 0; i-- {
		uIntData += uint(data[i]) * multiplier
		multiplier *= 256
	}
	return uIntData, nil
}

func NewUintEncoder() *UintEncoder {
	return &UintEncoder{}
}
