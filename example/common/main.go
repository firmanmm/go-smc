package main

import (
	"fmt"
	"log"

	"github.com/firmanmm/gosmc"
)

func main() {
	data := map[string]interface{}{
		"A Key":  "This is a key",
		"Number": int32(12345),
		"Float":  1234.5678,
	}
	codec := gosmc.NewSimpleMessageCodec()
	encoded, err := codec.Encode(data)
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Printf("Encoded Length %d\n", len(encoded))

	decodedData, err := codec.Decode(encoded)
	if err != nil {
		log.Fatalln(err.Error())
	}
	//Since the underlying data is map it is need to be accessed as map[interface{}]interface{}, see conversion table
	castedDecodedData := decodedData.(map[interface{}]interface{})
	fmt.Println("This is original data")
	fmt.Println(data)
	fmt.Println("This is decoded data")
	fmt.Println(decodedData)
	fmt.Println("This is casted data")
	fmt.Println(castedDecodedData)
	fmt.Println("This is example of accessing decoded data")
	fmt.Printf("key: %s content: %s\n", "A Key", castedDecodedData["A Key"].(string))
	//Since the underlying data is int32 it is need to be accessed as int, see conversion table
	fmt.Printf("key: %s content: %d\n", "Number", castedDecodedData["Number"].(int))
	fmt.Printf("key: %s content: %f\n", "Float", castedDecodedData["Float"].(float64))
}
