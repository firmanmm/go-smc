# Simple Message Codec Optimization

## Description
There is a litte optimization you can do when passing data to SMC. It can be achieved by just passing the right data type. If the passed data type is natively supported (refer to the table below) then we can avoid using reflection. By just bypassing the reflection stuff we can net a pretty big performance gain.

| Type  | Converted |
| :------------- | :------------- |
| bool  | bool |
| []byte  | []byte |
| string | string |
| int, int8, int16, int32, int64 | int |
| uint, uint8, uint16, uint32, uint64 | uint |
| float32, float64  | float64 |
| []interface{} | []interface{} |
| map[string]interface{}, map[interface{}]interface{}  | map[interface{}]interface{} |

The rule is as follow, if it's not in the table then it needs to find which one has the closest value. The `find` stuff is using reflection and it can be quite consuming.
That means `map[string]string` is not the same with `map[string]interface{}` or even `map[interface{}]interface{}`. 
## Benchmark
Let's see the number shall we?
```
goos: windows
goarch: amd64
pkg: github.com/firmanmm/gosmc/example/optimization
BenchmarkMapStringString-8               	    1314	    905254 ns/op	  317328 B/op	   10038 allocs/op
BenchmarkMapStringInterface-8            	    1656	    721666 ns/op	  284534 B/op	    8036 allocs/op
BenchmarkMapInterfaceInterface-8         	    1647	    751338 ns/op	  283980 B/op	    8036 allocs/op
BenchmarkMapNestedStringString-8         	     609	   2029757 ns/op	  881336 B/op	   21040 allocs/op
BenchmarkMapNestedInterfaceInterface-8   	     760	   1567411 ns/op	  702200 B/op	   16038 allocs/op
PASS
ok  	github.com/firmanmm/gosmc/example/optimization	9.488s
```

## Conclusion
As you can see that it is much faster to process `map[string]interface{}` and `map[interface{}]interface{}` than `map[string]string`. It's because there is manual implementation for that 2 types while `map[string]string` has to be encoded by using reflection to determine their type at runtime. The nested version also share the same faith. As long as the underlying type is supported then we can avoid reflection ~~*if only Golang has generic*~~.