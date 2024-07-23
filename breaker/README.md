# Fastest breaker research

Running example:

```bash
go test -bench=.* -benchmem ./...
```

Result example:

```bash
goos: linux
goarch: amd64
pkg: github.com/akramarenkov/goresearch/breaker
cpu: AMD Ryzen 5 3600 6-Core Processor              
BenchmarkIdle-12                 1000000000            0.2622 ns/op          0 B/op           0 allocs/op
BenchmarkChannelOpened-12        384499546             3.222 ns/op           0 B/op           0 allocs/op
BenchmarkChannelClosed-12        305772325             3.736 ns/op           0 B/op           0 allocs/op
BenchmarkAtomic-12               1000000000            0.2755 ns/op          0 B/op           0 allocs/op
PASS
ok      github.com/akramarenkov/goresearch/breaker    3.702s
```

Running example without optimizations:

```bash
go test -bench=.* -benchmem -gcflags=-N ./...
```

Result example without optimizations:

```bash
goos: linux
goarch: amd64
pkg: github.com/akramarenkov/goresearch/breaker
cpu: AMD Ryzen 5 3600 6-Core Processor              
BenchmarkIdle-12                 1000000000            0.9879 ns/op          0 B/op           0 allocs/op
BenchmarkChannelOpened-12        240061744             5.896 ns/op           0 B/op           0 allocs/op
BenchmarkChannelClosed-12        263083366             4.617 ns/op           0 B/op           0 allocs/op
BenchmarkAtomic-12               551909416             2.164 ns/op           0 B/op           0 allocs/op
PASS
ok      github.com/akramarenkov/goresearch/breaker    7.195s
```
