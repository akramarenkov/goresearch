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
BenchmarkOne-12                  1000000000            0.2397 ns/op          0 B/op           0 allocs/op
BenchmarkOneVariadic-12          1000000000            0.2627 ns/op          0 B/op           0 allocs/op
BenchmarkTwo-12                  1000000000            0.2672 ns/op          0 B/op           0 allocs/op
BenchmarkTwoVariadic-12          1000000000            0.7583 ns/op          0 B/op           0 allocs/op
BenchmarkThree-12                1000000000            0.2715 ns/op          0 B/op           0 allocs/op
BenchmarkThreeVariadic-12        940598256             1.199 ns/op           0 B/op           0 allocs/op
BenchmarkFour-12                 1000000000            0.2400 ns/op          0 B/op           0 allocs/op
BenchmarkFourVariadic-12         703097440             1.485 ns/op           0 B/op           0 allocs/op
PASS
ok      github.com/akramarenkov/goresearch/variadic    4.838s
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
BenchmarkOne-12                  997701146             1.202 ns/op           0 B/op           0 allocs/op
BenchmarkOneVariadic-12          486445124             2.174 ns/op           0 B/op           0 allocs/op
BenchmarkTwo-12                  694978266             1.677 ns/op           0 B/op           0 allocs/op
BenchmarkTwoVariadic-12          437557695             2.445 ns/op           0 B/op           0 allocs/op
BenchmarkThree-12                619435599             1.692 ns/op           0 B/op           0 allocs/op
BenchmarkThreeVariadic-12        350588073             3.490 ns/op           0 B/op           0 allocs/op
BenchmarkFour-12                 494646198             2.167 ns/op           0 B/op           0 allocs/op
BenchmarkFourVariadic-12         296527790             4.156 ns/op           0 B/op           0 allocs/op
PASS
ok      github.com/akramarenkov/goresearch/variadic    11.117s
```
