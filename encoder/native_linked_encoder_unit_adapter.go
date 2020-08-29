package encoder

//Adapter for linked version
type NativeLinkedEncoderUnitAdapter struct {
	value IValueEncoderUnit
}

func (n *NativeLinkedEncoderUnitAdapter) Encode(data interface{}) (*LinkedByte, error) {
	res, err := n.value.Encode(data)
	if err != nil {
		return nil, err
	}
	output := NewLinkedByte()
	output.Write(res)
	return output, nil
}

func (n *NativeLinkedEncoderUnitAdapter) Decode(data []byte) (interface{}, error) {
	return n.value.Decode(data)
}

func NewNativeLinkedEncoderUnitAdapter(value IValueEncoderUnit) *NativeLinkedEncoderUnitAdapter {
	return &NativeLinkedEncoderUnitAdapter{
		value: value,
	}
}
