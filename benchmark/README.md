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
BenchmarkMakeWithoutNLoop-12                              1000000000         0.0000002 ns/op     0 B/op           0 allocs/op
BenchmarkMakeWithoutNLoopWithAssignment-12                1000000000         0.0000006 ns/op     0 B/op           0 allocs/op
BenchmarkMakeWithNLoop-12                                 1000000000         0.2676 ns/op        0 B/op           0 allocs/op
BenchmarkMakeWithNLoopWithAssignment-12                   4442838          258.7 ns/op           0 B/op           0 allocs/op
BenchmarkMakeWithNLoopScopeEscape-12                      4609045          259.5 ns/op        1024 B/op           1 allocs/op
BenchmarkMakeWithNLoopWithAssignmentScopeEscape-12        1721224          670.7 ns/op        1024 B/op           1 allocs/op
BenchmarkMakeWithNLoopSizeFromEnv-12                      4796232          258.0 ns/op        1024 B/op           1 allocs/op
BenchmarkMakeWithNLoopWithAssignmentSizeFromEnv-12        883053          1146 ns/op          1024 B/op           1 allocs/op
PASS
ok      github.com/akramarenkov/goresearch/benchmark    7.578s
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
BenchmarkMakeWithoutNLoop-12                              1000000000         0.0000004 ns/op     0 B/op           0 allocs/op
BenchmarkMakeWithoutNLoopWithAssignment-12                1000000000         0.0000023 ns/op     0 B/op           0 allocs/op
BenchmarkMakeWithNLoop-12                                 48516979          24.40 ns/op          0 B/op           0 allocs/op
BenchmarkMakeWithNLoopWithAssignment-12                   576064          1988 ns/op             0 B/op           0 allocs/op
BenchmarkMakeWithNLoopScopeEscape-12                      4454593          260.8 ns/op        1024 B/op           1 allocs/op
BenchmarkMakeWithNLoopWithAssignmentScopeEscape-12        356882          2967 ns/op          1024 B/op           1 allocs/op
BenchmarkMakeWithNLoopSizeFromEnv-12                      4759490          262.6 ns/op        1024 B/op           1 allocs/op
BenchmarkMakeWithNLoopWithAssignmentSizeFromEnv-12        419067          2695 ns/op          1024 B/op           1 allocs/op
PASS
ok      github.com/akramarenkov/goresearch/benchmark    8.257s
```
