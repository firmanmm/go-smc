package gosmc

import (
	"encoding/json"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

func BenchmarkArrayOfByteJson(b *testing.B) {
	data := make([]byte, 10000)
	for i := 0; i < len(data); i++ {
		data[i] = byte(i % 256)
	}

	var dest interface{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res, err := json.Marshal(data)
		if err != nil {
			b.Error(err.Error())
		}
		if err = json.Unmarshal(res, &dest); err != nil {
			b.Error(err.Error())
		}
	}

}

func BenchmarkArrayOfByteJsoniter(b *testing.B) {
	data := make([]byte, 10000)
	for i := 0; i < len(data); i++ {
		data[i] = byte(i % 256)
	}

	var dest interface{}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res, err := jsoniter.Marshal(data)
		if err != nil {
			b.Error(err.Error())
		}
		if err = jsoniter.Unmarshal(res, &dest); err != nil {
			b.Error(err.Error())
		}
	}

}

func BenchmarkArrayOfByteSMC(b *testing.B) {
	encoder := NewSimpleMessageCodec()

	data := make([]byte, 10000)
	for i := 0; i < len(data); i++ {
		data[i] = byte(i % 256)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res, err := encoder.Encode(data)
		if err != nil {
			b.Error(err.Error())
		}
		if _, err = encoder.Decode(res); err != nil {
			b.Error(err.Error())
		}
	}
}

func BenchmarkArrayOfByteSMCWithJsoniter(b *testing.B) {
	encoder := NewSimpleMessageCodecWithJsoniter()

	data := make([]byte, 10000)
	for i := 0; i < len(data); i++ {
		data[i] = byte(i % 256)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res, err := encoder.Encode(data)
		if err != nil {
			b.Error(err.Error())
		}
		if _, err = encoder.Decode(res); err != nil {
			b.Error(err.Error())
		}
	}
}

func BenchmarkNestedArrayOfByteJson(b *testing.B) {
	childData := make([]byte, 10000)
	for i := 0; i < len(childData); i++ {
		childData[i] = byte(i % 256)
	}

	data := make([][]byte, 100)
	for i := 0; i < 10; i++ {
		data[i] = childData
	}

	var dest interface{}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res, err := json.Marshal(data)
		if err != nil {
			b.Error(err.Error())
		}
		if err = json.Unmarshal(res, &dest); err != nil {
			b.Error(err.Error())
		}
	}

}

func BenchmarkNestedArrayOfByteJsoniter(b *testing.B) {
	childData := make([]byte, 10000)
	for i := 0; i < len(childData); i++ {
		childData[i] = byte(i % 256)
	}

	data := make([][]byte, 100)
	for i := 0; i < 10; i++ {
		data[i] = childData
	}

	var dest interface{}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res, err := jsoniter.Marshal(data)
		if err != nil {
			b.Error(err.Error())
		}
		if err = jsoniter.Unmarshal(res, &dest); err != nil {
			b.Error(err.Error())
		}
	}

}

func BenchmarkNestedArrayOfByteSMC(b *testing.B) {
	encoder := NewSimpleMessageCodec()

	childData := make([]byte, 10000)
	for i := 0; i < len(childData); i++ {
		childData[i] = byte(i % 256)
	}

	data := make([][]byte, 100)
	for i := 0; i < 10; i++ {
		data[i] = childData
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res, err := encoder.Encode(data)
		if err != nil {
			b.Error(err.Error())
		}
		if _, err = encoder.Decode(res); err != nil {
			b.Error(err.Error())
		}
	}
}

func BenchmarkNestedArrayOfByteSMCWithJsoniter(b *testing.B) {
	encoder := NewSimpleMessageCodecWithJsoniter()

	childData := make([]byte, 10000)
	for i := 0; i < len(childData); i++ {
		childData[i] = byte(i % 256)
	}

	data := make([][]byte, 100)
	for i := 0; i < 10; i++ {
		data[i] = childData
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res, err := encoder.Encode(data)
		if err != nil {
			b.Error(err.Error())
		}
		if _, err = encoder.Decode(res); err != nil {
			b.Error(err.Error())
		}
	}
}

func BenchmarkInterfaceMapJsoniter(b *testing.B) {

	data := map[interface{}]interface{}{
		1:              1123.312,
		"Not A Number": 13123,
		-1:             "11111",
		-2:             -2,
		"ww":           "www",
	}

	var dest interface{}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res, err := jsoniter.Marshal(data)
		if err != nil {
			b.Error(err.Error())
		}
		if err = jsoniter.Unmarshal(res, &dest); err != nil {
			b.Error(err.Error())
		}
	}

}

