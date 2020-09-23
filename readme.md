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
BenchmarkArrayOfByteJson-8                    	    8575	    125611 ns/op	   28674 B/op	       6 allocs/op
BenchmarkArrayOfByteJsoniter-8                	   23362	     50930 ns/op	   40871 B/op	       5 allocs/op
BenchmarkArrayOfByteSMC-8                     	  601572	      1991 ns/op	   10328 B/op	       6 allocs/op
BenchmarkArrayOfByteSMCWithJsoniter-8         	  633214	      1925 ns/op	   10328 B/op	       6 allocs/op
BenchmarkNestedArrayOfByteJson-8              	     985	   1231257 ns/op	  305718 B/op	      43 allocs/op
BenchmarkNestedArrayOfByteJsoniter-8          	    2227	    530238 ns/op	  438761 B/op	     141 allocs/op
BenchmarkNestedArrayOfByteSMC-8               	   10000	    104623 ns/op	  332489 B/op	     634 allocs/op
BenchmarkNestedArrayOfByteSMCWithJsoniter-8   	   12691	     93124 ns/op	  332489 B/op	     634 allocs/op
BenchmarkInterfaceMapJsoniter-8               	  273681	      4450 ns/op	     775 B/op	      25 allocs/op
BenchmarkInterfaceMapSMC-8                    	  176852	      6982 ns/op	    2869 B/op	     123 allocs/op
BenchmarkInterfaceMapSMCWithJsoniter-8        	 1558557	       767 ns/op	     136 B/op	       8 allocs/op
BenchmarkDeepInterfaceMapJsoniter-8           	    2559	    455600 ns/op	   80858 B/op	    2625 allocs/op
BenchmarkDeepInterfaceMapSMC-8                	     871	   1401530 ns/op	 2228537 B/op	   13723 allocs/op
BenchmarkDeepInterfaceMapSMCWithJsoniter-8    	 1542577	       767 ns/op	     136 B/op	       8 allocs/op
BenchmarkStringJson-8                         	  121490	      9991 ns/op	    2385 B/op	       5 allocs/op
BenchmarkStringJsoniter-8                     	  211096	      5740 ns/op	    2210 B/op	       4 allocs/op
BenchmarkStringSMC-8                          	 1257272	       959 ns/op	    3256 B/op	       8 allocs/op
BenchmarkStringSMCWithJsoniter-8              	 1255970	       966 ns/op	    3256 B/op	       8 allocs/op
BenchmarkListStringJson-8                     	    1279	    935755 ns/op	  218879 B/op	     213 allocs/op
BenchmarkListStringJsoniter-8                 	    2227	    530238 ns/op	  230235 B/op	     311 allocs/op
BenchmarkListStringSMC-8                      	    7519	    164213 ns/op	  547610 B/op	    1014 allocs/op
BenchmarkListStringSMCWithJsoniter-8          	    7077	    163051 ns/op	  547611 B/op	    1014 allocs/op
BenchmarkListOfMapJsoniter-8                  	    2794	    427263 ns/op	   82602 B/op	    2511 allocs/op
BenchmarkListOfMapSMC-8                       	    1648	    722601 ns/op	  318224 B/op	   12514 allocs/op
BenchmarkListOfMapSMCWithJsoniter-8           	   13941	     86204 ns/op	   23505 B/op	    1014 allocs/op
PASS
ok  	github.com/firmanmm/gosmc	36.890s
```

As you can see. This Simple Message Codec provides higher throughput in certain usecase compared to `jsoniter` and `json`. However, you can also see that this message codec may also take higher memory compared to `jsoniter`. So pick your choice between speed and memory. If you want to get the best of both world you can use `smc with jsoniter` variant which use `jsoniter` to handle `map` and `struct`. I always open if you want to improve it or maybe you want to create your own version and need some assistance.

## Todo
- Improve Map performance (Currently convert to List to handle the map)
- Make example

## Note
I highly recommend that you use `jsoniter` as that is more battle tested than this.
