package main

import (
	"fmt"
	"math"
)

type pair struct {
	i, j int
}

func iabs(x int) int {
	return int(math.Abs(float64(x)))
}

func main() {
	n := 0
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		s := 0
		fmt.Scan(&s)
		m := make([]pair, s*s)
		for j := 0; j < s; j++ {
			for k := 0; k < s; k++ {
				r := 0
				fmt.Scan(&r)
				m[r-1] = pair{j, k}
			}
		}
		rc, rm, dc, dm, pi, pj := 0, 0, -1, 0, -1, -1
		for r, p := range m {
			i, j := p.i, p.j
			if (i == pi && iabs(j-pj) <= 1) || (j == pj && iabs(i-pi) <= 1) {
				dc++
			} else {
				rc, dc = r, 0
			}
			if dc > dm {
				rm, dm = rc, dc
			}
			pi, pj = i, j
		}
		fmt.Printf("Case #%d: %v %v\n", i+1, rm+1, dm+1)
	}
}
