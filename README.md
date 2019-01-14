## damm checksum algorithm

[![Go Report Card](https://goreportcard.com/badge/jancajthaml-go/damm)](https://goreportcard.com/report/jancajthaml-go/damm)

### Usage ###

```
import "github.com/jancajthaml-go/damm"

ok := damm.Validate("00123014764700968325")

digit, error := damm.Digit("x")

checksum := damm.Generate("1")
```

### Performance ###

```
BenchmarkDammSmall-4           	500000000	        3.80 ns/op	       0 B/op	       0 allocs/op
BenchmarkDammLarge-4           	50000000	        38.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkDammSmallParallel-4   	1000000000	      2.09 ns/op	       0 B/op	       0 allocs/op
BenchmarkDammLargeParallel-4   	100000000	        15.2 ns/op	       0 B/op	       0 allocs/op
```

verify your performance by running `make benchmark`

### Resources ###

* [Wikipedia - Damm algorithm](https://en.wikipedia.org/wiki/Damm_algorithm)
