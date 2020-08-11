package encoder

type IntEncoder struct {
}

func (s *IntEncoder) Encode(data interface{}) ([]byte, error) {
	byteArray := make([]byte, 9)
	intData := data.(int)
	var processedData uint
	if intData >= 0 {
		processedData = uint(intData)
	} else {
		processedData = uint(-intData)
		byteArray[8] = 1
	}
	spaceUsed := 1
	for i := 7; processedData > 0; i-- {
		byteArray[i] = byte(processedData % 256)
		processedData /= 256
		spaceUsed++
	}
	return byteArray[9-spaceUsed:], nil
}

func (s *IntEncoder) Decode(data []byte) (interface{}, error) {
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
