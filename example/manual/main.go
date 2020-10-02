package main

import (
	"crypto/sha512"
	"fmt"
	"log"

	"github.com/firmanmm/gosmc"
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
	Payload      map[interface{}]interface{}
}

func main() {
	fingerprint := sha512.Sum512([]byte("A Fingerprint"))

	data := ParentOrganism{
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
		Payload: map[interface{}]interface{}{
			"A Key":    "This is a key",
			"Number":   int32(12345),
			1000:       1234.56789012345678901234567890,
			uint(1000): 8765, //Look the same with previous key, but beware, its a different data type
			"1000":     "THis is a string",
		},
	}

	//Lets just use pure implementation
	encoder := gosmc.NewSimpleMessageCodec()
	encoded, err := encoder.ManualEncode(data, func(data interface{}, write gosmc.WriterFunc) error {
		organismData := data.(ParentOrganism)
		//Lets just encode what we need
		if err := write(organismData.Name); err != nil {
			return err
		}
		//Lets speed up a bit
		if err := write(
			organismData.Age,
			organismData.Species,
			organismData.Fingerprint,
			organismData.Child,
			organismData.PointerChild.Name,
			organismData.Payload); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Fatalln(err.Error())
	}
	decoded, err := encoder.ManualDecode(encoded, func(reader *gosmc.ManualReader) (interface{}, error) {
		result := ParentOrganism{}
		//Lets read the first 6 data
		data, err := reader.ReadN(6)
		if err != nil {
			return nil, err
		}
		//See conversion table to know which data type you need to use
		result.Name = data[0].(string)
		result.Age = data[1].(uint)
		result.Species = data[2].(string)
		result.Fingerprint = data[3].([]byte)
		//Since we are passing a struct, it will be converted to map[string]interface{}, see conversion table
		rawChild := data[4].(map[string]interface{})
		child := Organism{
			Name:    rawChild["Name"].(string),
			Age:     rawChild["Age"].(uint),
			Species: rawChild["Species"].(string),
		}
		result.Child = child
		result.PointerChild = &Organism{
			Name: data[5].(string),
		}
		//It doesn't matter how you read, aslong as there are still valid data in there, read will not return error
		//Think of it like a queue
		payload, err := reader.Read()
		if err != nil {
			return nil, err
		}
		result.Payload = payload.(map[interface{}]interface{})
		return result, nil
	})
	if err != nil {
		log.Fatalln(err.Error())
	}
	//Lets see the original data
	fmt.Println(data) //OUTPUT : {Rendoru 22 Human true false 172.2 [152 77 121 122 155 37 109 58 12 79 188 10 231 17 219 248 232 113 251 119 237 205 193 131 23 73 134 70 240 208 100 84 245 35 9 23 43 145 216 219 102 98 216 246 250 44 157 223 130 103 222 194 7 72 19 0 244 218 107 123 93 199 9 120] {Doru 1 Digital Or Unknown} 0xc00006dc20 map[1000:1234.567890123457 1000:THis is a string A Key:This is a key Number:12345 1000:8765]}
	//Now lets see our decoded data
	fmt.Println(decoded) //OUTPUT : {Rendoru 22 Human false false 0 [152 77 121 122 155 37 109 58 12 79 188 10 231 17 219 248 232 113 251 119 237 205 193 131 23 73 134 70 240 208 100 84 245 35 9 23 43 145 216 219 102 98 216 246 250 44 157 223 130 103 222 194 7 72 19 0 244 218 107 123 93 199 9 120] {Doru 1 Digital Or Unknown} 0xc00006dd40 map[1000:1234.567890123457 1000:THis is a string A Key:This is a key Number:12345 1000:8765]}
}
