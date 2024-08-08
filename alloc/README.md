# Performance comparison of increasing slice length using append and make

Running example:

```bash
for bench in $( go test -list ".*" | grep "Benchmark" ); do go test -run=^$ "-bench=^${bench}$" -benchmem -benchtime=1x '.'  | grep "Benchmark"; sleep 2; done
```

Result example:

```bash
BenchmarkGrowAppend_0_0_10-12               1           381.0 ns/op        1024 B/op           1 allocs/op
BenchmarkGrowRemake_0_0_10-12               1           321.0 ns/op        1024 B/op           1 allocs/op
BenchmarkGrowAppend_0_0_20-12               1        670332 ns/op       1048576 B/op           1 allocs/op
BenchmarkGrowRemake_0_0_20-12               1         23044 ns/op       1048576 B/op           1 allocs/op
BenchmarkGrowAppend_0_0_29-12               1     247323515 ns/op     536870912 B/op           1 allocs/op
BenchmarkGrowRemake_0_0_29-12               1        412298 ns/op     536870920 B/op           2 allocs/op
BenchmarkGrowAppend_0_0_30-12               1     475732047 ns/op    1073745120 B/op           6 allocs/op
BenchmarkGrowRemake_0_0_30-12               1        928963 ns/op    1073741832 B/op           2 allocs/op
BenchmarkGrowAppend_0_0_31-12               1     944742222 ns/op    2147486944 B/op           6 allocs/op
BenchmarkGrowRemake_0_0_31-12               1       1508550 ns/op    2147483648 B/op           1 allocs/op
BenchmarkGrowAppend_0_0_32-12               1    1879972629 ns/op    4294970496 B/op           5 allocs/op
BenchmarkGrowRemake_0_0_32-12               1    1669865615 ns/op    4294967296 B/op           1 allocs/op
BenchmarkGrowAppend_0_0_33-12               1    3766471135 ns/op    8589937992 B/op           8 allocs/op
BenchmarkGrowRemake_0_0_33-12               1       5168696 ns/op    8589934592 B/op           1 allocs/op
BenchmarkGrowAppend_8_0_10-12               1           481.0 ns/op        1024 B/op           1 allocs/op
BenchmarkGrowRemake_8_0_10-12               1           481.0 ns/op        1024 B/op           1 allocs/op
BenchmarkGrowAppend_18_0_20-12              1        433758 ns/op       1048576 B/op           1 allocs/op
BenchmarkGrowRemake_18_0_20-12              1         21691 ns/op       1048576 B/op           1 allocs/op
BenchmarkGrowAppend_27_0_29-12              1     237050174 ns/op     536871104 B/op           3 allocs/op
BenchmarkGrowRemake_27_0_29-12              1     212481727 ns/op     536870920 B/op           2 allocs/op
BenchmarkGrowAppend_28_0_30-12              1     470653798 ns/op    1073741824 B/op           1 allocs/op
BenchmarkGrowRemake_28_0_30-12              1     421702971 ns/op    1073741824 B/op           1 allocs/op
BenchmarkGrowAppend_29_0_31-12              1     946562441 ns/op    2147486848 B/op           5 allocs/op
BenchmarkGrowRemake_29_0_31-12              1     654548821 ns/op    2147486944 B/op           6 allocs/op
BenchmarkGrowAppend_30_0_32-12              1    1900882898 ns/op    4294970592 B/op           6 allocs/op
BenchmarkGrowRemake_30_0_32-12              1    1661657284 ns/op    4294970600 B/op           7 allocs/op
BenchmarkGrowAppend_31_0_33-12              1    3761617881 ns/op    8589937984 B/op           7 allocs/op
BenchmarkGrowRemake_31_0_33-12              1    3329993613 ns/op    8589934592 B/op           1 allocs/op
BenchmarkGrowAppend_9_0_10-12               1          2144 ns/op          1280 B/op           1 allocs/op
BenchmarkGrowRemake_9_0_10-12               1           331.0 ns/op        1024 B/op           1 allocs/op
BenchmarkGrowAppend_19_0_20-12              1        513080 ns/op       1286144 B/op           1 allocs/op
BenchmarkGrowRemake_19_0_20-12              1          8516 ns/op       1048576 B/op           1 allocs/op
BenchmarkGrowAppend_28_0_29-12              1     290065923 ns/op     655371392 B/op           5 allocs/op
BenchmarkGrowRemake_28_0_29-12              1     213958962 ns/op     536871008 B/op           2 allocs/op
BenchmarkGrowAppend_29_0_30-12              1     578851000 ns/op    1310731488 B/op           6 allocs/op
BenchmarkGrowRemake_29_0_30-12              1     421103937 ns/op    1073745120 B/op           6 allocs/op
BenchmarkGrowAppend_30_0_31-12              1    1158314880 ns/op    2621451392 B/op           5 allocs/op
BenchmarkGrowRemake_30_0_31-12              1     844724278 ns/op    2147486944 B/op           6 allocs/op
BenchmarkGrowAppend_31_0_32-12              1    2297741812 ns/op    5242891488 B/op           6 allocs/op
BenchmarkGrowRemake_31_0_32-12              1    1670259177 ns/op    4294970600 B/op           7 allocs/op
BenchmarkGrowAppend_32_0_33-12              1    4618210227 ns/op   10485771488 B/op           6 allocs/op
BenchmarkGrowRemake_32_0_33-12              1    3338581454 ns/op    8589934592 B/op           1 allocs/op
BenchmarkGrowAppend_9_128_10-12             1          2766 ns/op          1536 B/op           1 allocs/op
BenchmarkGrowRemake_9_128_10-12             1           271.0 ns/op        1024 B/op           1 allocs/op
BenchmarkGrowAppend_9_129_10-12             1           461.0 ns/op        1152 B/op           1 allocs/op
BenchmarkGrowRemake_9_129_10-12             1           261.0 ns/op        1024 B/op           1 allocs/op
BenchmarkGrowAppend_19_516096_20-12         1        607720 ns/op       1302528 B/op           1 allocs/op
BenchmarkGrowRemake_19_516096_20-12         1         73921 ns/op       1048576 B/op           1 allocs/op
BenchmarkGrowAppend_19_516097_20-12         1           350.0 ns/op           0 B/op           0 allocs/op
BenchmarkGrowRemake_19_516097_20-12         1         58011 ns/op       1048576 B/op           1 allocs/op
BenchmarkGrowAppend_28_268427264_29-12      1     297635323 ns/op     671084032 B/op           9 allocs/op
BenchmarkGrowRemake_28_268427264_29-12      1     216119785 ns/op     536870912 B/op           1 allocs/op
BenchmarkGrowAppend_28_268427265_29-12      1           851.0 ns/op           0 B/op           0 allocs/op
BenchmarkGrowRemake_28_268427265_29-12      1     212680669 ns/op     536874120 B/op           6 allocs/op
BenchmarkGrowAppend_29_536862720_30-12      1     608533919 ns/op    1342172480 B/op           7 allocs/op
BenchmarkGrowRemake_29_536862720_30-12      1     424307876 ns/op    1073745120 B/op           6 allocs/op
BenchmarkGrowAppend_29_536862721_30-12      1          1333 ns/op             0 B/op           0 allocs/op
BenchmarkGrowRemake_29_536862721_30-12      1     426919061 ns/op    1073745120 B/op           6 allocs/op
BenchmarkGrowAppend_30_1073733632_31-12     1    1191303049 ns/op    2684349568 B/op           5 allocs/op
BenchmarkGrowRemake_30_1073733632_31-12     1     840410797 ns/op    2147486856 B/op           6 allocs/op
BenchmarkGrowAppend_30_1073733633_31-12     1          1062 ns/op             0 B/op           0 allocs/op
BenchmarkGrowRemake_30_1073733633_31-12     1     837067661 ns/op    2147487040 B/op           7 allocs/op
BenchmarkGrowAppend_31_2147475456_32-12     1    2382382925 ns/op    5368704224 B/op           6 allocs/op
BenchmarkGrowRemake_31_2147475456_32-12     1    1653759515 ns/op    4294970504 B/op           6 allocs/op
BenchmarkGrowAppend_31_2147475457_32-12     1          1032 ns/op             0 B/op           0 allocs/op
BenchmarkGrowRemake_31_2147475457_32-12     1    1657953654 ns/op    4294970496 B/op           5 allocs/op
BenchmarkGrowAppend_32_4294959104_33-12     1    4882759428 ns/op   10737410144 B/op           2 allocs/op
BenchmarkGrowRemake_32_4294959104_33-12     1    3311962635 ns/op    8589941096 B/op          11 allocs/op
BenchmarkGrowAppend_32_4294959105_33-12     1          1123 ns/op             0 B/op           0 allocs/op
BenchmarkGrowRemake_32_4294959105_33-12     1    3333256132 ns/op    8589937896 B/op           7 allocs/op
```

