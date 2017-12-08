## Performant and straight-forward implementation of damm checksum algorithm

[![Go Report Card](https://goreportcard.com/badge/jancajthaml-go/damm)](https://goreportcard.com/report/jancajthaml-go/damm)

### Usage ###

```
import "github.com/jancajthaml-go/damm"

ok := damm.Validate("00123014764700968325")

digit, error := damm.Digit("x")

signed := damm.Generate("1")
```

### Performance ###

```
BenchmarkDammSmall-4            500000000         3.98 ns/op
BenchmarkDammLarge-4            50000000          39.4 ns/op
BenchmarkDammSmallParallel-4    1000000000        1.99 ns/op
BenchmarkDammLargeParallel-4    100000000         14.7 ns/op
```

test on your own by running `make benchmark`

### Resources ###

* [Wikipedia - Damm algorithm](https://en.wikipedia.org/wiki/Damm_algorithm)