func BenchmarkInterfaceMapSMC(b *testing.B) {
	encoder := NewSimpleMessageCodec()

	data := map[interface{}]interface{}{
		1:              1123.312,
		"Not A Number": 13123,
		11:             "11111",
		2:              -2,
		"ww":           "www",
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res, err := encoder.Encode(data)
		if err != nil {
			b.Error(err.Error())
		}

		if _, err = encoder.Decode(res); err != nil {
			b.Error(err.Error())
		}
	}
}

func BenchmarkInterfaceMapSMCWithJsoniter(b *testing.B) {
	encoder := NewSimpleMessageCodecWithJsoniter()

	data := map[interface{}]interface{}{
		1:              1123.312,
		"Not A Number": 13123,
		11:             "11111",
		2:              -2,
		"ww":           "www",
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res, err := encoder.Encode(data)
		if err != nil {
			b.Error(err.Error())
		}

		if _, err = encoder.Decode(res); err != nil {
			b.Error(err.Error())
		}
	}
}

func BenchmarkDeepInterfaceMapJsoniter(b *testing.B) {

	data := map[interface{}]interface{}{
		1:              1123.312,
		"Not A Number": 13123,
		-1:             "11111",
		-2:             -2,
		"ww":           "www",
	}
	iter := data
	for i := 0; i < 100; i++ {
		child := map[interface{}]interface{}{
			1:              1123.312,
			"Not A Number": 13123,
			-1:             "11111",
			-2:             -2,
			"ww":           "www",
		}
		iter["child"] = child
		iter = child
	}

	var dest interface{}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res, err := jsoniter.Marshal(data)
		if err != nil {
			b.Error(err.Error())
		}
		if err = jsoniter.Unmarshal(res, &dest); err != nil {
			b.Error(err.Error())
		}
	}

}

func BenchmarkDeepInterfaceMapSMC(b *testing.B) {
	encoder := NewSimpleMessageCodec()

	data := map[interface{}]interface{}{
		1:              1123.312,
		"Not A Number": 13123,
		-1:             "11111",
		-2:             -2,
		"ww":           "www",
	}
	iter := data
	for i := 0; i < 100; i++ {
		child := map[interface{}]interface{}{
			1:              1123.312,
			"Not A Number": 13123,
			-1:             "11111",
			-2:             -2,
			"ww":           "www",
		}
		iter["child"] = child
		iter = child
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res, err := encoder.Encode(data)
		if err != nil {
			b.Error(err.Error())
		}

		if _, err = encoder.Decode(res); err != nil {
			b.Error(err.Error())
		}
	}
}

func BenchmarkDeepInterfaceMapSMCWithJsoniter(b *testing.B) {
	encoder := NewSimpleMessageCodecWithJsoniter()

	data := map[interface{}]interface{}{
		1:              1123.312,
		"Not A Number": 13123,
		-1:             "11111",
		-2:             -2,
		"ww":           "www",
	}
	iter := data
	for i := 0; i < 100; i++ {
		child := map[interface{}]interface{}{
			1:              1123.312,
			"Not A Number": 13123,
			-1:             "11111",
			-2:             -2,
			"ww":           "www",
		}
		iter["child"] = child
		iter = child
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res, err := encoder.Encode(data)
		if err != nil {
			b.Error(err.Error())
		}

		if _, err = encoder.Decode(res); err != nil {
			b.Error(err.Error())
		}
	}
}

func BenchmarkStringJson(b *testing.B) {
	data := "A"
	for i := 0; i < 10; i++ {
		data += data
	}

	var dest interface{}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res, err := json.Marshal(data)
		if err != nil {
			b.Error(err.Error())
		}
		if err = json.Unmarshal(res, &dest); err != nil {
			b.Error(err.Error())
		}
	}

}

func BenchmarkStringJsoniter(b *testing.B) {
	data := "A"
	for i := 0; i < 10; i++ {
		data += data
	}

	var dest interface{}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res, err := jsoniter.Marshal(data)
		if err != nil {
			b.Error(err.Error())
		}
		if err = jsoniter.Unmarshal(res, &dest); err != nil {
			b.Error(err.Error())
		}
	}

}

