package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		var k, v int
		fmt.Scanln(&k, &v)
		fmt.Printf("Case #%d: %v\n", i+1, solve(k, v))
	}
}

func solve(k, v int) int32 {
	var t int32
	fv := float64(v)
	for r := 0; r <= k; r++ {
		for g := 0; g <= k; g++ {
			for b := 0; b <= k; b++ {
				fr := float64(r)
				fg := float64(g)
				fb := float64(b)
				if math.Abs(fr-fg) <= fv && math.Abs(fr-fb) <= fv && math.Abs(fg-fb) <= fv {
					t += 1
				}
			}
		}
	}
	return t
}