Running example without optimizations:

```bash
for bench in $( go test -list ".*" | grep "Benchmark" ); do go test -run=^$ "-bench=^${bench}$" -benchmem -benchtime=1x -gcflags=-N  '.'  | grep "Benchmark"; sleep 2; done
```

Result example without optimizations:

```bash
BenchmarkGrowAppend_0_0_10-12               1           541.0 ns/op        2048 B/op           2 allocs/op
BenchmarkGrowRemake_0_0_10-12               1           340.0 ns/op        1024 B/op           1 allocs/op
BenchmarkGrowAppend_0_0_20-12               1        682783 ns/op       2097152 B/op           2 allocs/op
BenchmarkGrowRemake_0_0_20-12               1         26190 ns/op       1048576 B/op           1 allocs/op
BenchmarkGrowAppend_0_0_29-12               1     343669802 ns/op    1073748416 B/op          12 allocs/op
BenchmarkGrowRemake_0_0_29-12               1        417246 ns/op     536870912 B/op           1 allocs/op
BenchmarkGrowAppend_0_0_30-12               1     676430478 ns/op    2147490144 B/op          11 allocs/op
BenchmarkGrowRemake_0_0_30-12               1        867136 ns/op    1073741832 B/op           2 allocs/op
BenchmarkGrowAppend_0_0_31-12               1    1819177301 ns/op    4294970688 B/op           8 allocs/op
BenchmarkGrowRemake_0_0_31-12               1       1370235 ns/op    2147483648 B/op           1 allocs/op
BenchmarkGrowAppend_0_0_32-12               1    2706150791 ns/op    8589941280 B/op          13 allocs/op
BenchmarkGrowRemake_0_0_32-12               1       2760119 ns/op    4294967304 B/op           2 allocs/op
BenchmarkGrowAppend_0_0_33-12               1    5401486878 ns/op   17179878784 B/op          14 allocs/op
BenchmarkGrowRemake_0_0_33-12               1       4946232 ns/op    8589934592 B/op           1 allocs/op
BenchmarkGrowAppend_8_0_10-12               1          1863 ns/op          1792 B/op           2 allocs/op
BenchmarkGrowRemake_8_0_10-12               1           311.0 ns/op        1024 B/op           1 allocs/op
BenchmarkGrowAppend_18_0_20-12              1        572753 ns/op       1835008 B/op           2 allocs/op
BenchmarkGrowRemake_18_0_20-12              1          8616 ns/op       1048576 B/op           1 allocs/op
BenchmarkGrowAppend_27_0_29-12              1     329200200 ns/op     939527392 B/op           7 allocs/op
BenchmarkGrowRemake_27_0_29-12              1     210063677 ns/op     536874120 B/op           6 allocs/op
BenchmarkGrowAppend_28_0_30-12              1     625820898 ns/op    1879051392 B/op           6 allocs/op
BenchmarkGrowRemake_28_0_30-12              1     418270459 ns/op    1073745024 B/op           5 allocs/op
BenchmarkGrowAppend_29_0_31-12              1    1274873435 ns/op    3758099872 B/op           9 allocs/op
BenchmarkGrowRemake_29_0_31-12              1     831255118 ns/op    2147486856 B/op           6 allocs/op
BenchmarkGrowAppend_30_0_32-12              1    2506210733 ns/op    7516196064 B/op           7 allocs/op
BenchmarkGrowRemake_30_0_32-12              1    1688575194 ns/op    4294967400 B/op           3 allocs/op
BenchmarkGrowAppend_31_0_33-12              1    4957899968 ns/op   15032392224 B/op          13 allocs/op
BenchmarkGrowRemake_31_0_33-12              1    3317406610 ns/op    8589937888 B/op           6 allocs/op
BenchmarkGrowAppend_9_0_10-12               1          2445 ns/op          1792 B/op           2 allocs/op
BenchmarkGrowRemake_9_0_10-12               1           401.0 ns/op        1024 B/op           1 allocs/op
BenchmarkGrowAppend_19_0_20-12              1        628670 ns/op       1810432 B/op           2 allocs/op
BenchmarkGrowRemake_19_0_20-12              1         70635 ns/op       1048576 B/op           1 allocs/op
BenchmarkGrowAppend_28_0_29-12              1     418665388 ns/op     923807136 B/op           9 allocs/op
BenchmarkGrowRemake_28_0_29-12              1     212623396 ns/op     536874112 B/op           5 allocs/op
BenchmarkGrowAppend_29_0_30-12              1     830071318 ns/op    1847602496 B/op           8 allocs/op
BenchmarkGrowRemake_29_0_30-12              1     414894590 ns/op    1073745024 B/op           5 allocs/op
BenchmarkGrowAppend_30_0_31-12              1    1653001395 ns/op    3695190016 B/op           2 allocs/op
BenchmarkGrowRemake_30_0_31-12              1     833658876 ns/op    2147486944 B/op           6 allocs/op
BenchmarkGrowAppend_31_0_32-12              1    3323513612 ns/op    7390371936 B/op           3 allocs/op
BenchmarkGrowRemake_31_0_32-12              1    1671404322 ns/op    4294970504 B/op           6 allocs/op
BenchmarkGrowAppend_32_0_33-12              1    6578879084 ns/op   14780735584 B/op           3 allocs/op
BenchmarkGrowRemake_32_0_33-12              1    3316485198 ns/op    8589937888 B/op           6 allocs/op
BenchmarkGrowAppend_9_128_10-12             1          5220 ns/op          1920 B/op           2 allocs/op
BenchmarkGrowRemake_9_128_10-12             1           501.0 ns/op        1024 B/op           1 allocs/op
BenchmarkGrowAppend_9_129_10-12             1          2475 ns/op          1536 B/op           2 allocs/op
BenchmarkGrowRemake_9_129_10-12             1           831.0 ns/op        1024 B/op           1 allocs/op
BenchmarkGrowAppend_19_516096_20-12         1        517368 ns/op       1310720 B/op           2 allocs/op
BenchmarkGrowRemake_19_516096_20-12         1         73931 ns/op       1048576 B/op           1 allocs/op
BenchmarkGrowAppend_19_516097_20-12         1          4920 ns/op          8192 B/op           1 allocs/op
BenchmarkGrowRemake_19_516097_20-12         1         58542 ns/op       1048584 B/op           2 allocs/op
BenchmarkGrowAppend_28_268427264_29-12      1     295778205 ns/op     671088736 B/op           3 allocs/op
BenchmarkGrowRemake_28_268427264_29-12      1     209979323 ns/op     536874120 B/op           6 allocs/op
BenchmarkGrowAppend_28_268427265_29-12      1         37211 ns/op          8192 B/op           1 allocs/op
BenchmarkGrowRemake_28_268427265_29-12      1      36842179 ns/op     536870920 B/op           2 allocs/op
BenchmarkGrowAppend_29_536862720_30-12      1     590879188 ns/op    1342177280 B/op           2 allocs/op
BenchmarkGrowRemake_29_536862720_30-12      1     415816718 ns/op    1073745032 B/op           6 allocs/op
BenchmarkGrowAppend_29_536862721_30-12      1         45487 ns/op          8288 B/op           2 allocs/op
BenchmarkGrowRemake_29_536862721_30-12      1     415373362 ns/op    1073745032 B/op           6 allocs/op
BenchmarkGrowAppend_30_1073733632_31-12     1    1187188353 ns/op    2684354560 B/op           2 allocs/op
BenchmarkGrowRemake_30_1073733632_31-12     1     831150336 ns/op    2147486848 B/op           5 allocs/op
BenchmarkGrowAppend_30_1073733633_31-12     1         10570 ns/op          8192 B/op           1 allocs/op
BenchmarkGrowRemake_30_1073733633_31-12     1     831333796 ns/op    2147486952 B/op           7 allocs/op
BenchmarkGrowAppend_31_2147475456_32-12     1    2417969119 ns/op    5368712320 B/op           6 allocs/op
BenchmarkGrowRemake_31_2147475456_32-12     1    1666346324 ns/op    4294970600 B/op           7 allocs/op
BenchmarkGrowAppend_31_2147475457_32-12     1         43092 ns/op          8384 B/op           3 allocs/op
BenchmarkGrowRemake_31_2147475457_32-12     1    1655923498 ns/op    4294970496 B/op           5 allocs/op
BenchmarkGrowAppend_32_4294959104_33-12     1    4855751239 ns/op   10737418336 B/op           3 allocs/op
BenchmarkGrowRemake_32_4294959104_33-12     1    3340839134 ns/op    8589937800 B/op           6 allocs/op
BenchmarkGrowAppend_32_4294959105_33-12     1          8957 ns/op          8192 B/op           1 allocs/op
BenchmarkGrowRemake_32_4294959105_33-12     1    3321152201 ns/op    8589937992 B/op           8 allocs/op
```