func BenchmarkStringSMC(b *testing.B) {
	encoder := NewSimpleMessageCodec()

	data := "A"
	for i := 0; i < 10; i++ {
		data += data
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res, err := encoder.Encode(data)
		if err != nil {
			b.Error(err.Error())
		}
		if _, err = encoder.Decode(res); err != nil {
			b.Error(err.Error())
		}
	}

}

func BenchmarkStringSMCWithJsoniter(b *testing.B) {
	encoder := NewSimpleMessageCodecWithJsoniter()

	data := "A"
	for i := 0; i < 10; i++ {
		data += data
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res, err := encoder.Encode(data)
		if err != nil {
			b.Error(err.Error())
		}
		if _, err = encoder.Decode(res); err != nil {
			b.Error(err.Error())
		}
	}

}

func BenchmarkListStringJson(b *testing.B) {
	childData := "A"
	for i := 0; i < 10; i++ {
		childData += childData
	}

	data := make([]string, 100)
	for i := 0; i < 100; i++ {
		data[i] = childData
	}

	var dest interface{}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res, err := json.Marshal(data)
		if err != nil {
			b.Error(err.Error())
		}
		if err = json.Unmarshal(res, &dest); err != nil {
			b.Error(err.Error())
		}
	}

}

func BenchmarkListStringJsoniter(b *testing.B) {
	childData := "A"
	for i := 0; i < 10; i++ {
		childData += childData
	}

	data := make([]string, 100)
	for i := 0; i < 100; i++ {
		data[i] = childData
	}

	var dest interface{}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res, err := jsoniter.Marshal(data)
		if err != nil {
			b.Error(err.Error())
		}
		if err = jsoniter.Unmarshal(res, &dest); err != nil {
			b.Error(err.Error())
		}
	}

}

func BenchmarkListStringSMC(b *testing.B) {
	encoder := NewSimpleMessageCodec()

	childData := "A"
	for i := 0; i < 10; i++ {
		childData += childData
	}

	data := make([]string, 100)
	for i := 0; i < 100; i++ {
		data[i] = childData
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res, err := encoder.Encode(data)
		if err != nil {
			b.Error(err.Error())
		}
		if _, err = encoder.Decode(res); err != nil {
			b.Error(err.Error())
		}
	}
}

func BenchmarkListStringSMCWithJsoniter(b *testing.B) {
	encoder := NewSimpleMessageCodecWithJsoniter()

	childData := "A"
	for i := 0; i < 10; i++ {
		childData += childData
	}

	data := make([]string, 100)
	for i := 0; i < 100; i++ {
		data[i] = childData
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res, err := encoder.Encode(data)
		if err != nil {
			b.Error(err.Error())
		}
		if _, err = encoder.Decode(res); err != nil {
			b.Error(err.Error())
		}
	}

}

func BenchmarkListOfMapJsoniter(b *testing.B) {

	childData := map[interface{}]interface{}{
		1:              1123.312,
		"Not A Number": 13123,
		-1:             "11111",
		-2:             -2,
		"ww":           "www",
	}

	data := make([]interface{}, 100)
	for i := 0; i < 100; i++ {
		data[i] = childData
	}

	var dest interface{}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res, err := jsoniter.Marshal(data)
		if err != nil {
			b.Error(err.Error())
		}
		if err = jsoniter.Unmarshal(res, &dest); err != nil {
			b.Error(err.Error())
		}
	}

}

func BenchmarkListOfMapSMC(b *testing.B) {
	encoder := NewSimpleMessageCodec()

	childData := map[interface{}]interface{}{
		1:              1123.312,
		"Not A Number": 13123,
		-1:             "11111",
		-2:             -2,
		"ww":           "www",
	}

	data := make([]interface{}, 100)
	for i := 0; i < 100; i++ {
		data[i] = childData
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res, err := encoder.Encode(data)
		if err != nil {
			b.Error(err.Error())
		}
		if _, err = encoder.Decode(res); err != nil {
			b.Error(err.Error())
		}
	}

}

func BenchmarkListOfMapSMCWithJsoniter(b *testing.B) {
	encoder := NewSimpleMessageCodecWithJsoniter()

	childData := map[interface{}]interface{}{
		1:              1123.312,
		"Not A Number": 13123,
		-1:             "11111",
		-2:             -2,
		"ww":           "www",
	}

	data := make([]interface{}, 100)
	for i := 0; i < 100; i++ {
		data[i] = childData
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res, err := encoder.Encode(data)
		if err != nil {
			b.Error(err.Error())
		}
		if _, err = encoder.Decode(res); err != nil {
			b.Error(err.Error())
		}
	}

}
