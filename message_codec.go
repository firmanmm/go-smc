package gosmc

type IMessageCodec interface {
	Encode(interface{}) []byte
	Decode([]byte) interface{}
}

type SimpleMessageCodec struct {
}

func (s *SimpleMessageCodec) Encode(value interface{}) []byte {
	return nil
}

func (s *SimpleMessageCodec) Decode([]byte) interface{} {
	return nil
}

func NewSimpleMessageCodec() *SimpleMessageCodec {
	return &SimpleMessageCodec{}
}
