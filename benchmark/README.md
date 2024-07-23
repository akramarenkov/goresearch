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
BenchmarkMakeWithoutNLoopWithAssignment-12                1000000000          0.0000007 ns/op     0 B/op           0 allocs/op
BenchmarkMakeWithNLoop-12                                 1000000000          0.2396 ns/op        0 B/op           0 allocs/op
BenchmarkMakeWithNLoopWithAssignment-12                   2453742           498.1 ns/op           0 B/op           0 allocs/op
BenchmarkMakeWithNLoopScopeEscape-12                      4780209           257.4 ns/op        1024 B/op           1 allocs/op
BenchmarkMakeWithNLoopWithAssignmentScopeEscape-12        1786213           678.0 ns/op        1024 B/op           1 allocs/op
BenchmarkMakeWithNLoopSizeFromEnv-12                      4578804           260.8 ns/op        1024 B/op           1 allocs/op
BenchmarkMakeWithNLoopWithAssignmentSizeFromEnv-12        1752130           671.2 ns/op        1024 B/op           1 allocs/op
BenchmarkAssignmentSlice-12                               4481353           253.0 ns/op           0 B/op           0 allocs/op
BenchmarkAppendSlice-12                                   3280669           342.7 ns/op           0 B/op           0 allocs/op
PASS
ok      github.com/akramarenkov/goresearch/benchmark    11.624s
```

Running example without optimizations:

```bash
GORESEARCH_SIZE=1000 go test -bench=.* -benchmem -gcflags=-N ./...
```

Result example without optimizations:

```bash
goos: linux
goarch: amd64
pkg: github.com/akramarenkov/goresearch/benchmark
cpu: AMD Ryzen 5 3600 6-Core Processor              
BenchmarkMakeWithoutNLoop-12                              1000000000         0.0000002 ns/op     0 B/op           0 allocs/op
BenchmarkMakeWithoutNLoopWithAssignment-12                1000000000         0.0000022 ns/op     0 B/op           0 allocs/op
BenchmarkMakeWithNLoop-12                                 49878760          24.10 ns/op          0 B/op           0 allocs/op
BenchmarkMakeWithNLoopWithAssignment-12                   588502          1959 ns/op             0 B/op           0 allocs/op
BenchmarkMakeWithNLoopScopeEscape-12                      4711892          264.0 ns/op        1024 B/op           1 allocs/op
BenchmarkMakeWithNLoopWithAssignmentScopeEscape-12        446679          2640 ns/op          1024 B/op           1 allocs/op
BenchmarkMakeWithNLoopSizeFromEnv-12                      4550565          264.2 ns/op        1024 B/op           1 allocs/op
BenchmarkMakeWithNLoopWithAssignmentSizeFromEnv-12        414697          2628 ns/op          1024 B/op           1 allocs/op
BenchmarkAssignmentSlice-12                               804276          1293 ns/op             0 B/op           0 allocs/op
BenchmarkAppendSlice-12                                   652683          2038 ns/op             0 B/op           0 allocs/op
PASS
ok      github.com/akramarenkov/goresearch/benchmark    12.189s
```
