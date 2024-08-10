# Slice cloning performance comparison

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
BenchmarkSlicesClone-12                24062085            62.38 ns/op          48 B/op           1 allocs/op
BenchmarkThroughAppend-12              17961103            61.09 ns/op          48 B/op           1 allocs/op
BenchmarkThroughCopy-12                25089546            44.47 ns/op          48 B/op           1 allocs/op
Benchmark_16B_SlicesClone-12           19176639            55.73 ns/op          16 B/op           1 allocs/op
Benchmark_16B_ThroughAppend-12         19494410            54.64 ns/op          16 B/op           1 allocs/op
Benchmark_16B_ThroughCopy-12           27642034            36.50 ns/op          16 B/op           1 allocs/op
Benchmark_32B_SlicesClone-12           21036723            57.29 ns/op          32 B/op           1 allocs/op
Benchmark_32B_ThroughAppend-12         19364068            56.78 ns/op          32 B/op           1 allocs/op
Benchmark_32B_ThroughCopy-12           25485942            39.77 ns/op          32 B/op           1 allocs/op
Benchmark_64B_SlicesClone-12           18757338            62.37 ns/op          64 B/op           1 allocs/op
Benchmark_64B_ThroughAppend-12         18438795            64.93 ns/op          64 B/op           1 allocs/op
Benchmark_64B_ThroughCopy-12           22870263            45.25 ns/op          64 B/op           1 allocs/op
Benchmark_128B_SlicesClone-12          16841658            72.79 ns/op         128 B/op           1 allocs/op
Benchmark_128B_ThroughAppend-12        16675440            75.38 ns/op         128 B/op           1 allocs/op
Benchmark_128B_ThroughCopy-12          20407275            55.31 ns/op         128 B/op           1 allocs/op
Benchmark_1KB_SlicesClone-12            4504489           269.0 ns/op         1024 B/op           1 allocs/op
Benchmark_1KB_ThroughAppend-12          4342705           269.1 ns/op         1024 B/op           1 allocs/op
Benchmark_1KB_ThroughCopy-12            4547306           259.4 ns/op         1024 B/op           1 allocs/op
Benchmark_1MB_SlicesClone-12               8775        137473 ns/op        1048577 B/op           1 allocs/op
Benchmark_1MB_ThroughAppend-12             7191        144163 ns/op        1048577 B/op           1 allocs/op
Benchmark_1MB_ThroughCopy-12               7610        146825 ns/op        1048577 B/op           1 allocs/op
PASS
ok      github.com/akramarenkov/goresearch/clone    26.892s
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
BenchmarkSlicesClone-12                17740669            73.90 ns/op          48 B/op           1 allocs/op
BenchmarkThroughAppend-12              16436650            69.62 ns/op          48 B/op           1 allocs/op
BenchmarkThroughCopy-12                18995998            62.39 ns/op          48 B/op           1 allocs/op
Benchmark_16B_SlicesClone-12           17306532            65.47 ns/op          16 B/op           1 allocs/op
Benchmark_16B_ThroughAppend-12         17332122            65.13 ns/op          16 B/op           1 allocs/op
Benchmark_16B_ThroughCopy-12           20066188            57.41 ns/op          16 B/op           1 allocs/op
Benchmark_32B_SlicesClone-12           17174340            66.10 ns/op          32 B/op           1 allocs/op
Benchmark_32B_ThroughAppend-12         17532056            65.19 ns/op          32 B/op           1 allocs/op
Benchmark_32B_ThroughCopy-12           23416042            58.18 ns/op          32 B/op           1 allocs/op
Benchmark_64B_SlicesClone-12           16527429            71.60 ns/op          64 B/op           1 allocs/op
Benchmark_64B_ThroughAppend-12         15868380            71.28 ns/op          64 B/op           1 allocs/op
Benchmark_64B_ThroughCopy-12           18265540            62.28 ns/op          64 B/op           1 allocs/op
Benchmark_128B_SlicesClone-12          14197839            81.46 ns/op         128 B/op           1 allocs/op
Benchmark_128B_ThroughAppend-12        14795859            80.17 ns/op         128 B/op           1 allocs/op
Benchmark_128B_ThroughCopy-12          16149037            72.66 ns/op         128 B/op           1 allocs/op
Benchmark_1KB_SlicesClone-12            4514607           271.0 ns/op         1024 B/op           1 allocs/op
Benchmark_1KB_ThroughAppend-12          4411366           275.1 ns/op         1024 B/op           1 allocs/op
Benchmark_1KB_ThroughCopy-12            4357315           276.8 ns/op         1024 B/op           1 allocs/op
Benchmark_1MB_SlicesClone-12               6889        153539 ns/op        1048576 B/op           1 allocs/op
Benchmark_1MB_ThroughAppend-12             8148        133012 ns/op        1048577 B/op           1 allocs/op
Benchmark_1MB_ThroughCopy-12               5588        186729 ns/op        1048576 B/op           1 allocs/op
PASS
ok      github.com/akramarenkov/goresearch/clone    27.377s
```
