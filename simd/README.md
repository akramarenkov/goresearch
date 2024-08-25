# Example of SIMD using

Running example:

```bash
go test -bench=.* -benchmem ./...
```

Result example:

```bash
goos: linux
goarch: amd64
pkg: github.com/akramarenkov/goresearch/simd
cpu: AMD Ryzen 5 3600 6-Core Processor              
BenchmarkSlowMultiplication_64-12            35475258         32.38 ns/op        0 B/op        0 allocs/op
BenchmarkSlowMultiplicationSIMD_64-12        25582794         47.29 ns/op        0 B/op        0 allocs/op
BenchmarkSlowMultiplication_128-12           17236124         68.49 ns/op        0 B/op        0 allocs/op
BenchmarkSlowMultiplicationSIMD_128-12       24155043         48.62 ns/op        0 B/op        0 allocs/op
BenchmarkSlowMultiplication_256-12            9351682        128.2 ns/op         0 B/op        0 allocs/op
BenchmarkSlowMultiplicationSIMD_256-12       21983127         52.47 ns/op        0 B/op        0 allocs/op
BenchmarkSlowMultiplication_512-12            4230076        252.6 ns/op         0 B/op        0 allocs/op
BenchmarkSlowMultiplicationSIMD_512-12       18295626         64.06 ns/op        0 B/op        0 allocs/op
BenchmarkSlowMultiplication_1024-12           2404146        498.9 ns/op         0 B/op        0 allocs/op
BenchmarkSlowMultiplicationSIMD_1024-12      11706397         98.60 ns/op        0 B/op        0 allocs/op
BenchmarkSlowMultiplication_65536-12            34680      32296 ns/op           0 B/op        0 allocs/op
BenchmarkSlowMultiplicationSIMD_65536-12       301449       4910 ns/op           0 B/op        0 allocs/op
PASS
ok   github.com/akramarenkov/goresearch/simd 23.617s
```

Running example without optimizations:

```bash
go test -bench=.* -benchmem -gcflags=-N ./...
```

Result example without optimizations:

```bash
goos: linux
goarch: amd64
pkg: github.com/akramarenkov/goresearch/simd
cpu: AMD Ryzen 5 3600 6-Core Processor              
BenchmarkSlowMultiplication_64-12            18685464        111.3 ns/op         0 B/op        0 allocs/op
BenchmarkSlowMultiplicationSIMD_64-12        24609580         46.95 ns/op        0 B/op        0 allocs/op
BenchmarkSlowMultiplication_128-12           10655684        111.8 ns/op         0 B/op        0 allocs/op
BenchmarkSlowMultiplicationSIMD_128-12       24303078         48.54 ns/op        0 B/op        0 allocs/op
BenchmarkSlowMultiplication_256-12            5016102        214.0 ns/op         0 B/op        0 allocs/op
BenchmarkSlowMultiplicationSIMD_256-12       22098590         53.73 ns/op        0 B/op        0 allocs/op
BenchmarkSlowMultiplication_512-12            2926927        423.9 ns/op         0 B/op        0 allocs/op
BenchmarkSlowMultiplicationSIMD_512-12       17910554         65.77 ns/op        0 B/op        0 allocs/op
BenchmarkSlowMultiplication_1024-12           1429669        833.8 ns/op         0 B/op        0 allocs/op
BenchmarkSlowMultiplicationSIMD_1024-12      12035121        100.2 ns/op         0 B/op        0 allocs/op
BenchmarkSlowMultiplication_65536-12            23061      53137 ns/op           0 B/op        0 allocs/op
BenchmarkSlowMultiplicationSIMD_65536-12       302212       4298 ns/op           0 B/op        0 allocs/op
PASS
ok   github.com/akramarenkov/goresearch/simd 25.170s
```
