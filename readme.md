# Go Simple Message Codec (SMC)

### Preface
This is a simple message codec `rewritten` based on version that is used on my `sync-mq` library. I decided to share it with you guys since i think it's ready to be used by the public. Feel free to improve this repository.

## About
Well this is a message codec. It's mainly used to pack and unpack together a bunch of data structure into an array of byte to be transported via network. You can think of it like json marshal and unmarshal but friendlier to machine instead of human. This is a work in progress version that may change in the future. However, if you want to use it right now, i suggest you to fork this project so there is no breaking change in the future. (PS: I do it often as i see fit, so I highly recommend that). 

## Example
Please see example folder and test file to know more about how to use it
```go
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
}

```

## Conversion Table
SMC will convert certain data type to another data type in the process. If you decode something, please use the highest available data type that is compatible with that type.
| Type | Implementation | Converted |
| :------------- | :------------- | :---------- |
| bool | Any | bool |
| []byte | Pure | []byte |
| []byte | Jsoniter | string |
| string | Any | string |
| int, int8, int16, int32, int64 | Any | int |
| uint, uint8, uint16, uint32, uint64 | Any | uint |
| float32, float64 | Any | float64 |
| []Any (Except []byte) | Any | []interface{} |
| map[Any]Any | Pure | map[interface{}]interface{} |
| map[Any]Any | Jsoniter | map[string]interface{} |
| Struct | Any | map[string]interface{} |
| *Any | Any | Dereferenced and follow conversion table |

Also, since `Jsoniter` convert all map keys to string, you simply can't do what the example shows you. That means `int(1000)`, `uint(1000)`, and "1000" are actually the same, and the behaviour is unknown. If you want safety, then you can use the pure implementation if you need to deal with that kind of key.

## Benchmark
Well this is the comparison of `smc` against `json`, `jsoniter` and `smc backed with jsoniter` :

```
goos: windows
goarch: amd64
pkg: github.com/firmanmm/gosmc
BenchmarkArrayOfByteJson-8                    	    9998	    114121 ns/op	   28732 B/op	       6 allocs/op
BenchmarkArrayOfByteJsoniter-8                	   23407	     51513 ns/op	   40893 B/op	       5 allocs/op
BenchmarkArrayOfByteSMC-8                     	  429724	      2639 ns/op	   10370 B/op	       6 allocs/op
BenchmarkArrayOfByteSMCWithJsoniter-8         	  481291	      2543 ns/op	   10370 B/op	       6 allocs/op
BenchmarkNestedArrayOfByteJson-8              	    1054	   1138289 ns/op	  299651 B/op	      43 allocs/op
BenchmarkNestedArrayOfByteJsoniter-8          	    2269	    541976 ns/op	  440338 B/op	     142 allocs/op
BenchmarkNestedArrayOfByteSMC-8               	   20998	     57091 ns/op	  121285 B/op	     317 allocs/op
BenchmarkNestedArrayOfByteSMCWithJsoniter-8   	   20708	     58420 ns/op	  120994 B/op	     317 allocs/op
BenchmarkInterfaceMapJsoniter-8               	  261368	      4587 ns/op	     775 B/op	      25 allocs/op
BenchmarkInterfaceMapSMC-8                    	  308613	      4046 ns/op	    1040 B/op	      49 allocs/op
BenchmarkInterfaceMapSMCWithJsoniter-8        	  245552	      5028 ns/op	     913 B/op	      29 allocs/op
BenchmarkDeepInterfaceMapJsoniter-8           	    2614	    466630 ns/op	   81048 B/op	    2625 allocs/op
BenchmarkDeepInterfaceMapSMC-8                	    2732	    446463 ns/op	  109350 B/op	    5349 allocs/op
BenchmarkDeepInterfaceMapSMCWithJsoniter-8    	    2614	    470051 ns/op	   89394 B/op	    2630 allocs/op
BenchmarkStringJson-8                         	  128005	      9249 ns/op	    2385 B/op	       5 allocs/op
BenchmarkStringJsoniter-8                     	  211084	      5731 ns/op	    2211 B/op	       4 allocs/op
BenchmarkStringSMC-8                          	 1130359	      1013 ns/op	    2257 B/op	       7 allocs/op
BenchmarkStringSMCWithJsoniter-8              	 1212872	      1016 ns/op	    2257 B/op	       7 allocs/op
BenchmarkListStringJson-8                     	    1366	    879078 ns/op	  223814 B/op	     213 allocs/op
BenchmarkListStringJsoniter-8                 	    2226	    534493 ns/op	  230925 B/op	     311 allocs/op
BenchmarkListStringSMC-8                      	   12252	     96708 ns/op	  229908 B/op	     507 allocs/op
BenchmarkListStringSMCWithJsoniter-8          	   12277	     96348 ns/op	  228699 B/op	     507 allocs/op
BenchmarkListOfMapJsoniter-8                  	    2864	    427987 ns/op	   82532 B/op	    2511 allocs/op
BenchmarkListOfMapSMC-8                       	    3248	    385667 ns/op	  101598 B/op	    4707 allocs/op
BenchmarkListOfMapSMCWithJsoniter-8           	    2455	    485462 ns/op	   91171 B/op	    2807 allocs/op
BenchmarkStructJson-8                         	  138255	      8779 ns/op	    2754 B/op	      48 allocs/op
BenchmarkStructJsoniter-8                     	  152240	      7992 ns/op	    2787 B/op	      58 allocs/op
BenchmarkStructSMC-8                          	  128002	      9288 ns/op	    3266 B/op	     130 allocs/op
BenchmarkStructSMCWithJsoniter-8              	  144969	      8497 ns/op	    3173 B/op	      63 allocs/op
PASS
ok  	github.com/firmanmm/gosmc	42.231s
```

