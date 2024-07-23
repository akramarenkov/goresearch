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
BenchmarkReference-12              1000000000           0.2400 ns/op        0 B/op           0 allocs/op
BenchmarkNotDeferred-12            1000000000           0.2734 ns/op        0 B/op           0 allocs/op
BenchmarkDeferred-12               459613740            2.651 ns/op         0 B/op           0 allocs/op
BenchmarkReferenceIdle-12          1000000000           0.2397 ns/op        0 B/op           0 allocs/op
BenchmarkNotDeferredIdle-12        1000000000           0.2521 ns/op        0 B/op           0 allocs/op
BenchmarkDeferredIdle-12           499279069            2.397 ns/op         0 B/op           0 allocs/op
BenchmarkReference2-12             1000000000           0.5130 ns/op        0 B/op           0 allocs/op
BenchmarkNotDeferred2-12           318114310            3.580 ns/op         0 B/op           0 allocs/op
BenchmarkDeferred2-12              290891766            4.199 ns/op         0 B/op           0 allocs/op
BenchmarkReference4-12             1000000000           0.5121 ns/op        0 B/op           0 allocs/op
BenchmarkNotDeferred4-12           173232747            6.958 ns/op         0 B/op           0 allocs/op
BenchmarkDeferred4-12              103334169           11.35 ns/op          0 B/op           0 allocs/op
BenchmarkReference6-12             1000000000           0.2401 ns/op        0 B/op           0 allocs/op
BenchmarkNotDeferred6-12           108287259           10.84 ns/op          0 B/op           0 allocs/op
BenchmarkDeferred6-12              75953253            15.10 ns/op          0 B/op           0 allocs/op
BenchmarkReference8-12             1000000000           0.2709 ns/op        0 B/op           0 allocs/op
BenchmarkNotDeferred8-12           83960371            14.51 ns/op          0 B/op           0 allocs/op
BenchmarkDeferred8-12              54324934            22.58 ns/op          0 B/op           0 allocs/op
BenchmarkReference10-12            1000000000           0.4799 ns/op        0 B/op           0 allocs/op
BenchmarkNotDeferred10-12          69950504            17.50 ns/op          0 B/op           0 allocs/op
BenchmarkDeferred10-12             11628180           103.1 ns/op           0 B/op           0 allocs/op
BenchmarkReference11-12            1000000000           0.4795 ns/op        0 B/op           0 allocs/op
BenchmarkNotDeferred11-12          60796129            18.72 ns/op          0 B/op           0 allocs/op
BenchmarkDeferred11-12             10498614           114.0 ns/op           0 B/op           0 allocs/op
PASS
ok      github.com/akramarenkov/goresearch/deferred    28.192s
```

Running example without optimizations:

```bash
go test -bench=.* -benchmem -gcflags=-N ./...
```

Result example without optimizations:

```bash
goos: linux
goarch: amd64
pkg: github.com/akramarenkov/goresearch/deferred
cpu: AMD Ryzen 5 3600 6-Core Processor              
BenchmarkReference-12              1000000000          1.198 ns/op         0 B/op           0 allocs/op
BenchmarkNotDeferred-12            613934799           1.732 ns/op         0 B/op           0 allocs/op
BenchmarkDeferred-12               50022240           24.62 ns/op          0 B/op           0 allocs/op
BenchmarkReferenceIdle-12          1000000000          0.9864 ns/op        0 B/op           0 allocs/op
BenchmarkNotDeferredIdle-12        665928440           2.172 ns/op         0 B/op           0 allocs/op
BenchmarkDeferredIdle-12           43373122           25.24 ns/op          0 B/op           0 allocs/op
BenchmarkReference2-12             587326773           1.711 ns/op         0 B/op           0 allocs/op
BenchmarkNotDeferred2-12           230482862           5.009 ns/op         0 B/op           0 allocs/op
BenchmarkDeferred2-12              32276529           37.65 ns/op          0 B/op           0 allocs/op
BenchmarkReference4-12             353146765           3.209 ns/op         0 B/op           0 allocs/op
BenchmarkNotDeferred4-12           131471248           9.037 ns/op         0 B/op           0 allocs/op
BenchmarkDeferred4-12              18569334           65.00 ns/op          0 B/op           0 allocs/op
BenchmarkReference6-12             245129386           4.555 ns/op         0 B/op           0 allocs/op
BenchmarkNotDeferred6-12           85426002           13.32 ns/op          0 B/op           0 allocs/op
BenchmarkDeferred6-12              14991230           80.17 ns/op          0 B/op           0 allocs/op
BenchmarkReference8-12             191491674           6.167 ns/op         0 B/op           0 allocs/op
BenchmarkNotDeferred8-12           66419971           17.50 ns/op          0 B/op           0 allocs/op
BenchmarkDeferred8-12              11993130          101.9 ns/op           0 B/op           0 allocs/op
BenchmarkReference10-12            161488378           7.636 ns/op         0 B/op           0 allocs/op
BenchmarkNotDeferred10-12          54891165           21.80 ns/op          0 B/op           0 allocs/op
BenchmarkDeferred10-12             8454690           120.0 ns/op           0 B/op           0 allocs/op
BenchmarkReference11-12            146565667           8.146 ns/op         0 B/op           0 allocs/op
BenchmarkNotDeferred11-12          48604437           24.71 ns/op          0 B/op           0 allocs/op
BenchmarkDeferred11-12             9393800           128.1 ns/op           0 B/op           0 allocs/op
PASS
ok      github.com/akramarenkov/goresearch/deferred    37.981s
```
