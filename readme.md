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
BenchmarkArrayOfByteJson-8                    	    9255	    123821 ns/op	   28664 B/op	       6 allocs/op
BenchmarkArrayOfByteJsoniter-8                	   24062	     50236 ns/op	   40865 B/op	       5 allocs/op
BenchmarkArrayOfByteSMC-8                     	  501296	      2389 ns/op	   10664 B/op	      15 allocs/op
BenchmarkArrayOfByteSMCWithJsoniter-8         	  462754	      2457 ns/op	   10664 B/op	      15 allocs/op
BenchmarkNestedArrayOfByteJson-8              	     985	   1217079 ns/op	  295947 B/op	      43 allocs/op
BenchmarkNestedArrayOfByteJsoniter-8          	    2312	    521975 ns/op	  430576 B/op	     141 allocs/op
BenchmarkNestedArrayOfByteSMC-8               	   10000	    107216 ns/op	  159554 B/op	    1742 allocs/op
BenchmarkNestedArrayOfByteSMCWithJsoniter-8   	   10000	    107516 ns/op	  159554 B/op	    1742 allocs/op
BenchmarkInterfaceMapJsoniter-8               	  273226	      4450 ns/op	     775 B/op	      25 allocs/op
BenchmarkInterfaceMapSMC-8                    	  101966	     11904 ns/op	    5941 B/op	     230 allocs/op
BenchmarkInterfaceMapSMCWithJsoniter-8        	  231220	      5236 ns/op	    1219 B/op	      38 allocs/op
BenchmarkDeepInterfaceMapJsoniter-8           	    2673	    448495 ns/op	   80888 B/op	    2625 allocs/op
BenchmarkDeepInterfaceMapSMC-8                	     865	   1378966 ns/op	  661729 B/op	   25830 allocs/op
BenchmarkDeepInterfaceMapSMCWithJsoniter-8    	    2614	    453658 ns/op	   89532 B/op	    2639 allocs/op
BenchmarkStringJson-8                         	  121497	      9875 ns/op	    2385 B/op	       5 allocs/op
BenchmarkStringJsoniter-8                     	  214731	      5638 ns/op	    2210 B/op	       4 allocs/op
BenchmarkStringSMC-8                          	  802202	      1461 ns/op	    3592 B/op	      17 allocs/op
BenchmarkStringSMCWithJsoniter-8              	  802154	      1468 ns/op	    3592 B/op	      17 allocs/op
BenchmarkListStringJson-8                     	    1292	    928657 ns/op	  222351 B/op	     213 allocs/op
BenchmarkListStringJsoniter-8                 	    2268	    524172 ns/op	  226412 B/op	     311 allocs/op
BenchmarkListStringSMC-8                      	     961	   1246370 ns/op	  606777 B/op	   23422 allocs/op
BenchmarkListStringSMCWithJsoniter-8          	    2269	    587235 ns/op	  135572 B/op	    4322 allocs/op
BenchmarkListOfMapJsoniter-8                  	    2864	    422055 ns/op	   82568 B/op	    2511 allocs/op
BenchmarkListOfMapSMC-8                       	     946	   1248247 ns/op	  606774 B/op	   23422 allocs/op
BenchmarkListOfMapSMCWithJsoniter-8           	    2269	    535823 ns/op	  135570 B/op	    4322 allocs/op
PASS
ok  	github.com/firmanmm/gosmc	32.317s
```

As you can see. This Simple Message Codec provides higher throughput in certain usecase compared to `jsoniter` and `json`. However, you can also see that this message codec can also take significantly higher memory compared to both of them. So pick your choice between speed and memory. If you want to get the best of both world you can use `smc with jsoniter` variant which use `jsoniter` to handle `map` and `struct`. I always open if you want to improve it or maybe you want to create your own version and need some assistance.

## Note
I highly recommend that you use `jsoniter` as that is more battle tested than this.
