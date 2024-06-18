# Benchmark research

Running example:

```bash
GORESEARCH_SIZE=1000 go test -bench=.* -benchmem ./...
```

Result example:

```bash
goos: linux
goarch: amd64
pkg: github.com/akramarenkov/goresearch/benchmark
cpu: AMD Ryzen 5 3600 6-Core Processor              
BenchmarkMakeWithoutNLoop-12                              1000000000         0.0000002 ns/op       0 B/op           0 allocs/op
BenchmarkMakeWithoutNLoopWithAssignment-12                1000000000         0.0000005 ns/op       0 B/op           0 allocs/op
BenchmarkMakeWithNLoop-12                                 1000000000         0.2555 ns/op          0 B/op           0 allocs/op
BenchmarkMakeWithNLoopWithAssignment-12                    2343388           496.3 ns/op           0 B/op           0 allocs/op
BenchmarkMakeWithNLoopSizeFromEnv-12                       4770296           254.8 ns/op        1024 B/op           1 allocs/op
BenchmarkMakeWithNLoopSizeFromEnvWithAssignment-12         1797872           659.7 ns/op        1024 B/op           1 allocs/op
BenchmarkAssignmentSlice-12                                4436733           252.5 ns/op           0 B/op           0 allocs/op
BenchmarkAssignmentSliceUseCap-12                          4410078           251.7 ns/op           0 B/op           0 allocs/op
BenchmarkAssignmentSliceConsiderMake-12                    4248307           252.1 ns/op           0 B/op           0 allocs/op
BenchmarkAppendSlice-12                                    3490934           344.6 ns/op           0 B/op           0 allocs/op
BenchmarkAppendSliceUseCap-12                              2445586           505.3 ns/op           0 B/op           0 allocs/op
BenchmarkAppendSliceConsiderMake-12                        2466327           499.0 ns/op           0 B/op           0 allocs/op
PASS
ok      github.com/akramarenkov/goresearch/benchmark    14.473s
```
