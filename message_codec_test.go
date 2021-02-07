package gosmc

import (
	"crypto/sha512"
	"reflect"
	"testing"

	"github.com/firmanmm/gosmc/encoder"
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"github.com/vmihailenco/msgpack"
)

type Organism struct {
	Name    string
	Age     uint
	Species string
}

type ParentOrganism struct {
	Name         string
	Age          uint
	Species      string
	Active       bool
	Passive      bool
	Weight       float64
	Fingerprint  []byte
	Child        Organism
	PointerChild *Organism
}

func _GetStructSource() ParentOrganism {

	fingerprint := sha512.Sum512([]byte("A Fingerprint"))

	return ParentOrganism{
		Name:        "Rendoru",
		Age:         22,
		Species:     "Human",
		Active:      true,
		Passive:     false,
		Weight:      172.2,
		Fingerprint: fingerprint[:],
		Child: Organism{
			Name:    "Doru",
			Age:     1,
			Species: "Digital Or Unknown",
		},
		PointerChild: &Organism{
			Name:    "Ren",
			Age:     1,
			Species: "Digital",
		},
	}
}

type MockValueEncoder struct {
}

func (m *MockValueEncoder) Encode(dataType encoder.ValueEncoderType, data interface{}) ([]byte, error) {
	return []byte{byte(data.(int))}, nil
}

func (m *MockValueEncoder) Decode(data []byte) (interface{}, error) {
	return int(data[0]), nil
}

func TestMessageCodecBehaviour(t *testing.T) {

	testData := []struct {
		Name     string
		Value    interface{}
		HasError bool
	}{
		{
			"Accepted",
			1,
			false,
		},
	}

	codec := NewSimpleMessageCodec()

	for _, val := range testData {
		t.Run(val.Name, func(t *testing.T) {
			encoded, err := codec.Encode(val.Value)
			if err != nil != val.HasError {
				t.Errorf("Expected error value of %v but got %v", val.HasError, err != nil)
			}
			if reflect.DeepEqual(encoded, val.Value) {
				t.Errorf("Expected data to be transformed but nothing happens, %v", encoded)
			}
			decoded, err := codec.Decode(encoded)
			if err != nil != val.HasError {
				t.Errorf("Expected error value of %v but got %v", val.HasError, err != nil)
			}
			if !reflect.DeepEqual(val.Value, decoded) {
				t.Errorf("Expected %v but got %v", val.Value, decoded)
			}
		})
	}
}

