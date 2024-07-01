# Research on the impact of defer on performance

Running example:

```bash
go test -bench=.* -benchmem ./...
```

Result example:

```bash
goos: linux
goarch: amd64
pkg: github.com/akramarenkov/goresearch/deferred
cpu: AMD Ryzen 5 3600 6-Core Processor              
BenchmarkReference-12              1000000000           0.2663 ns/op        0 B/op           0 allocs/op
BenchmarkNotDeferred-12            1000000000           0.7198 ns/op        0 B/op           0 allocs/op
BenchmarkDeferred-12               395775301            2.863 ns/op         0 B/op           0 allocs/op
BenchmarkReferenceIdle-12          1000000000           0.2687 ns/op        0 B/op           0 allocs/op
BenchmarkNotDeferredIdle-12        1000000000           0.2536 ns/op        0 B/op           0 allocs/op
BenchmarkDeferredIdle-12           443267794            2.421 ns/op         0 B/op           0 allocs/op
BenchmarkReference2-12             1000000000           0.2719 ns/op        0 B/op           0 allocs/op
BenchmarkNotDeferred2-12           227299069            5.172 ns/op         0 B/op           0 allocs/op
BenchmarkDeferred2-12              237940282            5.160 ns/op         0 B/op           0 allocs/op
BenchmarkReference4-12             1000000000           0.2579 ns/op        0 B/op           0 allocs/op
BenchmarkNotDeferred4-12           115775539           10.38 ns/op          0 B/op           0 allocs/op
BenchmarkDeferred4-12              103743432           11.35 ns/op          0 B/op           0 allocs/op
BenchmarkReference6-12             1000000000           0.2710 ns/op        0 B/op           0 allocs/op
BenchmarkNotDeferred6-12           81396103            15.16 ns/op          0 B/op           0 allocs/op
BenchmarkDeferred6-12              70203306            16.97 ns/op          0 B/op           0 allocs/op
BenchmarkReference8-12             1000000000           0.2719 ns/op        0 B/op           0 allocs/op
BenchmarkNotDeferred8-12           58132306            20.47 ns/op          0 B/op           0 allocs/op
BenchmarkDeferred8-12              54209341            22.31 ns/op          0 B/op           0 allocs/op
BenchmarkReference10-12            1000000000           0.2407 ns/op        0 B/op           0 allocs/op
BenchmarkNotDeferred10-12          48086668            25.23 ns/op          0 B/op           0 allocs/op
BenchmarkDeferred10-12             11167503           103.2 ns/op           0 B/op           0 allocs/op
BenchmarkReference11-12            1000000000           0.2403 ns/op        0 B/op           0 allocs/op
BenchmarkNotDeferred11-12          41147313            27.30 ns/op          0 B/op           0 allocs/op
BenchmarkDeferred11-12             10444593           120.2 ns/op           0 B/op           0 allocs/op
PASS
ok      github.com/akramarenkov/goresearch/deferred    27.290s

```
