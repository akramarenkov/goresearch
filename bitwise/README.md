# Comparison of the performance of bitwise operations with regular ones in cases where they can replace them

Running example:

```bash
go test -bench=.* -benchmem ./...
```

Result example:

```bash
goos: linux
goarch: amd64
pkg: github.com/akramarenkov/goresearch/bitwise
cpu: AMD Ryzen 5 3600 6-Core Processor              
BenchmarkAnd2-12            1000000000          0.6177 ns/op        0 B/op        0 allocs/op
BenchmarkAnd2Bitwise-12     1000000000          0.4585 ns/op        0 B/op        0 allocs/op
BenchmarkAnd3-12            1000000000          0.8610 ns/op        0 B/op        0 allocs/op
BenchmarkAnd3Bitwise-12     1000000000          0.7534 ns/op        0 B/op        0 allocs/op
BenchmarkAnd4-12            1000000000          1.079 ns/op         0 B/op        0 allocs/op
BenchmarkAnd4Bitwise-12     1000000000          0.7567 ns/op        0 B/op        0 allocs/op
BenchmarkAnd5-12             677626140          1.476 ns/op         0 B/op        0 allocs/op
BenchmarkAnd5Bitwise-12     1000000000          0.9075 ns/op        0 B/op        0 allocs/op
PASS
ok   github.com/akramarenkov/goresearch/bitwise 9.226s
```

Running example without optimizations:

```bash
go test -bench=.* -benchmem -gcflags=-N ./...
```

Result example without optimizations:

```bash
goos: linux
goarch: amd64
pkg: github.com/akramarenkov/goresearch/bitwise
cpu: AMD Ryzen 5 3600 6-Core Processor              
BenchmarkAnd2-12            493059032          2.474 ns/op        0 B/op        0 allocs/op
BenchmarkAnd2Bitwise-12     830754897          1.569 ns/op        0 B/op        0 allocs/op
BenchmarkAnd3-12            315412124          3.483 ns/op        0 B/op        0 allocs/op
BenchmarkAnd3Bitwise-12     735718498          1.516 ns/op        0 B/op        0 allocs/op
BenchmarkAnd4-12            255215498          4.544 ns/op        0 B/op        0 allocs/op
BenchmarkAnd4Bitwise-12     668503632          1.516 ns/op        0 B/op        0 allocs/op
BenchmarkAnd5-12            195563047          6.274 ns/op        0 B/op        0 allocs/op
BenchmarkAnd5Bitwise-12     635011461          1.847 ns/op        0 B/op        0 allocs/op
PASS
ok   github.com/akramarenkov/goresearch/bitwise 12.770s
```
