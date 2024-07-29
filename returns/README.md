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
BenchmarkDirect-12                          1000000000           0.7457 ns/op         0 B/op           0 allocs/op
BenchmarkStruct-12                          176745098            6.658 ns/op          0 B/op           0 allocs/op
BenchmarkStructPointer-12                   24288409            48.33 ns/op          48 B/op           1 allocs/op
BenchmarkStructPointerPassthrough-12        842854261            1.242 ns/op          0 B/op           0 allocs/op
PASS
ok      github.com/akramarenkov/goresearch/returns    5.115s
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
BenchmarkDirect-12                          248453119            4.951 ns/op          0 B/op           0 allocs/op
BenchmarkStruct-12                          154329919            7.844 ns/op          0 B/op           0 allocs/op
BenchmarkStructPointer-12                   27789607            51.52 ns/op          48 B/op           1 allocs/op
BenchmarkStructPointerPassthrough-12        381699048            2.996 ns/op          0 B/op           0 allocs/op
PASS
ok      github.com/akramarenkov/goresearch/returns    6.492s
```
