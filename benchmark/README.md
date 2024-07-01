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
BenchmarkMakeWithoutNLoop-12                              1000000000          0.0000002 ns/op     0 B/op           0 allocs/op
BenchmarkMakeWithoutNLoopWithAssignment-12                1000000000          0.0000004 ns/op     0 B/op           0 allocs/op
BenchmarkMakeWithNLoop-12                                 1000000000          0.2593 ns/op        0 B/op           0 allocs/op
BenchmarkMakeWithNLoopWithAssignment-12                   2310674           499.4 ns/op           0 B/op           0 allocs/op
BenchmarkMakeWithNLoopSizeFromEnv-12                      4736898           250.0 ns/op        1024 B/op           1 allocs/op
BenchmarkMakeWithNLoopSizeFromEnvWithAssignment-12        1766924           650.6 ns/op        1024 B/op           1 allocs/op
BenchmarkAssignmentSlice-12                               4440802           255.2 ns/op           0 B/op           0 allocs/op
BenchmarkAssignmentSliceUseCap-12                         4440178           251.4 ns/op           0 B/op           0 allocs/op
BenchmarkAssignmentSliceConsiderMake-12                   4227254           250.0 ns/op           0 B/op           0 allocs/op
BenchmarkAppendSlice-12                                   3175544           346.1 ns/op           0 B/op           0 allocs/op
BenchmarkAppendSliceUseCap-12                             2458405           506.1 ns/op           0 B/op           0 allocs/op
BenchmarkAppendSliceConsiderMake-12                       2119771           507.9 ns/op           0 B/op           0 allocs/op
PASS
ok      github.com/akramarenkov/goresearch/benchmark    14.263s
```
