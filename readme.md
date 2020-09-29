# Go Simple Message Codec (SMC)

### Preface
This is a simple message codec `rewritten` based on version that is used on my `sync-mq` library. I decided to share it with you guys since i think it's ready to be used by the public. Feel free to improve this repository.

## About
Well this is a message codec. It's mainly used to pack and unpack together a bunch of data structure into an array of byte to be transported via network. You can think of it like json marshal and unmarshal but friendlier to machine instead of human. This is a work in progress version that may change in the future. However, if you want to use it right now, i suggest you to fork this project so there is no breaking change in the future. (PS: I do it often as i see fit, so I highly recommend that). 

## Usage
Sample code :
```golang
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

}
```
Sample Output :
```
Encoded Length 80
This is original data
map[A Key:This is a key Float:1234.5678 Number:12345]
This is decoded data
map[A Key:This is a key Float:1234.5678 Number:12345]
This is casted data
map[A Key:This is a key Float:1234.5678 Number:12345]
This is example of accessing decoded data
key: A Key content: This is a key
key: Number content: 12345
key: Float content: 1234.567800
```

## Conversion Table
SMC will convert certain data type to another data type in the process. If you decode something, please use the highest available data type that is compatible with that type.
| Type | Converted |
| :------------- | :---------- |
| []byte | []byte |
| string | string |
| int, int8, int16, int32, int64 | int |
| uint, uint8, uint16, uint32, uint64 | uint |
| float32, float64 | float64 |
| []Any (Except []byte) | []interface{} |
| map[Any]Any | map[interface{}]interface{} |

## Benchmark

Well this is the comparison of `smc` against `json`, `jsoniter` and `smc backed with jsoniter` :

```
goos: windows
goarch: amd64
pkg: github.com/firmanmm/gosmc
BenchmarkArrayOfByteJson-8                    	    9230	    125778 ns/op	   28666 B/op	       6 allocs/op
BenchmarkArrayOfByteJsoniter-8                	   23592	     51026 ns/op	   40879 B/op	       5 allocs/op
BenchmarkArrayOfByteSMC-8                     	  572936	      1974 ns/op	   10328 B/op	       6 allocs/op
BenchmarkArrayOfByteSMCWithJsoniter-8         	  633241	      2035 ns/op	   10328 B/op	       6 allocs/op
BenchmarkNestedArrayOfByteJson-8              	     976	   1229291 ns/op	  305274 B/op	      43 allocs/op
BenchmarkNestedArrayOfByteJsoniter-8          	    2226	    534987 ns/op	  440582 B/op	     142 allocs/op
BenchmarkNestedArrayOfByteSMC-8               	   12558	     98081 ns/op	  332457 B/op	     633 allocs/op
BenchmarkNestedArrayOfByteSMCWithJsoniter-8   	   12546	     97698 ns/op	  332457 B/op	     633 allocs/op
BenchmarkInterfaceMapJsoniter-8               	  273262	      4471 ns/op	     775 B/op	      25 allocs/op
BenchmarkInterfaceMapSMC-8                    	  210964	      5848 ns/op	    2190 B/op	     113 allocs/op
BenchmarkInterfaceMapSMCWithJsoniter-8        	  255840	      4650 ns/op	     888 B/op	      29 allocs/op
BenchmarkDeepInterfaceMapJsoniter-8           	    2614	    456710 ns/op	   80758 B/op	    2625 allocs/op
BenchmarkDeepInterfaceMapSMC-8                	    1167	   1028953 ns/op	 1477599 B/op	   12713 allocs/op
BenchmarkDeepInterfaceMapSMCWithJsoniter-8    	    2613	    460689 ns/op	   89135 B/op	    2630 allocs/op
BenchmarkStringJson-8                         	  121538	      9921 ns/op	    2385 B/op	       5 allocs/op
BenchmarkStringJsoniter-8                     	  210966	      5678 ns/op	    2210 B/op	       4 allocs/op
BenchmarkStringSMC-8                          	 1221459	       990 ns/op	    3256 B/op	       8 allocs/op
BenchmarkStringSMCWithJsoniter-8              	 1217826	       975 ns/op	    3256 B/op	       8 allocs/op
BenchmarkListStringJson-8                     	    1279	    938095 ns/op	  223310 B/op	     213 allocs/op
BenchmarkListStringJsoniter-8                 	    2227	    531134 ns/op	  228187 B/op	     311 allocs/op
BenchmarkListStringSMC-8                      	    7063	    164928 ns/op	  547578 B/op	    1013 allocs/op
BenchmarkListStringSMCWithJsoniter-8          	    7077	    165592 ns/op	  547578 B/op	    1013 allocs/op
BenchmarkListOfMapJsoniter-8                  	    2797	    433602 ns/op	   82551 B/op	    2511 allocs/op
BenchmarkListOfMapSMC-8                       	    2004	    600705 ns/op	  250190 B/op	   11513 allocs/op
BenchmarkListOfMapSMCWithJsoniter-8           	    2456	    491762 ns/op	  115041 B/op	    3213 allocs/op
PASS
ok  	github.com/firmanmm/gosmc	36.211s
```

As you can see. This Simple Message Codec provides higher throughput in certain usecase compared to `jsoniter` and `json`. However, you can also see that this message codec may also take higher memory compared to `jsoniter`. So pick your choice between speed and memory. If you want to get the best of both world you can use `smc with jsoniter` variant which use `jsoniter` to handle `map` and `struct`. I always open if you want to improve it or maybe you want to create your own version and need some assistance.

## Todo
- Improve Map performance
- Make more example

## Note
I highly recommend that you use `jsoniter` as that is more battle tested than this.
