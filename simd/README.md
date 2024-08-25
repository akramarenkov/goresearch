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
BenchmarkSlowMultiplication_64-12            35526039         33.94 ns/op        0 B/op        0 allocs/op
BenchmarkSlowMultiplicationSIMD_64-12        23476220         43.36 ns/op        0 B/op        0 allocs/op
BenchmarkSlowMultiplication_128-12           16290385         69.02 ns/op        0 B/op        0 allocs/op
BenchmarkSlowMultiplicationSIMD_128-12       27991550         42.65 ns/op        0 B/op        0 allocs/op
BenchmarkSlowMultiplication_256-12            9307885        130.3 ns/op         0 B/op        0 allocs/op
BenchmarkSlowMultiplicationSIMD_256-12       27408306         44.07 ns/op        0 B/op        0 allocs/op
BenchmarkSlowMultiplication_512-12            4720165        258.8 ns/op         0 B/op        0 allocs/op
BenchmarkSlowMultiplicationSIMD_512-12       26265427         45.18 ns/op        0 B/op        0 allocs/op
BenchmarkSlowMultiplication_1024-12           2272365        500.0 ns/op         0 B/op        0 allocs/op
BenchmarkSlowMultiplicationSIMD_1024-12      26367222         46.75 ns/op        0 B/op        0 allocs/op
BenchmarkSlowMultiplication_65536-12            34507      31467 ns/op           0 B/op        0 allocs/op
BenchmarkSlowMultiplicationSIMD_65536-12      3999472        284.6 ns/op         0 B/op        0 allocs/op
PASS
ok   github.com/akramarenkov/goresearch/simd 20.264s
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
BenchmarkSlowMultiplication_64-12            18008630         63.86 ns/op        0 B/op        0 allocs/op
BenchmarkSlowMultiplicationSIMD_64-12        22823816         46.80 ns/op        0 B/op        0 allocs/op
BenchmarkSlowMultiplication_128-12           10049812        110.1 ns/op         0 B/op        0 allocs/op
BenchmarkSlowMultiplicationSIMD_128-12       26123348         46.05 ns/op        0 B/op        0 allocs/op
BenchmarkSlowMultiplication_256-12            5574208        312.2 ns/op         0 B/op        0 allocs/op
BenchmarkSlowMultiplicationSIMD_256-12       25394832         46.56 ns/op        0 B/op        0 allocs/op
BenchmarkSlowMultiplication_512-12            2676610        650.4 ns/op         0 B/op        0 allocs/op
BenchmarkSlowMultiplicationSIMD_512-12       25492640         48.16 ns/op        0 B/op        0 allocs/op
BenchmarkSlowMultiplication_1024-12           1414107        820.4 ns/op         0 B/op        0 allocs/op
BenchmarkSlowMultiplicationSIMD_1024-12      24025528         48.51 ns/op        0 B/op        0 allocs/op
BenchmarkSlowMultiplication_65536-12            22348      52704 ns/op           0 B/op        0 allocs/op
BenchmarkSlowMultiplicationSIMD_65536-12      3787648        293.7 ns/op         0 B/op        0 allocs/op
PASS
ok   github.com/akramarenkov/goresearch/simd 23.381s
```