Comparing the output generated and we get :
```
=== RUN   TestSizeComparison
--- PASS: TestSizeComparison (0.02s)
    message_codec_test.go:294: Size Comparison :
    message_codec_test.go:295: Jsoniter : 418001
    message_codec_test.go:296: SMC : 365004
PASS
ok      github.com/firmanmm/gosmc       0.539s
```
And if we add slice of byte in to the mix then we get : 
```
=== RUN   TestSizeComparisonWithArrayOfByte
--- PASS: TestSizeComparisonWithArrayOfByte (0.02s)
    message_codec_test.go:325: Size Comparison :
    message_codec_test.go:326: Jsoniter : 523001
    message_codec_test.go:327: SMC : 446004
PASS
ok      github.com/firmanmm/gosmc       0.524s
```
Well, it clearly gives lower output size compared to `jsoniter` with or without having `slice of byte` in the map entry.

As you can see. This Simple Message Codec provides higher throughput in certain usecase compared to `jsoniter` and `json`. However, you can also see that this message codec may also take higher memory compared to `jsoniter`. So pick your choice between speed and memory. If you want to get the best of both world you can use `smc with jsoniter` variant which use `jsoniter` to handle `map` and `struct`. I always open if you want to improve it or maybe you want to create your own version and need some assistance.

## Behavioural Note
When using `jsoniter`, the decoded `map` version will be in the `map[string]interface{}` format where all keys will be converted into `string`. It also apply to codec instantiated with `NewSimpleMessageCodecWithJsoniter`. 

When using `NewSimpleMessageCodec` or `pure` implementation, the decoded `map` will become `map[interface{}]interface{}` and will maintain their original data type. So `int` key will remain `int` and not converted to `string` like `jsoniter`.

## Todo
- Improve Struct performance (Need to be better than jsoniter)
- Make more example
- Need to try using Unsafe Operation

## Note
I highly recommend that you use `jsoniter` as that is more battle tested than this. Also `jsoniter` is easier to read than this (Good luck if you want to read the output when using SMC).
