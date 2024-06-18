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
BenchmarkChannelOpened-12        366570735             3.230 ns/op           0 B/op           0 allocs/op
BenchmarkChannelClosed-12        305249197             3.602 ns/op           0 B/op           0 allocs/op
BenchmarkAtomic-12               1000000000            0.2718 ns/op          0 B/op           0 allocs/op
PASS
ok      github.com/akramarenkov/goresearch/breaker    3.341s
```
