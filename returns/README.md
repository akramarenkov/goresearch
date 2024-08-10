# Performance study of direct and structure-wrapped return of output arguments

Running example:

```bash
go test -bench=.* -benchmem ./...
```

Result example:

```bash
goos: linux
goarch: amd64
pkg: github.com/akramarenkov/goresearch/returns
cpu: AMD Ryzen 5 3600 6-Core Processor              
BenchmarkDirect-12                          1000000000             0.5833 ns/op         0 B/op           0 allocs/op
BenchmarkStruct-12                          184950351              6.591 ns/op          0 B/op           0 allocs/op
BenchmarkStructPointer-12                   29304154              46.21 ns/op          48 B/op           1 allocs/op
BenchmarkStructPointerPassthrough-12        779165852              1.472 ns/op          0 B/op           0 allocs/op
PASS
ok      github.com/akramarenkov/goresearch/returns    5.234s
```

Running example without optimizations:

```bash
go test -bench=.* -benchmem -gcflags=-N ./...
```

Result example without optimizations:

```bash
goos: linux
goarch: amd64
pkg: github.com/akramarenkov/goresearch/returns
cpu: AMD Ryzen 5 3600 6-Core Processor              
BenchmarkDirect-12                          246946584             5.064 ns/op          0 B/op           0 allocs/op
BenchmarkStruct-12                          146410693             8.040 ns/op          0 B/op           0 allocs/op
BenchmarkStructPointer-12                   28596062             50.33 ns/op          48 B/op           1 allocs/op
BenchmarkStructPointerPassthrough-12        289283244             3.891 ns/op          0 B/op           0 allocs/op
PASS
ok      github.com/akramarenkov/goresearch/returns    6.793s
```
