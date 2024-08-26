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
BenchmarkSlowMultiplication_64-12               70544082         17.00 ns/op        0 B/op        0 allocs/op
BenchmarkSlowMultiplicationSIMD_64-12           27676060         46.08 ns/op        0 B/op        0 allocs/op
BenchmarkSlowMultiplication_128-12              30903054         37.16 ns/op        0 B/op        0 allocs/op
BenchmarkSlowMultiplicationSIMD_128-12          22795730         45.10 ns/op        0 B/op        0 allocs/op
BenchmarkSlowMultiplication_256-12              17937307         68.92 ns/op        0 B/op        0 allocs/op
BenchmarkSlowMultiplicationSIMD_256-12          27377499         45.56 ns/op        0 B/op        0 allocs/op
BenchmarkSlowMultiplication_512-12               8889192        133.6 ns/op         0 B/op        0 allocs/op
BenchmarkSlowMultiplicationSIMD_512-12          25445817         47.60 ns/op        0 B/op        0 allocs/op
BenchmarkSlowMultiplication_1024-12              4661424        257.6 ns/op         0 B/op        0 allocs/op
BenchmarkSlowMultiplicationSIMD_1024-12         23792515         48.71 ns/op        0 B/op        0 allocs/op
BenchmarkSlowMultiplication_65536-12               73405      16165 ns/op           0 B/op        0 allocs/op
BenchmarkSlowMultiplicationSIMD_65536-12         3783394        294.9 ns/op         0 B/op        0 allocs/op
BenchmarkSlowMultiplicationF64_64-12            38461934         30.84 ns/op        0 B/op        0 allocs/op
BenchmarkSlowMultiplicationF64SIMD_64-12        25011186         49.49 ns/op        0 B/op        0 allocs/op
BenchmarkSlowMultiplicationF64_128-12           14580063         81.12 ns/op        0 B/op        0 allocs/op
BenchmarkSlowMultiplicationF64SIMD_128-12       20383242         58.09 ns/op        0 B/op        0 allocs/op
BenchmarkSlowMultiplicationF64_256-12            6802171        175.4 ns/op         0 B/op        0 allocs/op
BenchmarkSlowMultiplicationF64SIMD_256-12       14547673         79.97 ns/op        0 B/op        0 allocs/op
BenchmarkSlowMultiplicationF64_512-12            3089491        365.7 ns/op         0 B/op        0 allocs/op
BenchmarkSlowMultiplicationF64SIMD_512-12        9548721        128.5 ns/op         0 B/op        0 allocs/op
BenchmarkSlowMultiplicationF64_1024-12           1582714        725.7 ns/op         0 B/op        0 allocs/op
BenchmarkSlowMultiplicationF64SIMD_1024-12       4863008        217.2 ns/op         0 B/op        0 allocs/op
BenchmarkSlowMultiplicationF64_65536-12            25393      48470 ns/op           0 B/op        0 allocs/op
BenchmarkSlowMultiplicationF64SIMD_65536-12        97886      11835 ns/op           0 B/op        0 allocs/op
PASS
ok   github.com/akramarenkov/goresearch/simd 42.754s
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
BenchmarkSlowMultiplication_64-12               17978544         66.96 ns/op        0 B/op        0 allocs/op
BenchmarkSlowMultiplicationSIMD_64-12           24070126         48.17 ns/op        0 B/op        0 allocs/op
BenchmarkSlowMultiplication_128-12              10117098        110.3 ns/op         0 B/op        0 allocs/op
BenchmarkSlowMultiplicationSIMD_128-12          24024332         49.04 ns/op        0 B/op        0 allocs/op
BenchmarkSlowMultiplication_256-12               5424510        293.2 ns/op         0 B/op        0 allocs/op
BenchmarkSlowMultiplicationSIMD_256-12          24320710         49.58 ns/op        0 B/op        0 allocs/op
BenchmarkSlowMultiplication_512-12               2707010        422.1 ns/op         0 B/op        0 allocs/op
BenchmarkSlowMultiplicationSIMD_512-12          23624630         50.75 ns/op        0 B/op        0 allocs/op
BenchmarkSlowMultiplication_1024-12              1410066        834.5 ns/op         0 B/op        0 allocs/op
BenchmarkSlowMultiplicationSIMD_1024-12         23529157         52.76 ns/op        0 B/op        0 allocs/op
BenchmarkSlowMultiplication_65536-12               22845      53431 ns/op           0 B/op        0 allocs/op
BenchmarkSlowMultiplicationSIMD_65536-12         4112481        297.2 ns/op         0 B/op        0 allocs/op
BenchmarkSlowMultiplicationF64_64-12             7661340        156.5 ns/op         0 B/op        0 allocs/op
BenchmarkSlowMultiplicationF64SIMD_64-12        23257471         53.33 ns/op        0 B/op        0 allocs/op
BenchmarkSlowMultiplicationF64_128-12            3696250        332.9 ns/op         0 B/op        0 allocs/op
BenchmarkSlowMultiplicationF64SIMD_128-12       19548321         60.38 ns/op        0 B/op        0 allocs/op
BenchmarkSlowMultiplicationF64_256-12            1810239        662.3 ns/op         0 B/op        0 allocs/op
BenchmarkSlowMultiplicationF64SIMD_256-12       14408590         81.98 ns/op        0 B/op        0 allocs/op
BenchmarkSlowMultiplicationF64_512-12             757501       1338 ns/op           0 B/op        0 allocs/op
BenchmarkSlowMultiplicationF64SIMD_512-12        9367772        130.8 ns/op         0 B/op        0 allocs/op
BenchmarkSlowMultiplicationF64_1024-12            444704       2900 ns/op           0 B/op        0 allocs/op
BenchmarkSlowMultiplicationF64SIMD_1024-12       5144238        229.7 ns/op         0 B/op        0 allocs/op
BenchmarkSlowMultiplicationF64_65536-12             6212     177668 ns/op           0 B/op        0 allocs/op
BenchmarkSlowMultiplicationF64SIMD_65536-12        99056      11867 ns/op           0 B/op        0 allocs/op
PASS
ok   github.com/akramarenkov/goresearch/simd 45.398s
```
