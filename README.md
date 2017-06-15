# Autocsv

Generate autocompletion API from CSV data.

# Usage

```
Usage ./autocsv [options]
  -file string
        File datasource
  -needed-fields string
        Expected fields
  -search-field string
        Searchable field
  -separator string
        csv field separator (default ";")

```

# Endpoint

`/autocomplete?q=queryword`

# Benchmark

```
BenchmarkSearchA-4                         10000            209261 ns/op           35496 B/op        218 allocs/op
BenchmarkSearchMai-4                       10000            161396 ns/op           16664 B/op         75 allocs/op
BenchmarkSearchMairie-4                    10000            141730 ns/op           12728 B/op         44 allocs/op
BenchmarkSearchMairieUpper-4               10000            143580 ns/op           12744 B/op         46 allocs/op
BenchmarkSearchEgalite-4                   10000            133071 ns/op            9176 B/op         13 allocs/op
BenchmarkSearchNotFound-4                  10000            126792 ns/op            8800 B/op          9 allocs/op
BenchmarkSearchRegexA-4                     2000           1089574 ns/op           78992 B/op        269 allocs/op
BenchmarkSearchRegexMai-4                   2000           1049027 ns/op           59224 B/op        115 allocs/op
BenchmarkSearchRegexMairie-4                2000           1059808 ns/op           56920 B/op         87 allocs/op
BenchmarkSearchRegexMairieUpper-4           2000           1040144 ns/op           56936 B/op         89 allocs/op
BenchmarkSearchRegexEgalite-4               2000           1079702 ns/op           53448 B/op         56 allocs/op
BenchmarkSearchRegexNotFound-4              2000           1036235 ns/op           53361 B/op         54 allocs/op
PASS
ok      github.com/theodelaune/autocsv  22.656s
```


