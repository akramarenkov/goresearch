# Research on the impact of defer on performance

Running example:

```bash
go test -bench=.* -benchmem ./...
```

Result example:

```bash
goos: linux
goarch: amd64
pkg: github.com/akramarenkov/goresearch/defered
cpu: AMD Ryzen 5 3600 6-Core Processor              
BenchmarkReference-12             1000000000         0.2620 ns/op          0 B/op           0 allocs/op
BenchmarkNotDefered-12            1000000000         0.7521 ns/op          0 B/op           0 allocs/op
BenchmarkDefered-12               397150920          2.802 ns/op           0 B/op           0 allocs/op
BenchmarkReferenceIdle-12         1000000000         0.2540 ns/op          0 B/op           0 allocs/op
BenchmarkNotDeferedIdle-12        1000000000         0.2718 ns/op          0 B/op           0 allocs/op
BenchmarkDeferedIdle-12           446325628          2.462 ns/op           0 B/op           0 allocs/op
BenchmarkReference2-12            1000000000         0.2702 ns/op          0 B/op           0 allocs/op
BenchmarkNotDefered2-12           222535483          5.052 ns/op           0 B/op           0 allocs/op
BenchmarkDefered2-12              223446622          5.033 ns/op           0 B/op           0 allocs/op
BenchmarkReference4-12            1000000000         0.2405 ns/op          0 B/op           0 allocs/op
BenchmarkNotDefered4-12           118954453          10.37 ns/op           0 B/op           0 allocs/op
BenchmarkDefered4-12              100000000          11.11 ns/op           0 B/op           0 allocs/op
BenchmarkReference6-12            1000000000         0.2593 ns/op          0 B/op           0 allocs/op
BenchmarkNotDefered6-12           79372106           15.14 ns/op           0 B/op           0 allocs/op
BenchmarkDefered6-12              72624326           16.99 ns/op           0 B/op           0 allocs/op
BenchmarkReference8-12            1000000000         0.2649 ns/op          0 B/op           0 allocs/op
BenchmarkNotDefered8-12           59370925           20.27 ns/op           0 B/op           0 allocs/op
BenchmarkDefered8-12              51873313           22.59 ns/op           0 B/op           0 allocs/op
BenchmarkReference10-12           1000000000         0.2675 ns/op          0 B/op           0 allocs/op
BenchmarkNotDefered10-12          48052515           25.57 ns/op           0 B/op           0 allocs/op
BenchmarkDefered10-12             11604866           105.6 ns/op           0 B/op           0 allocs/op
BenchmarkReference11-12           1000000000         0.2578 ns/op          0 B/op           0 allocs/op
BenchmarkNotDefered11-12          43687959           28.02 ns/op           0 B/op           0 allocs/op
BenchmarkDefered11-12             10537629           116.7 ns/op           0 B/op           0 allocs/op
PASS
ok      github.com/akramarenkov/goresearch/defered    26.866s
```
