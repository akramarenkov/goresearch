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
BenchmarkSlicesClone-12                29252478            61.12 ns/op          48 B/op           1 allocs/op
BenchmarkThroughAppend-12              19126020            58.59 ns/op          48 B/op           1 allocs/op
BenchmarkThroughCopy-12                24433706            44.34 ns/op          48 B/op           1 allocs/op
Benchmark_16B_SlicesClone-12           21101544            51.94 ns/op          16 B/op           1 allocs/op
Benchmark_16B_ThroughAppend-12         21869557            51.30 ns/op          16 B/op           1 allocs/op
Benchmark_16B_ThroughCopy-12           35118582            34.85 ns/op          16 B/op           1 allocs/op
Benchmark_32B_SlicesClone-12           19761015            61.07 ns/op          32 B/op           1 allocs/op
Benchmark_32B_ThroughAppend-12         20130310            57.88 ns/op          32 B/op           1 allocs/op
Benchmark_32B_ThroughCopy-12           29282737            41.29 ns/op          32 B/op           1 allocs/op
Benchmark_64B_SlicesClone-12           18775729            63.03 ns/op          64 B/op           1 allocs/op
Benchmark_64B_ThroughAppend-12         18048069            62.91 ns/op          64 B/op           1 allocs/op
Benchmark_64B_ThroughCopy-12           24403334            46.80 ns/op          64 B/op           1 allocs/op
Benchmark_128B_SlicesClone-12          16177312            73.36 ns/op         128 B/op           1 allocs/op
Benchmark_128B_ThroughAppend-12        15978940            74.41 ns/op         128 B/op           1 allocs/op
Benchmark_128B_ThroughCopy-12          21883706            56.76 ns/op         128 B/op           1 allocs/op
Benchmark_1KB_SlicesClone-12            4457072           270.9 ns/op         1024 B/op           1 allocs/op
Benchmark_1KB_ThroughAppend-12          4409181           269.6 ns/op         1024 B/op           1 allocs/op
Benchmark_1KB_ThroughCopy-12            4480916           262.0 ns/op         1024 B/op           1 allocs/op
Benchmark_1MB_SlicesClone-12               8138        145157 ns/op        1048576 B/op           1 allocs/op
Benchmark_1MB_ThroughAppend-12             9286        142085 ns/op        1048576 B/op           1 allocs/op
Benchmark_1MB_ThroughCopy-12               6028        190124 ns/op        1048577 B/op           1 allocs/op
PASS
ok      github.com/akramarenkov/goresearch/clone    29.038s
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
BenchmarkSlicesClone-12                20975946            73.26 ns/op          48 B/op           1 allocs/op
BenchmarkThroughAppend-12              17024802            71.51 ns/op          48 B/op           1 allocs/op
BenchmarkThroughCopy-12                18466411            62.52 ns/op          48 B/op           1 allocs/op
Benchmark_16B_SlicesClone-12           17908898            65.41 ns/op          16 B/op           1 allocs/op
Benchmark_16B_ThroughAppend-12         26786160            66.78 ns/op          16 B/op           1 allocs/op
Benchmark_16B_ThroughCopy-12           20597984            55.28 ns/op          16 B/op           1 allocs/op
Benchmark_32B_SlicesClone-12           18039428            67.03 ns/op          32 B/op           1 allocs/op
Benchmark_32B_ThroughAppend-12         16721265            68.01 ns/op          32 B/op           1 allocs/op
Benchmark_32B_ThroughCopy-12           19788736            58.54 ns/op          32 B/op           1 allocs/op
Benchmark_64B_SlicesClone-12           16824711            72.28 ns/op          64 B/op           1 allocs/op
Benchmark_64B_ThroughAppend-12         15792369            71.64 ns/op          64 B/op           1 allocs/op
Benchmark_64B_ThroughCopy-12           18750921            63.65 ns/op          64 B/op           1 allocs/op
Benchmark_128B_SlicesClone-12          13839547            82.98 ns/op         128 B/op           1 allocs/op
Benchmark_128B_ThroughAppend-12        14171996            83.38 ns/op         128 B/op           1 allocs/op
Benchmark_128B_ThroughCopy-12          15138931            77.00 ns/op         128 B/op           1 allocs/op
Benchmark_1KB_SlicesClone-12            4341067           284.0 ns/op         1024 B/op           1 allocs/op
Benchmark_1KB_ThroughAppend-12          4228815           281.6 ns/op         1024 B/op           1 allocs/op
Benchmark_1KB_ThroughCopy-12            4312430           280.7 ns/op         1024 B/op           1 allocs/op
Benchmark_1MB_SlicesClone-12               8715        133197 ns/op        1048576 B/op           1 allocs/op
Benchmark_1MB_ThroughAppend-12             8679        132752 ns/op        1048576 B/op           1 allocs/op
Benchmark_1MB_ThroughCopy-12               5630        233025 ns/op        1048577 B/op           1 allocs/op
PASS
ok      github.com/akramarenkov/goresearch/clone    29.616s
```
