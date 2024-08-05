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
BenchmarkOne-12                       1000000000             0.2531 ns/op           0 B/op           0 allocs/op
BenchmarkOneVariadic-12               1000000000             0.2630 ns/op           0 B/op           0 allocs/op
BenchmarkOneVariadicSlice-12          1000000000             0.2400 ns/op           0 B/op           0 allocs/op
BenchmarkTwo-12                       1000000000             0.2401 ns/op           0 B/op           0 allocs/op
BenchmarkTwoVariadic-12               1000000000             0.7191 ns/op           0 B/op           0 allocs/op
BenchmarkTwoVariadicSlice-12          1000000000             0.2538 ns/op           0 B/op           0 allocs/op
BenchmarkThree-12                     1000000000             0.2736 ns/op           0 B/op           0 allocs/op
BenchmarkThreeVariadic-12             1000000000             1.239 ns/op            0 B/op           0 allocs/op
BenchmarkThreeVariadicSlice-12        1000000000             0.2740 ns/op           0 B/op           0 allocs/op
BenchmarkFour-12                      1000000000             0.2446 ns/op           0 B/op           0 allocs/op
BenchmarkFourVariadic-12              718237180              1.473 ns/op            0 B/op           0 allocs/op
BenchmarkFourVariadicSlice-12         1000000000             0.5018 ns/op           0 B/op           0 allocs/op
BenchmarkSix-12                       1000000000             0.4931 ns/op           0 B/op           0 allocs/op
BenchmarkSixVariadic-12               483691861              2.261 ns/op            0 B/op           0 allocs/op
BenchmarkSixVariadicSlice-12          686432673              1.534 ns/op            0 B/op           0 allocs/op
BenchmarkEight-12                     1000000000             0.5029 ns/op           0 B/op           0 allocs/op
BenchmarkEightVariadic-12             374086616              3.011 ns/op            0 B/op           0 allocs/op
BenchmarkEightVariadicSlice-12        542135908              1.978 ns/op            0 B/op           0 allocs/op
BenchmarkTen-12                       1000000000             0.2731 ns/op           0 B/op           0 allocs/op
BenchmarkTenVariadic-12               265798716              4.097 ns/op            0 B/op           0 allocs/op
BenchmarkTenVariadicSlice-12          448342713              2.482 ns/op            0 B/op           0 allocs/op
PASS
ok      github.com/akramarenkov/goresearch/variadic    16.049s
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
BenchmarkOne-12                       930361135             1.553 ns/op           0 B/op           0 allocs/op
BenchmarkOneVariadic-12               487124655             2.234 ns/op           0 B/op           0 allocs/op
BenchmarkOneVariadicSlice-12          493168356             2.252 ns/op           0 B/op           0 allocs/op
BenchmarkTwo-12                       602174529             1.678 ns/op           0 B/op           0 allocs/op
BenchmarkTwoVariadic-12               421654069             2.398 ns/op           0 B/op           0 allocs/op
BenchmarkTwoVariadicSlice-12          423924050             2.739 ns/op           0 B/op           0 allocs/op
BenchmarkThree-12                     626398966             1.709 ns/op           0 B/op           0 allocs/op
BenchmarkThreeVariadic-12             342640392             3.221 ns/op           0 B/op           0 allocs/op
BenchmarkThreeVariadicSlice-12        389057337             3.088 ns/op           0 B/op           0 allocs/op
BenchmarkFour-12                      482025337             2.248 ns/op           0 B/op           0 allocs/op
BenchmarkFourVariadic-12              292298326             4.193 ns/op           0 B/op           0 allocs/op
BenchmarkFourVariadicSlice-12         312644694             3.897 ns/op           0 B/op           0 allocs/op
BenchmarkSix-12                       334237053             3.353 ns/op           0 B/op           0 allocs/op
BenchmarkSixVariadic-12               224491651             5.567 ns/op           0 B/op           0 allocs/op
BenchmarkSixVariadicSlice-12          219704061             5.265 ns/op           0 B/op           0 allocs/op
BenchmarkEight-12                     275642688             4.442 ns/op           0 B/op           0 allocs/op
BenchmarkEightVariadic-12             153079130             7.509 ns/op           0 B/op           0 allocs/op
BenchmarkEightVariadicSlice-12        176488580             6.615 ns/op           0 B/op           0 allocs/op
BenchmarkTen-12                       202830078             5.754 ns/op           0 B/op           0 allocs/op
BenchmarkTenVariadic-12               82405263             14.22  ns/op           0 B/op           0 allocs/op
BenchmarkTenVariadicSlice-12          147558308             7.813 ns/op           0 B/op           0 allocs/op
PASS
ok      github.com/akramarenkov/goresearch/variadic    32.961s
```
