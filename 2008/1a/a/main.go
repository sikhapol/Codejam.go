package main

import (
	"fmt"
	"sort"
)

func main() {
	var t int
	fmt.Scan(&t)
	for ti := 0; ti < t; ti++ {
		var n int
		fmt.Scan(&n)
		x, y := make([]int, n), make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&x[i])
		}
		for i := 0; i < n; i++ {
			fmt.Scan(&y[i])
		}
		sort.Ints(x)
		sort.Ints(y)
		sum := 0
		for i := 0; i < n; i++ {
			sum += x[i] * y[n-i-1]
		}
		fmt.Printf("Case #%d: %v\n", ti+1, sum)
	}
}
