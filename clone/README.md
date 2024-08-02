# Copy performance comparison

Running example:

```bash
go test -bench=.* -benchmem ./...
```

Result example:

```bash
goos: linux
goarch: amd64
pkg: github.com/akramarenkov/goresearch/clone
cpu: AMD Ryzen 5 3600 6-Core Processor              
BenchmarkSlicesClone-12                22397168            61.47 ns/op          48 B/op           1 allocs/op
BenchmarkThroughAppend-12              17758291            65.69 ns/op          48 B/op           1 allocs/op
BenchmarkThroughCopy-12                24323726            45.75 ns/op          48 B/op           1 allocs/op
Benchmark_16B_SlicesClone-12           21956949            51.24 ns/op          16 B/op           1 allocs/op
Benchmark_16B_ThroughAppend-12         22256680            50.58 ns/op          16 B/op           1 allocs/op
Benchmark_16B_ThroughCopy-12           32436123            32.75 ns/op          16 B/op           1 allocs/op
Benchmark_32B_SlicesClone-12           18928647            56.53 ns/op          32 B/op           1 allocs/op
Benchmark_32B_ThroughAppend-12         19661013            56.55 ns/op          32 B/op           1 allocs/op
Benchmark_32B_ThroughCopy-12           26779504            40.92 ns/op          32 B/op           1 allocs/op
Benchmark_64B_SlicesClone-12           17499022            66.10 ns/op          64 B/op           1 allocs/op
Benchmark_64B_ThroughAppend-12         18140481            62.78 ns/op          64 B/op           1 allocs/op
Benchmark_64B_ThroughCopy-12           22843294            46.59 ns/op          64 B/op           1 allocs/op
Benchmark_128B_SlicesClone-12          16527813            72.29 ns/op         128 B/op           1 allocs/op
Benchmark_128B_ThroughAppend-12        15715479            73.18 ns/op         128 B/op           1 allocs/op
Benchmark_128B_ThroughCopy-12          20151787            56.69 ns/op         128 B/op           1 allocs/op
Benchmark_1KB_SlicesClone-12            4529434           266.3 ns/op         1024 B/op           1 allocs/op
Benchmark_1KB_ThroughAppend-12          4387381           268.1 ns/op         1024 B/op           1 allocs/op
Benchmark_1KB_ThroughCopy-12            4668588           259.6 ns/op         1024 B/op           1 allocs/op
Benchmark_1MB_SlicesClone-12               9210        137293 ns/op        1048576 B/op           1 allocs/op
Benchmark_1MB_ThroughAppend-12             9500        138519 ns/op        1048576 B/op           1 allocs/op
Benchmark_1MB_ThroughCopy-12               7275        161109 ns/op        1048576 B/op           1 allocs/op
PASS
ok      github.com/akramarenkov/goresearch/clone    29.049s

```

Running example without optimizations:

```bash
go test -bench=.* -benchmem -gcflags=-N ./...
```

Result example without optimizations:

```bash
goos: linux
goarch: amd64
pkg: github.com/akramarenkov/goresearch/clone
cpu: AMD Ryzen 5 3600 6-Core Processor              
BenchmarkSlicesClone-12                19658838            73.49 ns/op          48 B/op           1 allocs/op
BenchmarkThroughAppend-12              18009960            69.18 ns/op          48 B/op           1 allocs/op
BenchmarkThroughCopy-12                19233308            60.21 ns/op          48 B/op           1 allocs/op
Benchmark_16B_SlicesClone-12           18136663            61.44 ns/op          16 B/op           1 allocs/op
Benchmark_16B_ThroughAppend-12         29872119            60.84 ns/op          16 B/op           1 allocs/op
Benchmark_16B_ThroughCopy-12           20516802            54.97 ns/op          16 B/op           1 allocs/op
Benchmark_32B_SlicesClone-12           17225055            66.36 ns/op          32 B/op           1 allocs/op
Benchmark_32B_ThroughAppend-12         18063064            63.77 ns/op          32 B/op           1 allocs/op
Benchmark_32B_ThroughCopy-12           21197215            57.13 ns/op          32 B/op           1 allocs/op
Benchmark_64B_SlicesClone-12           16715428            71.96 ns/op          64 B/op           1 allocs/op
Benchmark_64B_ThroughAppend-12         16289227            71.03 ns/op          64 B/op           1 allocs/op
Benchmark_64B_ThroughCopy-12           19149770            62.56 ns/op          64 B/op           1 allocs/op
Benchmark_128B_SlicesClone-12          13770097            85.92 ns/op         128 B/op           1 allocs/op
Benchmark_128B_ThroughAppend-12        14379310            80.69 ns/op         128 B/op           1 allocs/op
Benchmark_128B_ThroughCopy-12          15018981            75.44 ns/op         128 B/op           1 allocs/op
Benchmark_1KB_SlicesClone-12            4441278           275.6 ns/op         1024 B/op           1 allocs/op
Benchmark_1KB_ThroughAppend-12          4262001           279.7 ns/op         1024 B/op           1 allocs/op
Benchmark_1KB_ThroughCopy-12            4333434           278.4 ns/op         1024 B/op           1 allocs/op
Benchmark_1MB_SlicesClone-12               6598        153487 ns/op        1048577 B/op           1 allocs/op
Benchmark_1MB_ThroughAppend-12             7542        149300 ns/op        1048577 B/op           1 allocs/op
Benchmark_1MB_ThroughCopy-12               5481        185212 ns/op        1048577 B/op           1 allocs/op
PASS
ok      github.com/akramarenkov/goresearch/clone    27.909s

```
