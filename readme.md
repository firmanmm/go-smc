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
BenchmarkArrayOfByteJson-8                    	    8566	    124811 ns/op	   28672 B/op	       6 allocs/op
BenchmarkArrayOfByteJsoniter-8                	   23454	     50944 ns/op	   40853 B/op	       5 allocs/op
BenchmarkArrayOfByteSMC-8                     	  445623	      2511 ns/op	   10664 B/op	      15 allocs/op
BenchmarkArrayOfByteSMCWithJsoniter-8         	  481268	      2489 ns/op	   10664 B/op	      15 allocs/op
BenchmarkNestedArrayOfByteJson-8              	     985	   1222145 ns/op	  296522 B/op	      43 allocs/op
BenchmarkNestedArrayOfByteJsoniter-8          	    2227	    532925 ns/op	  435737 B/op	     141 allocs/op
BenchmarkNestedArrayOfByteSMC-8               	   10000	    113500 ns/op	  157730 B/op	    1740 allocs/op
BenchmarkNestedArrayOfByteSMCWithJsoniter-8   	   10000	    113095 ns/op	  157730 B/op	    1740 allocs/op
BenchmarkInterfaceMapJsoniter-8               	  273235	      4493 ns/op	     775 B/op	      25 allocs/op
BenchmarkInterfaceMapSMC-8                    	  102837	     11648 ns/op	    5845 B/op	     230 allocs/op
BenchmarkInterfaceMapSMCWithJsoniter-8        	  231387	      5289 ns/op	    1219 B/op	      38 allocs/op
BenchmarkDeepInterfaceMapJsoniter-8           	    2671	    451061 ns/op	   80821 B/op	    2625 allocs/op
BenchmarkDeepInterfaceMapSMC-8                	     871	   1362599 ns/op	  650435 B/op	   25830 allocs/op
BenchmarkDeepInterfaceMapSMCWithJsoniter-8    	    2673	    457810 ns/op	   89459 B/op	    2639 allocs/op
BenchmarkStringJson-8                         	  117924	      9937 ns/op	    2385 B/op	       5 allocs/op
BenchmarkStringJsoniter-8                     	  214732	      5671 ns/op	    2210 B/op	       4 allocs/op
BenchmarkStringSMC-8                          	  802111	      1510 ns/op	    3592 B/op	      17 allocs/op
BenchmarkStringSMCWithJsoniter-8              	  802089	      1510 ns/op	    3592 B/op	      17 allocs/op
BenchmarkListStringJson-8                     	    1279	    941160 ns/op	  221761 B/op	     213 allocs/op
BenchmarkListStringJsoniter-8                 	    2228	    531790 ns/op	  229545 B/op	     311 allocs/op
BenchmarkListStringSMC-8                      	    6672	    175634 ns/op	  360771 B/op	    2120 allocs/op
BenchmarkListStringSMCWithJsoniter-8          	    7076	    175905 ns/op	  360771 B/op	    2120 allocs/op
BenchmarkListOfMapJsoniter-8                  	    2864	    420314 ns/op	   82467 B/op	    2511 allocs/op
BenchmarkListOfMapSMC-8                       	     993	   1218326 ns/op	  595351 B/op	   23420 allocs/op
BenchmarkListOfMapSMCWithJsoniter-8           	    2227	    538283 ns/op	  133746 B/op	    4320 allocs/op
PASS
ok  	github.com/firmanmm/gosmc	31.887s
```

As you can see. This Simple Message Codec provides higher throughput in certain usecase compared to `jsoniter` and `json`. However, you can also see that this message codec can also take significantly higher memory compared to both of them. So pick your choice between speed and memory. If you want to get the best of both world you can use `smc with jsoniter` variant which use `jsoniter` to handle `map` and `struct`. I always open if you want to improve it or maybe you want to create your own version and need some assistance.

## Note
I highly recommend that you use `jsoniter` as that is more battle tested than this.
