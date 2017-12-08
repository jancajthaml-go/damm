// Package damm supports the computation of a decimal checksum. It uses a
// method proposed by H. Michael Damm in 2004. The checksum doesn't change for
// leading zeroes.
//
// Information about the algorithm is available on Wikipedia
//
// http://en.wikipedia.org/wiki/Damm_algorithm
//
package main

import "errors"

var m = [160]uint{0, 3, 1, 7, 5, 9, 8, 6, 4, 2, 0, 0, 0, 0, 0, 0, 7, 0, 9, 2, 1, 5, 4, 8, 6, 3, 0, 0, 0, 0, 0, 0, 4, 2, 0, 6, 8, 7, 1, 3, 5, 9, 0, 0, 0, 0, 0, 0, 1, 7, 5, 0, 9, 8, 3, 4, 2, 6, 0, 0, 0, 0, 0, 0, 6, 1, 2, 3, 0, 4, 5, 9, 7, 8, 0, 0, 0, 0, 0, 0, 3, 6, 7, 4, 2, 0, 9, 5, 8, 1, 0, 0, 0, 0, 0, 0, 5, 8, 6, 9, 7, 2, 0, 1, 3, 4, 0, 0, 0, 0, 0, 0, 8, 9, 4, 5, 3, 6, 2, 0, 1, 7, 0, 0, 0, 0, 0, 0, 9, 4, 3, 8, 6, 1, 7, 2, 0, 5, 0, 0, 0, 0, 0, 0, 2, 5, 8, 1, 4, 3, 6, 7, 9, 0, 0, 0, 0, 0, 0, 0}

// Digit returns damm digit for given numeric string
func Digit(cc string) (uint, error) {
	var (
		d uint
		i int
		x uint
		l int = len(cc)
	)
scan:
	d = uint(cc[i]) - 48
	if d > 9 {
		return 1, errors.New("string must contain only digits")
	}
	x = m[x<<4+d]
	i++
	if i != l {
		goto scan
	}
	return x, nil
}

// Validate returns true if numeric string is signed with valid damm digit
func Validate(cc string) (ok bool) {
	digit, err := Digit(cc)
	return err == nil && digit == 0
}

// Generate signs string with damm digit
func Generate(cc string) (string, error) {
	digit, err := Digit(cc)

	if err != nil {
		return cc, err
	}

	return cc + string(digit+48), nil
}
