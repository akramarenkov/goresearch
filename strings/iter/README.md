# Research of iteration by string

Running example:

```bash
go test -bench=.* -benchmem ./...
```

Result example:

```bash
goos: linux
goarch: amd64
pkg: github.com/akramarenkov/goresearch/strings/iter
cpu: AMD Ryzen 5 3600 6-Core Processor              
BenchmarkIterByte1-12      356452321         3.381 ns/op       0 B/op        0 allocs/op
BenchmarkIterRune1-12      170039088         7.108 ns/op       0 B/op        0 allocs/op
BenchmarkIterByte2-12      185622195         6.391 ns/op       0 B/op        0 allocs/op
BenchmarkIterRune2-12      148367166         8.099 ns/op       0 B/op        0 allocs/op
BenchmarkIterByte3-12      133930953         9.008 ns/op       0 B/op        0 allocs/op
BenchmarkIterRune3-12       84257686        12.14  ns/op       0 B/op        0 allocs/op
BenchmarkIterByte4-12       56060012        21.57  ns/op       0 B/op        0 allocs/op
BenchmarkIterRune4-12       79492867        16.24  ns/op       0 B/op        0 allocs/op
BenchmarkIterByte5-12       83891607        13.94  ns/op       0 B/op        0 allocs/op
BenchmarkIterRune5-12       61546449        20.18  ns/op       0 B/op        0 allocs/op
BenchmarkIterByte6-12       72871057        16.39  ns/op       0 B/op        0 allocs/op
BenchmarkIterRune6-12       49303701        23.16  ns/op       0 B/op        0 allocs/op
BenchmarkIterByte7-12       64479295        18.84  ns/op       0 B/op        0 allocs/op
BenchmarkIterRune7-12       41690031        26.70  ns/op       0 B/op        0 allocs/op
BenchmarkIterByte8-12       38568976        31.41  ns/op       0 B/op        0 allocs/op
BenchmarkIterRune8-12       39003493        30.59  ns/op       0 B/op        0 allocs/op
BenchmarkIterByte9-12       39360636        29.65  ns/op       0 B/op        0 allocs/op
BenchmarkIterRune9-12       36023589        33.83  ns/op       0 B/op        0 allocs/op
BenchmarkIterByte10-12      39029484        29.63  ns/op       0 B/op        0 allocs/op
BenchmarkIterRune10-12      28325775        42.48  ns/op       0 B/op        0 allocs/op
PASS
ok   github.com/akramarenkov/goresearch/strings/iter 27.592s
```

Running example without optimizations:

```bash
go test -bench=.* -benchmem -gcflags=-N ./...
```

Result example without optimizations:

```bash
goos: linux
goarch: amd64
pkg: github.com/akramarenkov/goresearch/strings/iter
cpu: AMD Ryzen 5 3600 6-Core Processor              
BenchmarkIterByte1-12      120176977         9.863 ns/op      0 B/op        0 allocs/op
BenchmarkIterRune1-12       64223108        17.54  ns/op      0 B/op        0 allocs/op
BenchmarkIterByte2-12       52027974        25.78  ns/op      0 B/op        0 allocs/op
BenchmarkIterRune2-12       27865891        41.68  ns/op      0 B/op        0 allocs/op
BenchmarkIterByte3-12       31477245        46.52  ns/op      0 B/op        0 allocs/op
BenchmarkIterRune3-12       19048593        58.95  ns/op      0 B/op        0 allocs/op
BenchmarkIterByte4-12       25187439        53.27  ns/op      0 B/op        0 allocs/op
BenchmarkIterRune4-12       15187072        75.42  ns/op      0 B/op        0 allocs/op
BenchmarkIterByte5-12       23575368        55.79  ns/op      0 B/op        0 allocs/op
BenchmarkIterRune5-12       12668823        83.33  ns/op      0 B/op        0 allocs/op
BenchmarkIterByte6-12       18054356        69.45  ns/op      0 B/op        0 allocs/op
BenchmarkIterRune6-12       10741094       110.3   ns/op      0 B/op        0 allocs/op
BenchmarkIterByte7-12       17189540        69.93  ns/op      0 B/op        0 allocs/op
BenchmarkIterRune7-12       10210806       131.8   ns/op      0 B/op        0 allocs/op
BenchmarkIterByte8-12       10068786       113.2   ns/op      0 B/op        0 allocs/op
BenchmarkIterRune8-12        8207583       145.1   ns/op      0 B/op        0 allocs/op
BenchmarkIterByte9-12       15192942       100.6   ns/op      0 B/op        0 allocs/op
BenchmarkIterRune9-12        8288505       153.7   ns/op      0 B/op        0 allocs/op
BenchmarkIterByte10-12      11353393       115.3   ns/op      0 B/op        0 allocs/op
BenchmarkIterRune10-12       6821952       176.8   ns/op      0 B/op        0 allocs/op
PASS
ok   github.com/akramarenkov/goresearch/strings/iter 28.339s
```
