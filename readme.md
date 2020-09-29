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
BenchmarkArrayOfByteJson-8                    	    9230	    126206 ns/op	   28724 B/op	       6 allocs/op
BenchmarkArrayOfByteJsoniter-8                	   23544	     50831 ns/op	   40886 B/op	       5 allocs/op
BenchmarkArrayOfByteSMC-8                     	  462762	      2431 ns/op	   10343 B/op	       5 allocs/op
BenchmarkArrayOfByteSMCWithJsoniter-8         	  481266	      3216 ns/op	   10342 B/op	       5 allocs/op
BenchmarkNestedArrayOfByteJson-8              	     932	   1231682 ns/op	  305641 B/op	      43 allocs/op
BenchmarkNestedArrayOfByteJsoniter-8          	    2227	    530687 ns/op	  436341 B/op	     141 allocs/op
BenchmarkNestedArrayOfByteSMC-8               	   14478	     80805 ns/op	  116616 B/op	     327 allocs/op
BenchmarkNestedArrayOfByteSMCWithJsoniter-8   	   15002	     79643 ns/op	  116196 B/op	     327 allocs/op
BenchmarkInterfaceMapJsoniter-8               	  267164	      4506 ns/op	     775 B/op	      25 allocs/op
BenchmarkInterfaceMapSMC-8                    	  167034	      7320 ns/op	    1210 B/op	      68 allocs/op
BenchmarkInterfaceMapSMCWithJsoniter-8        	  245592	      4926 ns/op	     881 B/op	      28 allocs/op
BenchmarkDeepInterfaceMapJsoniter-8           	    2614	    455934 ns/op	   80858 B/op	    2625 allocs/op
BenchmarkDeepInterfaceMapSMC-8                	    1125	   1014177 ns/op	  140587 B/op	    7668 allocs/op
BenchmarkDeepInterfaceMapSMCWithJsoniter-8    	    2674	    459504 ns/op	   89325 B/op	    2629 allocs/op
BenchmarkStringJson-8                         	  121495	      9924 ns/op	    2385 B/op	       5 allocs/op
BenchmarkStringJsoniter-8                     	  214729	      5685 ns/op	    2210 B/op	       4 allocs/op
BenchmarkStringSMC-8                          	 1000000	      1220 ns/op	    3251 B/op	       7 allocs/op
BenchmarkStringSMCWithJsoniter-8              	  925653	      1211 ns/op	    3251 B/op	       7 allocs/op
BenchmarkListStringJson-8                     	    1280	    938920 ns/op	  224190 B/op	     213 allocs/op
BenchmarkListStringJsoniter-8                 	    2269	    527911 ns/op	  231313 B/op	     311 allocs/op
BenchmarkListStringSMC-8                      	    8548	    134994 ns/op	  320637 B/op	     707 allocs/op
BenchmarkListStringSMCWithJsoniter-8          	    9254	    133531 ns/op	  319871 B/op	     707 allocs/op
BenchmarkListOfMapJsoniter-8                  	    2862	    426892 ns/op	   82506 B/op	    2511 allocs/op
BenchmarkListOfMapSMC-8                       	    1542	    809141 ns/op	  146674 B/op	    6811 allocs/op
BenchmarkListOfMapSMCWithJsoniter-8           	    2504	    496290 ns/op	   92483 B/op	    2907 allocs/op
PASS
ok  	github.com/firmanmm/gosmc	36.129s
```

As you can see. This Simple Message Codec provides higher throughput in certain usecase compared to `jsoniter` and `json`. However, you can also see that this message codec may also take higher memory compared to `jsoniter`. So pick your choice between speed and memory. If you want to get the best of both world you can use `smc with jsoniter` variant which use `jsoniter` to handle `map` and `struct`. I always open if you want to improve it or maybe you want to create your own version and need some assistance.

## Todo
- Improve Map performance
- Make more example

## Note
I highly recommend that you use `jsoniter` as that is more battle tested than this.
