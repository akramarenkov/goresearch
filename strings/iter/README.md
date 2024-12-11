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
BenchmarkIterLen1-12     171784060          6.883 ns/op        0 B/op        0 allocs/op
BenchmarkIterByte1-12    175178682          6.921 ns/op        0 B/op        0 allocs/op
BenchmarkIterRune1-12    217542111          5.533 ns/op        0 B/op        0 allocs/op
BenchmarkIterLen2-12     173568162          6.898 ns/op        0 B/op        0 allocs/op
BenchmarkIterByte2-12    175265710          6.809 ns/op        0 B/op        0 allocs/op
BenchmarkIterRune2-12    135241789          8.554 ns/op        0 B/op        0 allocs/op
BenchmarkIterLen3-12     126311188          9.615 ns/op        0 B/op        0 allocs/op
BenchmarkIterByte3-12    124640461         10.28  ns/op        0 B/op        0 allocs/op
BenchmarkIterRune3-12     60520033         18.55  ns/op        0 B/op        0 allocs/op
BenchmarkIterLen4-12      99666962         12.38  ns/op        0 B/op        0 allocs/op
BenchmarkIterByte4-12    100000000         11.79  ns/op        0 B/op        0 allocs/op
BenchmarkIterRune4-12     61135147         17.46  ns/op        0 B/op        0 allocs/op
BenchmarkIterLen5-12      80542312         15.39  ns/op        0 B/op        0 allocs/op
BenchmarkIterByte5-12     80098851         15.20  ns/op        0 B/op        0 allocs/op
BenchmarkIterRune5-12     38337985         28.93  ns/op        0 B/op        0 allocs/op
BenchmarkIterLen6-12      67770738         18.38  ns/op        0 B/op        0 allocs/op
BenchmarkIterByte6-12     67665331         17.86  ns/op        0 B/op        0 allocs/op
BenchmarkIterRune6-12     48639446         25.33  ns/op        0 B/op        0 allocs/op
BenchmarkIterLen7-12      59707495         19.80  ns/op        0 B/op        0 allocs/op
BenchmarkIterByte7-12     64772253         20.15  ns/op        0 B/op        0 allocs/op
BenchmarkIterRune7-12     29500804         38.40  ns/op        0 B/op        0 allocs/op
BenchmarkIterLen8-12      36189931         32.12  ns/op        0 B/op        0 allocs/op
BenchmarkIterByte8-12     36495513         32.57  ns/op        0 B/op        0 allocs/op
BenchmarkIterRune8-12     34689189         31.15  ns/op        0 B/op        0 allocs/op
BenchmarkIterLen9-12      43218392         33.45  ns/op        0 B/op        0 allocs/op
BenchmarkIterByte9-12     37201489         31.13  ns/op        0 B/op        0 allocs/op
BenchmarkIterRune9-12     23202889         48.70  ns/op        0 B/op        0 allocs/op
BenchmarkIterLen10-12     40737453         33.03  ns/op        0 B/op        0 allocs/op
BenchmarkIterByte10-12    31781962         32.97  ns/op        0 B/op        0 allocs/op
BenchmarkIterRune10-12    26357806         45.26  ns/op        0 B/op        0 allocs/op
PASS
ok   github.com/akramarenkov/goresearch/strings/iter 42.685s
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
BenchmarkIterLen1-12      104147431         11.06 ns/op        0 B/op        0 allocs/op
BenchmarkIterByte1-12     118581488         10.69 ns/op        0 B/op        0 allocs/op
BenchmarkIterRune1-12      67359666         17.91 ns/op        0 B/op        0 allocs/op
BenchmarkIterLen2-12       48647131         24.42 ns/op        0 B/op        0 allocs/op
BenchmarkIterByte2-12      64392363         18.87 ns/op        0 B/op        0 allocs/op
BenchmarkIterRune2-12      24995905         47.47 ns/op        0 B/op        0 allocs/op
BenchmarkIterLen3-12       36225318         33.83 ns/op        0 B/op        0 allocs/op
BenchmarkIterByte3-12      44782188         26.08 ns/op        0 B/op        0 allocs/op
BenchmarkIterRune3-12      23594517         62.47 ns/op        0 B/op        0 allocs/op
BenchmarkIterLen4-12       23476396         51.36 ns/op        0 B/op        0 allocs/op
BenchmarkIterByte4-12      27111330         45.05 ns/op        0 B/op        0 allocs/op
BenchmarkIterRune4-12      14438763         83.74 ns/op        0 B/op        0 allocs/op
BenchmarkIterLen5-12       20914140         76.84 ns/op        0 B/op        0 allocs/op
BenchmarkIterByte5-12      20159379         53.40 ns/op        0 B/op        0 allocs/op
BenchmarkIterRune5-12      12247974         97.31 ns/op        0 B/op        0 allocs/op
BenchmarkIterLen6-12       16623404         79.81 ns/op        0 B/op        0 allocs/op
BenchmarkIterByte6-12      19650949         60.06 ns/op        0 B/op        0 allocs/op
BenchmarkIterRune6-12      11066630        110.5  ns/op        0 B/op        0 allocs/op
BenchmarkIterLen7-12       18082068         96.89 ns/op        0 B/op        0 allocs/op
BenchmarkIterByte7-12      16677696         71.96 ns/op        0 B/op        0 allocs/op
BenchmarkIterRune7-12       8014578        128.7  ns/op        0 B/op        0 allocs/op
BenchmarkIterLen8-12       13012264         96.36 ns/op        0 B/op        0 allocs/op
BenchmarkIterByte8-12      15477874         76.05 ns/op        0 B/op        0 allocs/op
BenchmarkIterRune8-12       7828764        147.8  ns/op        0 B/op        0 allocs/op
BenchmarkIterLen9-12       13894723         85.40 ns/op        0 B/op        0 allocs/op
BenchmarkIterByte9-12      12947557         88.94 ns/op        0 B/op        0 allocs/op
BenchmarkIterRune9-12       6589920        171.5  ns/op        0 B/op        0 allocs/op
BenchmarkIterLen10-12      10181126        115.9  ns/op        0 B/op        0 allocs/op
BenchmarkIterByte10-12     14186715         93.47 ns/op        0 B/op        0 allocs/op
BenchmarkIterRune10-12      6246962        194.3  ns/op        0 B/op        0 allocs/op
PASS
ok   github.com/akramarenkov/goresearch/strings/iter 42.267s
```
