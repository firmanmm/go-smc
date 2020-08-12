# Go Simple Message Codec

### Preface
This is a simple message codex based on `rewritten` version that is used on my `sync-mq` library. I decided to share it with you guys since i think it's ready to be used by the public. Feel free to improve this repository.

# About
Well this is a message codec. It's mainly used to pack and unpack together a bunch of data structure into an array of byte to be transported via network. You can think of it like json marshal and unmarshal but friendlier to machine instead of human. This is a work in progress version that may change in the future. However, if you want to use it right now, i suggest you to fork this project so there is no breaking change in the future. (PS: I do it often as i see fit, so I highly recommend that). 

# Comparison

Well this is the comparison of `this` against `json` and `jsoniter` :

```
goos: windows
goarch: amd64
pkg: github.com/firmanmm/gosmc
BenchmarkArrayOfByteJson-8             	    9254	    125682 ns/op	   28661 B/op	       6 allocs/op
BenchmarkArrayOfByteJsoniter-8         	   23474	     51588 ns/op	   40895 B/op	       5 allocs/op
BenchmarkArrayOfByteSMC-8              	  548414	      1982 ns/op	   10328 B/op	       6 allocs/op
BenchmarkNestedArrayOfByteJson-8       	     978	   1251901 ns/op	  302938 B/op	      43 allocs/op
BenchmarkNestedArrayOfByteJsoniter-8   	    2188	    537489 ns/op	  441203 B/op	     142 allocs/op
BenchmarkNestedArrayOfByteSMC-8        	   12292	     96522 ns/op	  334282 B/op	     635 allocs/op
BenchmarkInterfaceMapJsoniter-8        	  273252	      4490 ns/op	     775 B/op	      25 allocs/op
BenchmarkInterfaceMapSMC-8             	  158312	      7701 ns/op	    3334 B/op	     125 allocs/op
BenchmarkDeepInterfaceMapJsoniter-8    	    2671	    454245 ns/op	   80962 B/op	    2625 allocs/op
BenchmarkDeepInterfaceMapSMC-8         	     819	   1494765 ns/op	 2278558 B/op	   13925 allocs/op
BenchmarkStringJson-8                  	  121454	     10048 ns/op	    2385 B/op	       5 allocs/op
BenchmarkStringJsoniter-8              	  210963	      5754 ns/op	    2210 B/op	       4 allocs/op
BenchmarkStringSMC-8                   	 1000000	      1000 ns/op	    3256 B/op	       8 allocs/op
BenchmarkListStringJson-8              	    1279	    938652 ns/op	  219324 B/op	     213 allocs/op
BenchmarkListStringJsoniter-8          	    2247	    533766 ns/op	  229207 B/op	     311 allocs/op
BenchmarkListStringSMC-8               	    1562	    794180 ns/op	  366237 B/op	   12715 allocs/op
BenchmarkListOfMapJsoniter-8           	    2862	    421894 ns/op	   82531 B/op	    2511 allocs/op
BenchmarkListOfMapSMC-8                	    6684	    164004 ns/op	  549402 B/op	    1015 allocs/op
```

As you can see. This Simple Message Codex provides higher throughput compared to `jsoniter` and `json`. However, you can also see that this message codex also take significantly higher memory compared to both of them. So pick your choice between speed and memory. I always open if you want to improve it or maybe you want to create your own version and need some assistance.