# Code Generator Section
## Note
Please ignore this folder, this folder is commited in case my internal tool's code base gone wrong
Atleast there is backup of generated repetitive code in here

## Benchmark
Benchmark was performed to see any benefit when using generated code. Below are the benchmark result
```
goos: windows
goarch: amd64
pkg: github.com/firmanmm/gosmc/encoder
BenchmarkListEncoder-8                    704304              1694 ns/op             744 B/op         26 allocs/op
BenchmarkListInterfaceEncoder-8           802149              1503 ns/op             696 B/op         23 allocs/op
BenchmarkMapEncoder-8                     135192              8704 ns/op            2615 B/op        106 allocs/op
BenchmarkCommonMapEncoder-8               185107              6404 ns/op            2007 B/op         80 allocs/op
PASS
ok      github.com/firmanmm/gosmc/encoder       6.281s
```

## Conclusion
There is noticable increase in throughput when using auto generated code.
However it's only good for List only,
For map, things got worse since the generated code hits 4000 LOC which is unmaintainable, 
but if you decided that it is worth the performance increase then be my guest.