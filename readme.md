# Go Simple Message Codec (SMC)

### Preface
This is a simple message codec `rewritten` based on version that is used on my `sync-mq` library. I decided to share it with you guys since i think it's ready to be used by the public. Feel free to improve this repository.

# About
Well this is a message codec. It's mainly used to pack and unpack together a bunch of data structure into an array of byte to be transported via network. You can think of it like json marshal and unmarshal but friendlier to machine instead of human. This is a work in progress version that may change in the future. However, if you want to use it right now, i suggest you to fork this project so there is no breaking change in the future. (PS: I do it often as i see fit, so I highly recommend that). 

# Benchmark

Well this is the comparison of `smc` against `json`, `jsoniter` and `smc backed with jsoniter` :

```
goos: windows
goarch: amd64
pkg: github.com/firmanmm/gosmc
BenchmarkArrayOfByteJson-8                    	    9231	    126300 ns/op	   28717 B/op	       6 allocs/op
BenchmarkArrayOfByteJsoniter-8                	   23452	     50819 ns/op	   40883 B/op	       5 allocs/op
BenchmarkArrayOfByteSMC-8                     	  522358	      2333 ns/op	   10371 B/op	       6 allocs/op
BenchmarkArrayOfByteSMCWithJsoniter-8         	  501392	      2343 ns/op	   10371 B/op	       6 allocs/op
BenchmarkNestedArrayOfByteJson-8              	     954	   1244085 ns/op	  303522 B/op	      43 allocs/op
BenchmarkNestedArrayOfByteJsoniter-8          	    2227	    538765 ns/op	  438760 B/op	     141 allocs/op
BenchmarkNestedArrayOfByteSMC-8               	   22701	     52590 ns/op	  120543 B/op	     317 allocs/op
BenchmarkNestedArrayOfByteSMCWithJsoniter-8   	   22701	     52764 ns/op	  120495 B/op	     317 allocs/op
BenchmarkInterfaceMapJsoniter-8               	  250674	      4400 ns/op	     775 B/op	      25 allocs/op
BenchmarkInterfaceMapSMC-8                    	  300546	      4072 ns/op	    1080 B/op	      53 allocs/op
BenchmarkInterfaceMapSMCWithJsoniter-8        	  245378	      4788 ns/op	     912 B/op	      29 allocs/op
BenchmarkDeepInterfaceMapJsoniter-8           	    2673	    443644 ns/op	   80891 B/op	    2625 allocs/op
BenchmarkDeepInterfaceMapSMC-8                	    2673	    444005 ns/op	  114187 B/op	    5853 allocs/op
BenchmarkDeepInterfaceMapSMCWithJsoniter-8    	    2671	    446579 ns/op	   89261 B/op	    2630 allocs/op
BenchmarkStringJson-8                         	  120351	      9961 ns/op	    2385 B/op	       5 allocs/op
BenchmarkStringJsoniter-8                     	  210970	      5682 ns/op	    2210 B/op	       4 allocs/op
BenchmarkStringSMC-8                          	 1000000	      1161 ns/op	    3281 B/op	       8 allocs/op
BenchmarkStringSMCWithJsoniter-8              	 1000000	      1142 ns/op	    3281 B/op	       8 allocs/op
BenchmarkListStringJson-8                     	    1293	    937167 ns/op	  222781 B/op	     213 allocs/op
BenchmarkListStringJsoniter-8                 	    2227	    531582 ns/op	  226823 B/op	     311 allocs/op
BenchmarkListStringSMC-8                      	   10000	    120976 ns/op	  334217 B/op	     607 allocs/op
BenchmarkListStringSMCWithJsoniter-8          	   10000	    118387 ns/op	  333851 B/op	     607 allocs/op
BenchmarkListOfMapJsoniter-8                  	    2864	    415439 ns/op	   82554 B/op	    2511 allocs/op
BenchmarkListOfMapSMC-8                       	    3163	    385953 ns/op	  105739 B/op	    5107 allocs/op
BenchmarkListOfMapSMCWithJsoniter-8           	    2559	    468085 ns/op	   91100 B/op	    2807 allocs/op
PASS
ok  	github.com/firmanmm/gosmc	32.629s
```

Comparing the output generated and we get :
```
=== RUN   TestSizeComparison
--- PASS: TestSizeComparison (0.01s)
    message_codec_test.go:252: Size Comparison :
    message_codec_test.go:253: Jsoniter : 92001
    message_codec_test.go:254: SMC : 74004
PASS
ok      github.com/firmanmm/gosmc       0.536s
```
And if we add slice of byte in to the mix then we get : 
```
=== RUN   TestSizeComparisonWithArrayOfByte
--- PASS: TestSizeComparisonWithArrayOfByte (0.01s)
    message_codec_test.go:282: Size Comparison :
    message_codec_test.go:283: Jsoniter : 197001
    message_codec_test.go:284: SMC : 155004
PASS
ok      github.com/firmanmm/gosmc       0.513s
```
Well, it clearly gives lower output size compared to `jsoniter` with or without having `slice of byte` in the map entry.

As you can see. This Simple Message Codec provides higher throughput in certain usecase compared to `jsoniter` and `json`. However, you can also see that this message codec may also take higher memory compared to `jsoniter`. So pick your choice between speed and memory. If you want to get the best of both world you can use `smc with jsoniter` variant which use `jsoniter` to handle `map` and `struct`. I always open if you want to improve it or maybe you want to create your own version and need some assistance.

## Behavioural Note
When using `jsoniter`, the decoded `map` version will be in the `map[string]interface{}` format where all keys will be converted into `string`. It also apply to codec instantiated with `NewSimpleMessageCodecWithJsoniter`. 

When using `NewSimpleMessageCodec` or `pure` implementation, the decoded `map` will become `map[interface{}]interface{}` and will maintain their original data type. So `int` key will remain `int` and not converted to `string` like `jsoniter`.

## Todo
- Improve Map performance (Need to be better than jsoniter)
- Struct support
- Make example

## Note
I highly recommend that you use `jsoniter` as that is more battle tested than this. Also `jsoniter` is easier to read than this (Good luck if you want to read the output when using SMC).