func TestMessageCodecIntegration(t *testing.T) {

	testData := []struct {
		Name       string
		Value      interface{}
		HasError   bool
		ExactMatch bool
	}{
		{
			"Nil",
			nil,
			false,
			false,
		},
		{
			"Int",
			-1,
			false,
			true,
		},
		{
			"Bool",
			true,
			false,
			true,
		},
		{
			"Uint",
			uint(1),
			false,
			true,
		},
		{
			"Float32",
			float64(float32(1234.1234)), //Loses precision since encoding float32 to float64
			false,
			true,
		},
		{
			"Float64",
			float64(1234.1234),
			false,
			true,
		},
		{
			"String",
			"This is a string",
			false,
			true,
		},
		{
			"ByteArray",
			[]byte{
				1, 2, 3, 4, 5, 6, 7, 8,
			},
			false,
			true,
		},
		{
			"String Array",
			[]string{
				"This is a string",
				"This is a string_2",
				"This is a string_3",
			},
			false,
			false,
		},
		{
			"Array of String Array",
			[][]string{
				{
					"This is a string",
					"This is a string_2",
					"This is a string_3",
				},
				{
					"This is a string_4",
					"This is a string_5",
					"This is a string_6",
				},
				{
					"This is a string_7",
					"This is a string_8",
					"This is a string_9",
				},
			},
			false,
			false,
		},
		{
			"Map",
			map[interface{}]interface{}{
				1:              1123.312,
				"Not A Number": 13123,
				-1:             "11111",
				-2:             -2,
				"ww":           "www",
				"Nil Key":      nil,
			},
			false,
			true,
		},
		{
			"Nested Map",
			map[interface{}]interface{}{
				1:              1123.312,
				"Not A Number": 13123,
				-1:             "11111",
				-2:             -2,
				"ww":           "www",
				"child": map[interface{}]interface{}{
					1:              1123.312,
					"Not A Number": 13123,
					-1:             "11111",
					-2:             -2,
					"ww":           "www",
				},
			},
			false,
			true,
		},
		{
			"List Map",
			[]interface{}{
				map[interface{}]interface{}{
					1:              1123.312,
					"Not A Number": 13123,
					-1:             "11111",
					-2:             -2,
					"ww":           "www",
				},
				map[interface{}]interface{}{
					1:              1123.312,
					"Not A Number": 13123,
					-1:             "11111",
					-2:             -2,
					"ww":           "www",
				},
			},
			false,
			true,
		},
		{
			"Struct",
			_GetStructSource(),
			false,
			true,
		},
	}

	codec := NewSimpleMessageCodec()

	for _, val := range testData {
		t.Run(val.Name, func(t *testing.T) {
			encoded, err := codec.Encode(val.Value)
			if err != nil != val.HasError {
				t.Errorf("Expected error value of %v but got %v", val.HasError, err != nil)
			}
			if reflect.DeepEqual(encoded, val.Value) {
				t.Errorf("Expected data to be transformed but nothing happens, %v", encoded)
			}
			decoded, err := codec.Decode(encoded)
			if err != nil != val.HasError {
				t.Errorf("Expected error value of %v but got %v", val.HasError, err != nil)
			}
			originalJSON, err := jsoniter.MarshalToString(val.Value)
			assert.Nil(t, err)
			decodedJSON, err := jsoniter.MarshalToString(decoded)
			assert.Nil(t, err)
			assert.JSONEq(t, originalJSON, decodedJSON)
		})
	}
}

func TestSizeComparison(t *testing.T) {

	sourceMap := map[interface{}]interface{}{
		11231313:       1123.312,
		"Not A Number": 13123,
		-114124141:     "11111",
		-2542341242:    -2,
		"ww":           "www",
		"Struct":       _GetStructSource(),
	}

	source := make([]interface{}, 0, 1000)
	for i := 0; i < 1000; i++ {
		source = append(source, sourceMap)
	}

	jsoniterRes, err := jsoniter.Marshal(source)
	assert.Nil(t, err)

	smcEncoder := NewSimpleMessageCodec()
	smcRes, err := smcEncoder.Encode(source)
	assert.Nil(t, err)

	msgpackRes, err := msgpack.Marshal(source)
	assert.Nil(t, err)

	t.Log("Size Comparison : ")
	t.Logf("Jsoniter : %d\n", len(jsoniterRes))
	t.Logf("SMC : %d\n", len(smcRes))
	t.Logf("Msgpack : %d\n", len(msgpackRes))

}

func TestSizeComparisonWithArrayOfByte(t *testing.T) {

	fingerprint := sha512.Sum512([]byte("This is fingerprint"))
	sourceMap := map[interface{}]interface{}{
		11231313:       1123.312,
		"Not A Number": 13123,
		-114124141:     "11111",
		-2542341242:    -2,
		"ww":           "www",
		"Fingerprint":  fingerprint[:],
		"Struct":       _GetStructSource(),
	}

	source := make([]interface{}, 0, 1000)
	for i := 0; i < 1000; i++ {
		source = append(source, sourceMap)
	}

	jsoniterRes, err := jsoniter.Marshal(source)
	assert.Nil(t, err)

	smcEncoder := NewSimpleMessageCodec()
	smcRes, err := smcEncoder.Encode(source)
	assert.Nil(t, err)

	msgpackRes, err := msgpack.Marshal(source)
	assert.Nil(t, err)

	t.Log("Size Comparison : ")
	t.Logf("Jsoniter : %d\n", len(jsoniterRes))
	t.Logf("SMC : %d\n", len(smcRes))
	t.Logf("Msgpack : %d\n", len(msgpackRes))

}
