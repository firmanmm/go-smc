# Go Simple Message Codec (SMC)

### Preface
This is a simple message codec `rewritten` based on version that is used on my `sync-mq` library. I decided to share it with you guys since i think it's ready to be used by the public. Feel free to improve this repository.

# About
Well this is a message codec. It's mainly used to pack and unpack together a bunch of data structure into an array of byte to be transported via network. You can think of it like json marshal and unmarshal but friendlier to machine instead of human. This is a work in progress version that may change in the future. However, if you want to use it right now, i suggest you to fork this project so there is no breaking change in the future. (PS: I do it often as i see fit, so I highly recommend that). 

# Comparison

Well this is the comparison of `smc` against `json`, `jsoniter` and `smc backed with jsoniter` :

```
goos: windows
goarch: amd64
pkg: github.com/firmanmm/gosmc
BenchmarkArrayOfByteJson-8                    	    8594	    125338 ns/op	   28681 B/op	       6 allocs/op
BenchmarkArrayOfByteJsoniter-8                	   23823	     50532 ns/op	   40874 B/op	       5 allocs/op
BenchmarkArrayOfByteSMC-8                     	  546876	      1900 ns/op	   10328 B/op	       6 allocs/op
BenchmarkArrayOfByteSMCWithJsoniter-8         	  668415	      1925 ns/op	   10328 B/op	       6 allocs/op
BenchmarkNestedArrayOfByteJson-8              	     985	   1225185 ns/op	  305719 B/op	      43 allocs/op
BenchmarkNestedArrayOfByteJsoniter-8          	    2268	    531647 ns/op	  442123 B/op	     142 allocs/op
BenchmarkNestedArrayOfByteSMC-8               	   12813	     93796 ns/op	  332457 B/op	     633 allocs/op
BenchmarkNestedArrayOfByteSMCWithJsoniter-8   	   12882	     93913 ns/op	  332457 B/op	     633 allocs/op
BenchmarkInterfaceMapJsoniter-8               	  273244	      4424 ns/op	     775 B/op	      25 allocs/op
BenchmarkInterfaceMapSMC-8                    	  179488	      6829 ns/op	    2805 B/op	     121 allocs/op
BenchmarkInterfaceMapSMCWithJsoniter-8        	  267171	      4629 ns/op	     888 B/op	      29 allocs/op
BenchmarkDeepInterfaceMapJsoniter-8           	    2671	    451805 ns/op	   80918 B/op	    2625 allocs/op
BenchmarkDeepInterfaceMapSMC-8                	     884	   1367423 ns/op	 2222074 B/op	   13521 allocs/op
BenchmarkDeepInterfaceMapSMCWithJsoniter-8    	    2671	    452549 ns/op	   89100 B/op	    2630 allocs/op
BenchmarkStringJson-8                         	  120280	      9892 ns/op	    2385 B/op	       5 allocs/op
BenchmarkStringJsoniter-8                     	  211132	      5673 ns/op	    2210 B/op	       4 allocs/op
BenchmarkStringSMC-8                          	 1208043	       965 ns/op	    3256 B/op	       8 allocs/op
BenchmarkStringSMCWithJsoniter-8              	 1228988	       975 ns/op	    3256 B/op	       8 allocs/op
BenchmarkListStringJson-8                     	    1279	   1072942 ns/op	  226192 B/op	     213 allocs/op
BenchmarkListStringJsoniter-8                 	    2227	    533405 ns/op	  229326 B/op	     311 allocs/op
BenchmarkListStringSMC-8                      	    7064	    162927 ns/op	  547578 B/op	    1013 allocs/op
BenchmarkListStringSMCWithJsoniter-8          	    7507	    164071 ns/op	  547578 B/op	    1013 allocs/op
BenchmarkListOfMapJsoniter-8                  	    2864	    417876 ns/op	   82581 B/op	    2511 allocs/op
BenchmarkListOfMapSMC-8                       	    1694	    710615 ns/op	  311791 B/op	   12313 allocs/op
BenchmarkListOfMapSMCWithJsoniter-8           	    2504	    475963 ns/op	  115041 B/op	    3213 allocs/op
PASS
ok  	github.com/firmanmm/gosmc	35.885s
```

As you can see. This Simple Message Codec provides higher throughput in certain usecase compared to `jsoniter` and `json`. However, you can also see that this message codec may also take higher memory compared to `jsoniter`. So pick your choice between speed and memory. If you want to get the best of both world you can use `smc with jsoniter` variant which use `jsoniter` to handle `map` and `struct`. I always open if you want to improve it or maybe you want to create your own version and need some assistance.

## Todo
- Improve Map performance (Currently convert to List to handle the map)
- Make example

## Note
I highly recommend that you use `jsoniter` as that is more battle tested than this.
