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
	//Lets just use our standard implementation
	codec := gosmc.NewSimpleMessageCodec()
	encoded, err := codec.Encode(data)
	if err != nil {
		log.Fatalln(err.Error())
	}
	//Do Some random stuff until finish

	//Lets try to decode the encoded data
	decodedData, err := codec.Decode(encoded)
	if err != nil {
		log.Fatalln(err.Error())
	}
	//Since the underlying data is struct it is need to be accessed as map[string]interface{}, see conversion table
	castedDecodedData := decodedData.(map[string]interface{})
	fmt.Println("This is original data")
	fmt.Println(data) //OUTPUT : {Rendoru 22 Human true false 172.2 [152 77 121 122 155 37 109 58 12 79 188 10 231 17 219 248 232 113 251 119 237 205 193 131 23 73 134 70 240 208 100 84 245 35 9 23 43 145 216 219 102 98 216 246 250 44 157 223 130 103 222 194 7 72 19 0 244 218 107 123 93 199 9 120] {Doru 1 Digital Or Unknown} 0xc00006dc20 map[1000:1234.5678 1000:THis is a string A Key:This is a key Number:12345 1000:8765]}
	fmt.Println("This is decoded data")
	fmt.Println(decodedData) //OUTPUT : map[Active:true Age:22 Child:map[Age:1 Name:Doru Species:Digital Or Unknown] Fingerprint:[152 77 121 122 155 37 109 58 12 79 188 10 231 17 219 248 232 113 251 119 237 205 193 131 23 73 134 70 240 208 100 84 245 35 9 23 43 145 216 219 102 98 216 246 250 44 157 223 130 103 222 194 7 72 19 0 244 218 107 123 93 199 9 120] Name:Rendoru Passive:false Payload:map[1000:1234.5678 1000:THis is a string A Key:This is a key Number:12345 1000:8765] PointerChild:map[Age:1 Name:Ren Species:Digital] Species:Human Weight:172.2]
	fmt.Println("This is example of accessing decoded data")
	fmt.Printf("Field: %s content: %s\n", "Name", castedDecodedData["Name"].(string)) //OUTPUT : Field: Name content: Rendoru
	//Since the underlying data is struct it is need to be accessed as map[string]interface{}, see conversion table
	child := castedDecodedData["Child"].(map[string]interface{})
	fmt.Printf("Child Field: %s content: %s\n", "Name", child["Name"].(string)) //OUTPUT : Child Field: Name content: Doru
	//Since the underlying data is map it is need to be accessed as map[interface{}]interface{}, see conversion table

	payload := castedDecodedData["Payload"].(map[interface{}]interface{})
	//Lets access the "1000" int key in the payload, beware of precision lost when encoding float64
	fmt.Printf("Payload Key: %d content: %.40f, original %.40f\n", 1000, payload[1000].(float64), 1234.56789012345678901234567890) //OUTPUT : Payload Key: 1000 content: 1234.5678901234568911604583263397216796875000, original 1234.5678901234568911604583263397216796875000
	//BEWARE : This behaviour is only available when using Pure implementation and not the Jsoniter one,
	//Its becuase SMC will store their original data type in the encoded form
	//Lets now access the "1000" uint key in the payload and see the magic stuff
	fmt.Printf("Payload Key: %d content: %d\n", uint(1000), payload[uint(1000)].(int)) //OUTPUT : Payload Key: 1000 content: 8765
	//Lets now access the "1000" string key in the payload and see the magic stuff
	fmt.Printf("Payload Key: %s content: %s\n", "1000", payload["1000"].(string)) //OUTPUT : Payload Key: 1000 content: THis is a string

	fmt.Println(string(encoded))
}
