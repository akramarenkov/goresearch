# Performance comparison of assignment and append in slice

The code in different packages is almost identical. They differ only in the presence/absence of calls to the b.ResetTimer()/b.StopTimer() functions and the order in which the Benchmark* functions are written.

Running example:

```bash
GORESEARCH_SIZE=1000 go test -bench=.* -benchmem ./...
```

Result example:

```bash
goos: linux
goarch: amd64
pkg: github.com/akramarenkov/goresearch/assignment
cpu: AMD Ryzen 5 3600 6-Core Processor              
BenchmarkAssignmentSlice-12         2450546           502.5 ns/op           0 B/op           0 allocs/op
BenchmarkAppendSlice-12             3272944           351.9 ns/op           0 B/op           0 allocs/op
PASS
ok      github.com/akramarenkov/goresearch/assignment    3.258s
goos: linux
goarch: amd64
pkg: github.com/akramarenkov/goresearch/assignment/five
cpu: AMD Ryzen 5 3600 6-Core Processor              
BenchmarkAppendSlice-12             3375187           357.8 ns/op           0 B/op           0 allocs/op
BenchmarkAssignmentSlice-12         2427840           505.8 ns/op           0 B/op           0 allocs/op
PASS
ok      github.com/akramarenkov/goresearch/assignment/five    3.302s
goos: linux
goarch: amd64
pkg: github.com/akramarenkov/goresearch/assignment/four
cpu: AMD Ryzen 5 3600 6-Core Processor              
BenchmarkAppendSlice-12             2433134           495.3 ns/op           0 B/op           0 allocs/op
BenchmarkAssignmentSlice-12         4353081           249.1 ns/op           0 B/op           0 allocs/op
PASS
ok      github.com/akramarenkov/goresearch/assignment/four    3.073s
goos: linux
goarch: amd64
pkg: github.com/akramarenkov/goresearch/assignment/one
cpu: AMD Ryzen 5 3600 6-Core Processor              
BenchmarkAssignmentSlice-12         4792650           255.7 ns/op           0 B/op           0 allocs/op
BenchmarkAppendSlice-12             3127789           348.5 ns/op           0 B/op           0 allocs/op
PASS
ok      github.com/akramarenkov/goresearch/assignment/one    2.967s
goos: linux
goarch: amd64
pkg: github.com/akramarenkov/goresearch/assignment/three
cpu: AMD Ryzen 5 3600 6-Core Processor              
BenchmarkAssignmentSlice-12         4848877           255.2 ns/op           0 B/op           0 allocs/op
BenchmarkAppendSlice-12             3477580           321.5 ns/op           0 B/op           0 allocs/op
PASS
ok      github.com/akramarenkov/goresearch/assignment/three    2.963s
goos: linux
goarch: amd64
pkg: github.com/akramarenkov/goresearch/assignment/two
cpu: AMD Ryzen 5 3600 6-Core Processor              
BenchmarkAssignmentSlice-12         4834138           254.1 ns/op           0 B/op           0 allocs/op
BenchmarkAppendSlice-12             3223677           354.6 ns/op           0 B/op           0 allocs/op
PASS
ok      github.com/akramarenkov/goresearch/assignment/two    3.005s
```

Running example without optimizations:

```bash
GORESEARCH_SIZE=1000 go test -bench=.* -benchmem -gcflags=-N ./...
```

Result example without optimizations:

```bash
goos: linux
goarch: amd64
pkg: github.com/akramarenkov/goresearch/assignment
cpu: AMD Ryzen 5 3600 6-Core Processor              
BenchmarkAssignmentSlice-12          937608          1741 ns/op           0 B/op           0 allocs/op
BenchmarkAppendSlice-12              594217          2017 ns/op           0 B/op           0 allocs/op
PASS
ok      github.com/akramarenkov/goresearch/assignment    3.819s
goos: linux
goarch: amd64
pkg: github.com/akramarenkov/goresearch/assignment/five
cpu: AMD Ryzen 5 3600 6-Core Processor              
BenchmarkAppendSlice-12              607272          1954 ns/op           0 B/op           0 allocs/op
BenchmarkAssignmentSlice-12          900616          1532 ns/op           0 B/op           0 allocs/op
PASS
ok      github.com/akramarenkov/goresearch/assignment/five    2.607s
goos: linux
goarch: amd64
pkg: github.com/akramarenkov/goresearch/assignment/four
cpu: AMD Ryzen 5 3600 6-Core Processor              
BenchmarkAppendSlice-12              652657          1872 ns/op           0 B/op           0 allocs/op
BenchmarkAssignmentSlice-12          795429          1518 ns/op           0 B/op           0 allocs/op
PASS
ok      github.com/akramarenkov/goresearch/assignment/four    3.223s
goos: linux
goarch: amd64
pkg: github.com/akramarenkov/goresearch/assignment/one
cpu: AMD Ryzen 5 3600 6-Core Processor              
BenchmarkAssignmentSlice-12          921229          1277 ns/op           0 B/op           0 allocs/op
BenchmarkAppendSlice-12              590760          2021 ns/op           0 B/op           0 allocs/op
PASS
ok      github.com/akramarenkov/goresearch/assignment/one    3.192s
goos: linux
goarch: amd64
pkg: github.com/akramarenkov/goresearch/assignment/three
cpu: AMD Ryzen 5 3600 6-Core Processor              
BenchmarkAssignmentSlice-12          939184          1534 ns/op           0 B/op           0 allocs/op
BenchmarkAppendSlice-12              634262          1726 ns/op           0 B/op           0 allocs/op
PASS
ok      github.com/akramarenkov/goresearch/assignment/three    3.446s
goos: linux
goarch: amd64
pkg: github.com/akramarenkov/goresearch/assignment/two
cpu: AMD Ryzen 5 3600 6-Core Processor              
BenchmarkAssignmentSlice-12          942877          1291 ns/op           0 B/op           0 allocs/op
BenchmarkAppendSlice-12              588236          1943 ns/op           0 B/op           0 allocs/op
PASS
ok      github.com/akramarenkov/goresearch/assignment/two    3.031s
```
