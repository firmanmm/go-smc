# Go Simple Message Codec (SMC)

### Preface
This is a simple message codec rewritten based on version that is used on my `sync-mq` library. I decided to share it with you guys since i think it's ready to be used by the public. Feel free to improve this repository.

# About
Well this is a message codec. It's mainly used to pack and unpack together a bunch of data structure into an array of byte to be transported via network. You can think of it like json marshal and unmarshal but friendlier to machine instead of human. This is a work in progress version that may change in the future. However, if you want to use it right now, i suggest you to fork this project so there is no breaking change in the future. (PS: I do it often as i see fit, so I highly recommend that). 

# Comparison

Well this is the comparison of `smc` against `json`, `jsoniter` and `smc backed with jsoniter` :

```
goos: windows
goarch: amd64
pkg: github.com/firmanmm/gosmc
BenchmarkArrayOfByteJson-8                    	    9230	    125237 ns/op	   28692 B/op	       6 allocs/op
BenchmarkArrayOfByteJsoniter-8                	   23182	     51756 ns/op	   40854 B/op	       5 allocs/op
BenchmarkArrayOfByteSMC-8                     	  240632	      4754 ns/op	   20936 B/op	      15 allocs/op
BenchmarkArrayOfByteSMCWithJsoniter-8         	  261561	      4789 ns/op	   20936 B/op	      15 allocs/op
BenchmarkNestedArrayOfByteJson-8              	     985	   1248471 ns/op	  298811 B/op	      43 allocs/op
BenchmarkNestedArrayOfByteJsoniter-8          	    2038	    556917 ns/op	  431982 B/op	     141 allocs/op
BenchmarkNestedArrayOfByteSMC-8               	    6015	    194994 ns/op	  627835 B/op	    1274 allocs/op
BenchmarkNestedArrayOfByteSMCWithJsoniter-8   	    6330	    192848 ns/op	  627835 B/op	    1274 allocs/op
BenchmarkInterfaceMapJsoniter-8               	  267380	      4469 ns/op	     775 B/op	      25 allocs/op
BenchmarkInterfaceMapSMC-8                    	   97821	     12581 ns/op	    6637 B/op	     199 allocs/op
BenchmarkInterfaceMapSMCWithJsoniter-8        	  197311	      5636 ns/op	    1465 B/op	      38 allocs/op
BenchmarkDeepInterfaceMapJsoniter-8           	    2673	    451853 ns/op	   80890 B/op	    2625 allocs/op
BenchmarkDeepInterfaceMapSMC-8                	     639	   1847946 ns/op	 1970628 B/op	   22249 allocs/op
BenchmarkDeepInterfaceMapSMCWithJsoniter-8    	    2673	    453346 ns/op	   97611 B/op	    2639 allocs/op
BenchmarkStringJson-8                         	  121492	     10031 ns/op	    2385 B/op	       5 allocs/op
BenchmarkStringJsoniter-8                     	  210961	      5702 ns/op	    2210 B/op	       4 allocs/op
BenchmarkStringSMC-8                          	  572259	      1990 ns/op	    4648 B/op	      17 allocs/op
BenchmarkStringSMCWithJsoniter-8              	  600578	      1958 ns/op	    4648 B/op	      17 allocs/op
BenchmarkListStringJson-8                     	    1279	    938069 ns/op	  220872 B/op	     213 allocs/op
BenchmarkListStringJsoniter-8                 	    2268	    569917 ns/op	  228418 B/op	     311 allocs/op
BenchmarkListStringSMC-8                      	     932	   1282013 ns/op	  701372 B/op	   20126 allocs/op
BenchmarkListStringSMCWithJsoniter-8          	    2072	    580991 ns/op	  179621 B/op	    4126 allocs/op
BenchmarkListOfMapJsoniter-8                  	    2797	    420029 ns/op	   82460 B/op	    2511 allocs/op
BenchmarkListOfMapSMC-8                       	     925	   1291712 ns/op	  701372 B/op	   20126 allocs/op
BenchmarkListOfMapSMCWithJsoniter-8           	    2146	    574899 ns/op	  179619 B/op	    4126 allocs/op
PASS
ok  	github.com/firmanmm/gosmc	32.642s
```

As you can see. This Simple Message Codec provides higher throughput in certain usecase compared to `jsoniter` and `json`. However, you can also see that this message codec can also take significantly higher memory compared to both of them. So pick your choice between speed and memory. If you want to get the best of both world you can use `smc with jsoniter` variant which use `jsoniter` to handle `map` and `struct`. I always open if you want to improve it or maybe you want to create your own version and need some assistance.

## Note
I highly recommend that you use `jsoniter` as that is more battle tested than this.
