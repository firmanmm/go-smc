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
BenchmarkArrayOfByteJson-8                    	    9231	    130972 ns/op	   28671 B/op	       6 allocs/op
BenchmarkArrayOfByteJsoniter-8                	   22716	     53416 ns/op	   40882 B/op	       5 allocs/op
BenchmarkArrayOfByteSMC-8                     	  480618	      2262 ns/op	   10328 B/op	       6 allocs/op
BenchmarkArrayOfByteSMCWithJsoniter-8         	  445596	      2262 ns/op	   10328 B/op	       6 allocs/op
BenchmarkNestedArrayOfByteJson-8              	     946	   1281103 ns/op	  302443 B/op	      43 allocs/op
BenchmarkNestedArrayOfByteJsoniter-8          	    2187	    559220 ns/op	  441219 B/op	     142 allocs/op
BenchmarkNestedArrayOfByteSMC-8               	   10000	    104883 ns/op	  334282 B/op	     635 allocs/op
BenchmarkNestedArrayOfByteSMCWithJsoniter-8   	   10000	    105454 ns/op	  334283 B/op	     635 allocs/op
BenchmarkInterfaceMapJsoniter-8               	  267372	      4634 ns/op	     775 B/op	      25 allocs/op
BenchmarkInterfaceMapSMC-8                    	  152240	      7649 ns/op	    2998 B/op	     123 allocs/op
BenchmarkInterfaceMapSMCWithJsoniter-8        	  256008	      4768 ns/op	     888 B/op	      29 allocs/op
BenchmarkDeepInterfaceMapJsoniter-8           	    2672	    467265 ns/op	   80933 B/op	    2625 allocs/op
BenchmarkDeepInterfaceMapSMC-8                	     717	   1617705 ns/op	 2244626 B/op	   13723 allocs/op
BenchmarkDeepInterfaceMapSMCWithJsoniter-8    	    2617	    465408 ns/op	   89214 B/op	    2630 allocs/op
BenchmarkStringJson-8                         	  111438	     10375 ns/op	    2385 B/op	       5 allocs/op
BenchmarkStringJsoniter-8                     	  203958	      5875 ns/op	    2210 B/op	       4 allocs/op
BenchmarkStringSMC-8                          	 1000000	      1053 ns/op	    3256 B/op	       8 allocs/op
BenchmarkStringSMCWithJsoniter-8              	 1000000	      1069 ns/op	    3256 B/op	       8 allocs/op
BenchmarkListStringJson-8                     	    1239	    973822 ns/op	  223814 B/op	     213 allocs/op
BenchmarkListStringJsoniter-8                 	    2227	    570699 ns/op	  230461 B/op	     311 allocs/op
BenchmarkListStringSMC-8                      	    1604	    794489 ns/op	  332638 B/op	   12515 allocs/op
BenchmarkListStringSMCWithJsoniter-8          	    2383	    496786 ns/op	  116872 B/op	    3215 allocs/op
BenchmarkListOfMapJsoniter-8                  	    2796	    430642 ns/op	   82552 B/op	    2511 allocs/op
BenchmarkListOfMapSMC-8                       	    1503	    801502 ns/op	  332636 B/op	   12515 allocs/op
BenchmarkListOfMapSMCWithJsoniter-8           	    2359	    500570 ns/op	  116867 B/op	    3215 allocs/op
PASS
ok  	github.com/firmanmm/gosmc	31.536s
```

As you can see. This Simple Message Codec provides higher throughput in certain usecase compared to `jsoniter` and `json`. However, you can also see that this message codec can also take significantly higher memory compared to both of them. So pick your choice between speed and memory. If you want to get the best of both world you can use `smc with jsoniter` variant which use `jsoniter` to handle `map` and `struct`. I always open if you want to improve it or maybe you want to create your own version and need some assistance.

## Note
I highly recommend that you use `jsoniter` as that is more battle tested than this.
