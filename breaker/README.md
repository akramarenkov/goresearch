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
BenchmarkChannelOpened-12        367069210             3.235 ns/op           0 B/op           0 allocs/op
BenchmarkChannelClosed-12        302200909             3.721 ns/op           0 B/op           0 allocs/op
BenchmarkAtomic-12               1000000000            0.2717 ns/op          0 B/op           0 allocs/op
PASS
ok      github.com/akramarenkov/goresearch/breaker    3.352s
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
BenchmarkChannelOpened-12        221509768             6.342 ns/op           0 B/op           0 allocs/op
BenchmarkChannelClosed-12        245995996             4.670 ns/op           0 B/op           0 allocs/op
BenchmarkAtomic-12               714703104             1.726 ns/op           0 B/op           0 allocs/op
PASS
ok      github.com/akramarenkov/goresearch/breaker    5.013s
```
