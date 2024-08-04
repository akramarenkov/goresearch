# Benchmark research

Running example:

```bash
go test -bench=.* -benchmem ./...
```

Result example:

```bash
goos: linux
goarch: amd64
pkg: github.com/akramarenkov/goresearch/benchmark
cpu: AMD Ryzen 5 3600 6-Core Processor              
BenchmarkMakeWithoutNLoop-12                              1000000000          0.0000002 ns/op     0 B/op           0 allocs/op
BenchmarkMakeWithoutNLoopWithAssignment-12                1000000000          0.0000004 ns/op     0 B/op           0 allocs/op
BenchmarkMakeWithNLoop-12                                 1000000000          0.2665 ns/op        0 B/op           0 allocs/op
BenchmarkMakeWithNLoopWithAssignment-12                   2309276           510.8 ns/op           0 B/op           0 allocs/op
BenchmarkMakeWithNLoopScopeEscape-12                      4861186           251.4 ns/op        1024 B/op           1 allocs/op
BenchmarkMakeWithNLoopWithAssignmentScopeEscape-12        1754312           692.7 ns/op        1024 B/op           1 allocs/op
BenchmarkMakeWithNLoopSizeFromEnv-12                      4692160           261.2 ns/op        1024 B/op           1 allocs/op
BenchmarkMakeWithNLoopWithAssignmentSizeFromEnv-12        1758030           704.8 ns/op        1024 B/op           1 allocs/op
PASS
ok      github.com/akramarenkov/goresearch/benchmark    8.826s
```

Running example without optimizations:

```bash
go test -bench=.* -benchmem -gcflags=-N ./...
```

Result example without optimizations:

```bash
goos: linux
goarch: amd64
pkg: github.com/akramarenkov/goresearch/benchmark
cpu: AMD Ryzen 5 3600 6-Core Processor              
BenchmarkMakeWithoutNLoop-12                              1000000000         0.0000003 ns/op     0 B/op           0 allocs/op
BenchmarkMakeWithoutNLoopWithAssignment-12                1000000000         0.0000030 ns/op     0 B/op           0 allocs/op
BenchmarkMakeWithNLoop-12                                 49350010          24.72 ns/op          0 B/op           0 allocs/op
BenchmarkMakeWithNLoopWithAssignment-12                   588489          2037 ns/op             0 B/op           0 allocs/op
BenchmarkMakeWithNLoopScopeEscape-12                      4471000          261.4 ns/op        1024 B/op           1 allocs/op
BenchmarkMakeWithNLoopWithAssignmentScopeEscape-12        444516          2755 ns/op          1024 B/op           1 allocs/op
BenchmarkMakeWithNLoopSizeFromEnv-12                      4556365          258.9 ns/op        1024 B/op           1 allocs/op
BenchmarkMakeWithNLoopWithAssignmentSizeFromEnv-12        377108          2720 ns/op          1024 B/op           1 allocs/op
PASS
ok      github.com/akramarenkov/goresearch/benchmark    9.186s
```
