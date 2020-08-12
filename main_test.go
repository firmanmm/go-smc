package gosmc

import (
	"os"
	"reflect"
	"testing"
)

type TestKit struct {
	codec IMessageCodec
}

func (t *TestKit) setUp() {

}

func (t *TestKit) setDown() {

}

func NewTestKit() *TestKit {

	messageCodec := NewSimpleMessageCodec()

	testKit := &TestKit{
		codec: messageCodec,
	}
	return testKit
}

var testKit *TestKit

func TestMain(m *testing.M) {
	testKit = NewTestKit()
	testKit.setUp()
	code := m.Run()
	testKit.setDown()
	os.Exit(code)
}

func TestNativeDataType(t *testing.T) {

	testData := []struct {
		Name string
		Data interface{}
	}{

		//Can't use reflect deep equal, because everything is an interface
		// {
		// 	"Name",
		// 	byte(8),
		// },
		// {
		// 	"Rune",
		// 	rune(16),
		// },
		// {
		// 	"Int",
		// 	int(-3264),
		// },
		// {
		// 	"Int32",
		// 	int32(-32),
		// },
		// {
		// 	"Int64",
		// 	int64(-64),
		// },
		// {
		// 	"Uint",
		// 	uint(3264),
		// },
		// {
		// 	"Uint32",
		// 	uint(32),
		// },
		// {
		// 	"Uint64",
		// 	uint64(64),
		// },
		// {
		// 	"Float32",
		// 	float32(32.32),
		// },
		// {
		// 	"Float64",
		// 	float64(64.64),
		// },
		// {
		// 	"ByteArray",
		// 	[]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		// },
		// {
		// 	"String",
		// 	"A String",
		// },
	}

	codec := testKit.codec
	for _, val := range testData {
		t.Run(val.Name, func(t *testing.T) {
			encoded, err := codec.Encode(val.Data)
			if err != nil {
				t.Errorf("Got an error %s", err.Error())
			}
			if reflect.DeepEqual(encoded, val.Data) {
				t.Errorf("Expected data to be transformed but nothing happens, %v", encoded)
			}
			decoded, err := codec.Decode(encoded)
			if err != nil {
				t.Errorf("Got an error %s", err.Error())
			}
			if !reflect.DeepEqual(val.Data, decoded) {
				t.Errorf("Expected %v but got %v", val.Data, decoded)
			}
		})
	}

}
