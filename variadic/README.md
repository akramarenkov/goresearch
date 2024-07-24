# Research on the impact of using variadic function on performance

Running example:

```bash
go test -bench=.* -benchmem ./...
```

Result example:

```bash
goos: linux
goarch: amd64
pkg: github.com/akramarenkov/goresearch/variadic
cpu: AMD Ryzen 5 3600 6-Core Processor              
BenchmarkOne-12                  1000000000            0.2639 ns/op          0 B/op           0 allocs/op
BenchmarkOneVariadic-12          1000000000            0.2703 ns/op          0 B/op           0 allocs/op
BenchmarkTwo-12                  1000000000            0.2446 ns/op          0 B/op           0 allocs/op
BenchmarkTwoVariadic-12          1000000000            0.7611 ns/op          0 B/op           0 allocs/op
BenchmarkThree-12                1000000000            0.2762 ns/op          0 B/op           0 allocs/op
BenchmarkThreeVariadic-12        968327530             1.246 ns/op           0 B/op           0 allocs/op
BenchmarkFour-12                 1000000000            0.2573 ns/op          0 B/op           0 allocs/op
BenchmarkFourVariadic-12         723221860             1.485 ns/op           0 B/op           0 allocs/op
BenchmarkSix-12                  1000000000            0.5207 ns/op          0 B/op           0 allocs/op
BenchmarkSixVariadic-12          526631100             2.248 ns/op           0 B/op           0 allocs/op
BenchmarkEight-12                1000000000            0.5312 ns/op          0 B/op           0 allocs/op
BenchmarkEightVariadic-12        371642413             2.977 ns/op           0 B/op           0 allocs/op
BenchmarkTen-12                  1000000000            0.4988 ns/op          0 B/op           0 allocs/op
BenchmarkTenVariadic-12          277757012             4.137 ns/op           0 B/op           0 allocs/op
PASS
ok      github.com/akramarenkov/goresearch/variadic    11.140s
```

Running example without optimizations:

```bash
go test -bench=.* -benchmem -gcflags=-N ./...
```

Result example without optimizations:

```bash
goos: linux
goarch: amd64
pkg: github.com/akramarenkov/goresearch/variadic
cpu: AMD Ryzen 5 3600 6-Core Processor              
BenchmarkOne-12                  994632672            1.231 ns/op          0 B/op           0 allocs/op
BenchmarkOneVariadic-12          480926498            2.248 ns/op          0 B/op           0 allocs/op
BenchmarkTwo-12                  596505506            1.749 ns/op          0 B/op           0 allocs/op
BenchmarkTwoVariadic-12          467258839            2.523 ns/op          0 B/op           0 allocs/op
BenchmarkThree-12                591193303            1.858 ns/op          0 B/op           0 allocs/op
BenchmarkThreeVariadic-12        344229830            3.330 ns/op          0 B/op           0 allocs/op
BenchmarkFour-12                 546973488            2.239 ns/op          0 B/op           0 allocs/op
BenchmarkFourVariadic-12         272517946            4.177 ns/op          0 B/op           0 allocs/op
BenchmarkSix-12                  336804388            3.351 ns/op          0 B/op           0 allocs/op
BenchmarkSixVariadic-12          201673435            5.795 ns/op          0 B/op           0 allocs/op
BenchmarkEight-12                253323657            4.492 ns/op          0 B/op           0 allocs/op
BenchmarkEightVariadic-12        153463149            8.031 ns/op          0 B/op           0 allocs/op
BenchmarkTen-12                  210968979            5.750 ns/op          0 B/op           0 allocs/op
BenchmarkTenVariadic-12          84912114            14.15 ns/op           0 B/op           0 allocs/op
PASS
ok      github.com/akramarenkov/goresearch/variadic    21.155s
```
