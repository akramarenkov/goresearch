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
BenchmarkMakeWithoutNLoop-12                1000000000         0.0000004 ns/op       0 B/op           0 allocs/op
BenchmarkMakeWithNLoop-12                    4699849           252.5 ns/op           0 B/op           0 allocs/op
BenchmarkMakeWithNLoopSizeFromEnv-12         1779364           678.3 ns/op        1024 B/op           1 allocs/op
BenchmarkAssignmentSlice-12                  4781977           256.6 ns/op           0 B/op           0 allocs/op
BenchmarkAppendSlice-12                      2457417           485.3 ns/op           0 B/op           0 allocs/op
PASS

```
