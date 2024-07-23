# Research on the impact of defer on performance

Running example:

```bash
go test -bench=.* -benchmem ./...
```

Result example:

```bash
goos: linux
goarch: amd64
pkg: github.com/akramarenkov/goresearch/deferred
cpu: AMD Ryzen 5 3600 6-Core Processor              
BenchmarkIdleReference-12            1000000000           0.2404 ns/op        0 B/op           0 allocs/op
BenchmarkIdleNotDeferred-12          1000000000           0.2570 ns/op        0 B/op           0 allocs/op
BenchmarkIdleDeferred-12             494790894            2.528 ns/op         0 B/op           0 allocs/op
BenchmarkOneReference-12             1000000000           0.7324 ns/op        0 B/op           0 allocs/op
BenchmarkOneNotDeferred-12           1000000000           1.285 ns/op         0 B/op           0 allocs/op
BenchmarkOneDeferred-12              400940272            2.920 ns/op         0 B/op           0 allocs/op
BenchmarkTwoReference-12             792121762            1.441 ns/op         0 B/op           0 allocs/op
BenchmarkTwoNotDeferred-12           239845552            4.998 ns/op         0 B/op           0 allocs/op
BenchmarkTwoDeferred-12              223878168            5.213 ns/op         0 B/op           0 allocs/op
BenchmarkFourReference-12            373334564            2.955 ns/op         0 B/op           0 allocs/op
BenchmarkFourNotDeferred-12          121362987            9.861 ns/op         0 B/op           0 allocs/op
BenchmarkFourDeferred-12             100000000           11.20 ns/op          0 B/op           0 allocs/op
BenchmarkSixReference-12             261360696            4.421 ns/op         0 B/op           0 allocs/op
BenchmarkSixNotDeferred-12           74071412            15.48 ns/op          0 B/op           0 allocs/op
BenchmarkSixDeferred-12              74974321            16.38 ns/op          0 B/op           0 allocs/op
BenchmarkEightReference-12           208485084            5.883 ns/op         0 B/op           0 allocs/op
BenchmarkEightNotDeferred-12         58996754            19.98 ns/op          0 B/op           0 allocs/op
BenchmarkEightDeferred-12            54326890            22.58 ns/op          0 B/op           0 allocs/op
BenchmarkTenReference-12             158046957            7.365 ns/op         0 B/op           0 allocs/op
BenchmarkTenNotDeferred-12           49323393            24.95 ns/op          0 B/op           0 allocs/op
BenchmarkTenDeferred-12              10196566           115.2 ns/op           0 B/op           0 allocs/op
BenchmarkElevenReference-12          151637247            8.008 ns/op         0 B/op           0 allocs/op
BenchmarkElevenNotDeferred-12        40563180            28.23 ns/op          0 B/op           0 allocs/op
BenchmarkElevenDeferred-12           10656690           115.4 ns/op           0 B/op           0 allocs/op
PASS
ok      github.com/akramarenkov/goresearch/deferred    36.102s
```

Running example without optimizations:

```bash
go test -bench=.* -benchmem -gcflags=-N ./...
```

Result example without optimizations:

```bash
goos: linux
goarch: amd64
pkg: github.com/akramarenkov/goresearch/deferred
cpu: AMD Ryzen 5 3600 6-Core Processor              
BenchmarkIdleReference-12            1000000000           0.9588 ns/op        0 B/op           0 allocs/op
BenchmarkIdleNotDeferred-12          616774564            1.777 ns/op         0 B/op           0 allocs/op
BenchmarkIdleDeferred-12             52924790            23.67 ns/op          0 B/op           0 allocs/op
BenchmarkOneReference-12             821764441            1.231 ns/op         0 B/op           0 allocs/op
BenchmarkOneNotDeferred-12           418142164            2.839 ns/op         0 B/op           0 allocs/op
BenchmarkOneDeferred-12              49587547            24.29 ns/op          0 B/op           0 allocs/op
BenchmarkTwoReference-12             592550656            1.806 ns/op         0 B/op           0 allocs/op
BenchmarkTwoNotDeferred-12           215584528            5.252 ns/op         0 B/op           0 allocs/op
BenchmarkTwoDeferred-12              31485750            37.43 ns/op          0 B/op           0 allocs/op
BenchmarkFourReference-12            385017488            3.204 ns/op         0 B/op           0 allocs/op
BenchmarkFourNotDeferred-12          116617800           10.48 ns/op          0 B/op           0 allocs/op
BenchmarkFourDeferred-12             20572290            57.79 ns/op          0 B/op           0 allocs/op
BenchmarkSixReference-12             246032240            5.675 ns/op         0 B/op           0 allocs/op
BenchmarkSixNotDeferred-12           77408386            15.46 ns/op          0 B/op           0 allocs/op
BenchmarkSixDeferred-12              14758776            80.76 ns/op          0 B/op           0 allocs/op
BenchmarkEightReference-12           191229417            6.151 ns/op         0 B/op           0 allocs/op
BenchmarkEightNotDeferred-12         57844016            21.28 ns/op          0 B/op           0 allocs/op
BenchmarkEightDeferred-12            12251443            99.64 ns/op          0 B/op           0 allocs/op
BenchmarkTenReference-12             154981694            7.579 ns/op         0 B/op           0 allocs/op
BenchmarkTenNotDeferred-12           46516177            26.48 ns/op          0 B/op           0 allocs/op
BenchmarkTenDeferred-12              10175427           117.1 ns/op           0 B/op           0 allocs/op
BenchmarkElevenReference-12          143367564            8.375 ns/op         0 B/op           0 allocs/op
BenchmarkElevenNotDeferred-12        42253572            29.29 ns/op          0 B/op           0 allocs/op
BenchmarkElevenDeferred-12           8851676            136.7 ns/op           0 B/op           0 allocs/op
PASS
ok      github.com/akramarenkov/goresearch/deferred    36.604s
```
