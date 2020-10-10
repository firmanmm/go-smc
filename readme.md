# Go Simple Message Codec (SMC)

### Preface
This is a simple message codec `rewritten` based on version that is used on my `sync-mq` library. I decided to share it with you guys since i think it's ready to be used by the public. Feel free to improve this repository. I don't really like using solution that requires me to generate code in each change so here it goes.

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

Or if you prefer manual write and read structure then you can find it here [example/manual](example/manual/readme.md)

## Conversion Table
SMC will convert certain data type to another data type in the process. If you decode something, please use the highest available data type that is compatible with that type.
| Type | Implementation | Converted |
| :------------- | :------------- | :---------- |
| bool | Any | bool |
| []byte | Pure, Universal | []byte |
| []byte | Jsoniter | string |
| string | Any | string |
| int, int8, int16, int32, int64 | Any | int |
| uint, uint8, uint16, uint32, uint64 | Pure | uint |
| uint, uint8, uint16, uint32, uint64 | Universal, Jsoniter | int |
| float32, float64 | Any | float64 |
| []Any (Except []byte) | Any | []interface{} |
| map[Any]Any | Pure, Universal | map[interface{}]interface{} |
| map[Any]Any | Jsoniter | map[string]interface{} |
| Struct | Any | map[string]interface{} |
| *Any | Any | Dereferenced and follow conversion table |

Also, since `Jsoniter` convert all map keys to string, you simply can't do what the example shows you. That means `int(1000)`, `uint(1000)`, and "1000" are actually the same, and the behaviour is unknown. This same behaviour also happen when using `universal` encoder but it will only happen for `int(1000)` and `uint(1000)`, this because the `universal` encoder attempt to encode to the most commonly supported data type. An example of unsupported data type in other proggramming langguage is `uint` where it doesn't exist in `java` or `python`. If you want safety, then you can use the pure implementation if you need to deal with that kind of key.

### Conversion Optimization
Please refer to [example/optimization](example/optimization/readme.md) for conversion optimization.

## Benchmark
Well this is the comparison of `smc` against `json`, `jsoniter` and `smc backed with jsoniter` :

```
goos: windows
goarch: amd64
pkg: github.com/firmanmm/gosmc
BenchmarkArrayOfByteJson-8                    	    9230	    114540 ns/op	   28686 B/op	       6 allocs/op
BenchmarkArrayOfByteJsoniter-8                	   23493	     51199 ns/op	   40905 B/op	       5 allocs/op
BenchmarkArrayOfByteSMC-8                     	  480783	      2570 ns/op	   10370 B/op	       6 allocs/op
BenchmarkArrayOfByteSMCWithJsoniter-8         	  481306	      2574 ns/op	   10371 B/op	       6 allocs/op
BenchmarkNestedArrayOfByteJson-8              	    1064	   1129528 ns/op	  302500 B/op	      43 allocs/op
BenchmarkNestedArrayOfByteJsoniter-8          	    2148	    537203 ns/op	  439785 B/op	     142 allocs/op
BenchmarkNestedArrayOfByteSMC-8               	   20925	     56815 ns/op	  121252 B/op	     317 allocs/op
BenchmarkNestedArrayOfByteSMCWithJsoniter-8   	   21295	     57746 ns/op	  120695 B/op	     317 allocs/op
BenchmarkInterfaceMapJsoniter-8               	  261391	      4579 ns/op	     775 B/op	      25 allocs/op
BenchmarkInterfaceMapSMC-8                    	  388094	      3035 ns/op	     752 B/op	      37 allocs/op
BenchmarkInterfaceMapSMCWithJsoniter-8        	  245385	      5007 ns/op	     913 B/op	      29 allocs/op
BenchmarkDeepInterfaceMapJsoniter-8           	    2614	    458236 ns/op	   80942 B/op	    2625 allocs/op
BenchmarkDeepInterfaceMapSMC-8                	    3763	    315932 ns/op	   77043 B/op	    3937 allocs/op
BenchmarkDeepInterfaceMapSMCWithJsoniter-8    	    2617	    457683 ns/op	   89217 B/op	    2630 allocs/op
BenchmarkStringJson-8                         	  129331	      9200 ns/op	    2385 B/op	       5 allocs/op
BenchmarkStringJsoniter-8                     	  207441	      5712 ns/op	    2211 B/op	       4 allocs/op
BenchmarkStringSMC-8                          	 1000000	      1058 ns/op	    2257 B/op	       7 allocs/op
BenchmarkStringSMCWithJsoniter-8              	 1000000	      1082 ns/op	    2257 B/op	       7 allocs/op
BenchmarkListStringJson-8                     	    1398	    862498 ns/op	  223813 B/op	     213 allocs/op
BenchmarkListStringJsoniter-8                 	    2227	    533822 ns/op	  232963 B/op	     311 allocs/op
BenchmarkListStringSMC-8                      	   12651	     94054 ns/op	  229434 B/op	     507 allocs/op
BenchmarkListStringSMCWithJsoniter-8          	   13078	     92888 ns/op	  229375 B/op	     507 allocs/op
BenchmarkListOfMapJsoniter-8                  	    2862	    420968 ns/op	   82557 B/op	    2511 allocs/op
BenchmarkListOfMapSMC-8                       	    4455	    278275 ns/op	   72887 B/op	    3507 allocs/op
BenchmarkListOfMapSMCWithJsoniter-8           	    2505	    473782 ns/op	   91144 B/op	    2807 allocs/op
BenchmarkStructJson-8                         	  138244	      8823 ns/op	    2754 B/op	      48 allocs/op
BenchmarkStructJsoniter-8                     	  152300	      7950 ns/op	    2787 B/op	      58 allocs/op
BenchmarkStructSMC-8                          	  125336	      9398 ns/op	    3266 B/op	     130 allocs/op
BenchmarkStructSMCWithJsoniter-8              	  141517	      8640 ns/op	    3173 B/op	      63 allocs/op
PASS
ok  	github.com/firmanmm/gosmc	39.747s
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
- Need to try using Unsafe Operation (prefer not to)
- Make code generator since there are manual encode and decode support

## Note
I highly recommend that you use `jsoniter` as that is more battle tested than this. Also `jsoniter` is easier to read than this (Good luck if you want to read the output when using SMC).
