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
BenchmarkOne-12                       1000000000             0.2480 ns/op           0 B/op           0 allocs/op
BenchmarkOneVariadic-12               1000000000             0.2768 ns/op           0 B/op           0 allocs/op
BenchmarkOneVariadicSlice-12          1000000000             0.2419 ns/op           0 B/op           0 allocs/op
BenchmarkTwo-12                       1000000000             0.2578 ns/op           0 B/op           0 allocs/op
BenchmarkTwoVariadic-12               1000000000             0.7372 ns/op           0 B/op           0 allocs/op
BenchmarkTwoVariadicSlice-12          1000000000             0.2444 ns/op           0 B/op           0 allocs/op
BenchmarkThree-12                     1000000000             0.4948 ns/op           0 B/op           0 allocs/op
BenchmarkThreeVariadic-12             994219862              1.219 ns/op            0 B/op           0 allocs/op
BenchmarkThreeVariadicSlice-12        1000000000             0.5180 ns/op           0 B/op           0 allocs/op
BenchmarkFour-12                      1000000000             0.2855 ns/op           0 B/op           0 allocs/op
BenchmarkFourVariadic-12              738554181              1.550 ns/op            0 B/op           0 allocs/op
BenchmarkFourVariadicSlice-12         1000000000             0.5160 ns/op           0 B/op           0 allocs/op
BenchmarkSix-12                       1000000000             0.2808 ns/op           0 B/op           0 allocs/op
BenchmarkSixVariadic-12               466138543              2.317 ns/op            0 B/op           0 allocs/op
BenchmarkSixVariadicSlice-12          794829153              1.522 ns/op            0 B/op           0 allocs/op
BenchmarkEight-12                     1000000000             0.2675 ns/op           0 B/op           0 allocs/op
BenchmarkEightVariadic-12             397253168              2.947 ns/op            0 B/op           0 allocs/op
BenchmarkEightVariadicSlice-12        519103387              2.016 ns/op            0 B/op           0 allocs/op
BenchmarkTen-12                       1000000000             0.2537 ns/op           0 B/op           0 allocs/op
BenchmarkTenVariadic-12               292768915              4.256 ns/op            0 B/op           0 allocs/op
BenchmarkTenVariadicSlice-12          471175318              2.519 ns/op            0 B/op           0 allocs/op
PASS
ok      github.com/akramarenkov/goresearch/variadic    16.464s
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
BenchmarkOne-12                       997637875             1.409 ns/op           0 B/op           0 allocs/op
BenchmarkOneVariadic-12               541457486             2.196 ns/op           0 B/op           0 allocs/op
BenchmarkOneVariadicSlice-12          489893100             2.183 ns/op           0 B/op           0 allocs/op
BenchmarkTwo-12                       651310392             1.773 ns/op           0 B/op           0 allocs/op
BenchmarkTwoVariadic-12               479584719             2.609 ns/op           0 B/op           0 allocs/op
BenchmarkTwoVariadicSlice-12          348570208             3.303 ns/op           0 B/op           0 allocs/op
BenchmarkThree-12                     552116302             1.870 ns/op           0 B/op           0 allocs/op
BenchmarkThreeVariadic-12             371086665             3.280 ns/op           0 B/op           0 allocs/op
BenchmarkThreeVariadicSlice-12        303734491             4.031 ns/op           0 B/op           0 allocs/op
BenchmarkFour-12                      500484670             2.214 ns/op           0 B/op           0 allocs/op
BenchmarkFourVariadic-12              280298151             4.199 ns/op           0 B/op           0 allocs/op
BenchmarkFourVariadicSlice-12         268748270             4.677 ns/op           0 B/op           0 allocs/op
BenchmarkSix-12                       364905781             3.349 ns/op           0 B/op           0 allocs/op
BenchmarkSixVariadic-12               212314136             5.896 ns/op           0 B/op           0 allocs/op
BenchmarkSixVariadicSlice-12          187320420             6.552 ns/op           0 B/op           0 allocs/op
BenchmarkEight-12                     265325083             4.480 ns/op           0 B/op           0 allocs/op
BenchmarkEightVariadic-12             134082571             9.067 ns/op           0 B/op           0 allocs/op
BenchmarkEightVariadicSlice-12        130525062             9.036 ns/op           0 B/op           0 allocs/op
BenchmarkTen-12                       191407500             5.745 ns/op           0 B/op           0 allocs/op
BenchmarkTenVariadic-12               100000000            10.04 ns/op            0 B/op           0 allocs/op
BenchmarkTenVariadicSlice-12          89382726             13.15 ns/op            0 B/op           0 allocs/op
PASS
ok      github.com/akramarenkov/goresearch/variadic    32.909s
```
