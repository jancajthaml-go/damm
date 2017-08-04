package main

var m = [10][10]uint{
	0: {0, 3, 1, 7, 5, 9, 8, 6, 4, 2},
	1: {7, 0, 9, 2, 1, 5, 4, 8, 6, 3},
	2: {4, 2, 0, 6, 8, 7, 1, 3, 5, 9},
	3: {1, 7, 5, 0, 9, 8, 3, 4, 2, 6},
	4: {6, 1, 2, 3, 0, 4, 5, 9, 7, 8},
	5: {3, 6, 7, 4, 2, 0, 9, 5, 8, 1},
	6: {5, 8, 6, 9, 7, 2, 0, 1, 3, 4},
	7: {8, 9, 4, 5, 3, 6, 2, 0, 1, 7},
	8: {9, 4, 3, 8, 6, 1, 7, 2, 0, 5},
	9: {2, 5, 8, 1, 4, 3, 6, 7, 9, 0},
}

/**
 * Damm Checksum algorithm
 *
 * @see https://en.wikipedia.org/wiki/Damm_algorithm
 *
 * @author jan.cajthaml
 */
func Damm(cc string) (ok bool) {
	var (
		i int  = 0
		l int  = len(cc)
		x uint = 0
	)

scan:
	x = m[x][int(cc[i])-48]
	i++
	if i == l {
		goto eos
	}
	goto scan

eos:
	return x == 0
}
