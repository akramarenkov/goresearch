# Performance comparison of assignment and append in slice

Current package is a **main** package which was originally intended to be the only one.

Package **reverse**: the difference with the **main** package is the reverse order of writing the Benchmark* functions.

Running example:

```bash
GORESEARCH_SIZE=1000 go test -bench=.* -benchmem ./...
```

Result example:

```bash
goos: linux
goarch: amd64
pkg: github.com/akramarenkov/goresearch/assign
cpu: AMD Ryzen 5 3600 6-Core Processor              
BenchmarkAssignmentSlice-12         4833964           255.7 ns/op           0 B/op           0 allocs/op
BenchmarkAppendSlice-12             3513681           322.7 ns/op           0 B/op           0 allocs/op
PASS
ok      github.com/akramarenkov/goresearch/assign    2.976s
goos: linux
goarch: amd64
pkg: github.com/akramarenkov/goresearch/assign/reverse
cpu: AMD Ryzen 5 3600 6-Core Processor              
BenchmarkAppendSlice-12             3730080           330.9 ns/op           0 B/op           0 allocs/op
BenchmarkAssignmentSlice-12         2292518           492.5 ns/op           0 B/op           0 allocs/op
PASS
ok      github.com/akramarenkov/goresearch/assign/reverse    3.226s
```

Running example without optimizations:

```bash
GORESEARCH_SIZE=1000 go test -bench=.* -benchmem -gcflags=-N ./...
```

Result example without optimizations:

```bash
goos: linux
goarch: amd64
pkg: github.com/akramarenkov/goresearch/assign
cpu: AMD Ryzen 5 3600 6-Core Processor              
BenchmarkAssignmentSlice-12          819504          1516 ns/op           0 B/op           0 allocs/op
BenchmarkAppendSlice-12              754014          1737 ns/op           0 B/op           0 allocs/op
PASS
ok      github.com/akramarenkov/goresearch/assign    3.807s
goos: linux
goarch: amd64
pkg: github.com/akramarenkov/goresearch/assign/reverse
cpu: AMD Ryzen 5 3600 6-Core Processor              
BenchmarkAppendSlice-12              677061          1558 ns/op           0 B/op           0 allocs/op
BenchmarkAssignmentSlice-12          766302          1318 ns/op           0 B/op           0 allocs/op
PASS
ok      github.com/akramarenkov/goresearch/assign/reverse    2.105s
```
