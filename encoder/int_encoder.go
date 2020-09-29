package encoder

type IntEncoder struct {
}

const _INT_ENCODER_MAX_ARRAY_LIMIT = 9

func (i *IntEncoder) Encode(data interface{}) ([]byte, error) {
	byteArray := make([]byte, _INT_ENCODER_MAX_ARRAY_LIMIT)
	intData := data.(int)
	var unsignedData uint
	if intData >= 0 {
		unsignedData = uint(intData)
	} else {
		unsignedData = uint(-intData)
		byteArray[_INT_ENCODER_MAX_ARRAY_LIMIT-1] = 1
	}
	spaceUsed := 1
	for i := _INT_ENCODER_MAX_ARRAY_LIMIT - 2; unsignedData > 0; i-- {
		byteArray[i] = byte(unsignedData % 256)
		unsignedData /= 256
		spaceUsed++
	}
	return byteArray[_INT_ENCODER_MAX_ARRAY_LIMIT-spaceUsed:], nil
}

func (i *IntEncoder) Decode(data []byte) (interface{}, error) {
	intData := 0
	multiplier := 1
	for i := len(data) - 2; i >= 0; i-- {
		intData += int(data[i]) * multiplier
		multiplier *= 256
	}
	if data[len(data)-1] == 1 {
		intData = -intData
	}
	return intData, nil
}

func NewIntEncoder() *IntEncoder {
	return &IntEncoder{}
}
