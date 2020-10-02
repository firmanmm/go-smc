package main

import (
	"fmt"
	"testing"

	"github.com/firmanmm/gosmc"
)

func BenchmarkMapStringString(b *testing.B) {
	source := make(map[string]string)
	for i := 0; i < 1000; i++ {
		source[fmt.Sprintf("key_%d", i)] = fmt.Sprintf("This is a %d", i)
	}
	encoder := gosmc.NewSimpleMessageCodec()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		encoded, err := encoder.Encode(source)
		if err != nil {
			b.Error(err)
		}
		_, err = encoder.Decode(encoded)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkMapStringInterface(b *testing.B) {
	source := make(map[string]interface{})
	for i := 0; i < 1000; i++ {
		source[fmt.Sprintf("key_%d", i)] = fmt.Sprintf("This is a %d", i)
	}
	encoder := gosmc.NewSimpleMessageCodec()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		encoded, err := encoder.Encode(source)
		if err != nil {
			b.Error(err)
		}
		_, err = encoder.Decode(encoded)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkMapInterfaceInterface(b *testing.B) {
	source := make(map[interface{}]interface{})
	for i := 0; i < 1000; i++ {
		source[fmt.Sprintf("key_%d", i)] = fmt.Sprintf("This is a %d", i)
	}
	encoder := gosmc.NewSimpleMessageCodec()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		encoded, err := encoder.Encode(source)
		if err != nil {
			b.Error(err)
		}
		_, err = encoder.Decode(encoded)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkMapNestedStringString(b *testing.B) {
	source := make(map[string]map[string]string)
	for i := 0; i < 1000; i++ {
		newMap := make(map[string]string)
		source[fmt.Sprintf("key_%d", i)] = newMap
		for j := 0; j < 1000; j++ {
			newMap[fmt.Sprintf("key_%d", i)] = fmt.Sprintf("This is a %d", i)
		}
	}
	encoder := gosmc.NewSimpleMessageCodec()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		encoded, err := encoder.Encode(source)
		if err != nil {
			b.Error(err)
		}
		_, err = encoder.Decode(encoded)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkMapNestedInterfaceInterface(b *testing.B) {
	source := make(map[interface{}]interface{})
	for i := 0; i < 1000; i++ {
		newMap := make(map[string]interface{})
		source[fmt.Sprintf("key_%d", i)] = newMap
		for j := 0; j < 1000; j++ {
			newMap[fmt.Sprintf("key_%d", i)] = fmt.Sprintf("This is a %d", i)
		}
	}
	encoder := gosmc.NewSimpleMessageCodec()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		encoded, err := encoder.Encode(source)
		if err != nil {
			b.Error(err)
		}
		_, err = encoder.Decode(encoded)
		if err != nil {
			b.Error(err)
		}
	}
}
