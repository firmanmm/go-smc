package gosmc

type IMessageCodec interface {
	Encode(interface{}) ([]byte, error)
	Decode([]byte) (interface{}, error)
}

type IValueEncoder interface {
	Encode(ValueEncoderType, interface{}) ([]byte, error)
	Decode([]byte) (interface{}, error)
}

type ValueEncoderType int

var (
	Byte      ValueEncoderType = 0
	Int                        = 1
	Uint                       = 2
	Float                      = 3
	Complex                    = 4
	ByteArray                  = 5
	String                     = 6
)

type SimpleMessageCodec struct {
	valueEncoder IValueEncoder
}

func (s *SimpleMessageCodec) Encode(value interface{}) ([]byte, error) {
	return nil, nil
}

func (s *SimpleMessageCodec) Decode([]byte) (interface{}, error) {
	return nil, nil
}

func NewSimpleMessageCodec() *SimpleMessageCodec {
	return &SimpleMessageCodec{}
}
