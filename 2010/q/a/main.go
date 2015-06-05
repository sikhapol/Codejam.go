package main

import (
	"fmt"
)

func main() {
	n := 0 // number of test cases
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var c int // credit
		fmt.Scan(&c)
		var ni int // number of items
		fmt.Scan(&ni)
		items := make([]int, ni) // array of items
		for j := 0; j < ni; j++ {
			fmt.Scan(&items[j])
		}
		var i1, i2 int
		for j, found := 0, false; j < ni-1 && !found; j++ {
			for k := j + 1; k < ni; k++ {
				if items[j]+items[k] == c {
					i1, i2 = j, k
					found = true
					break
				}
			}
		}
		fmt.Printf("Case #%d: %v %v\n", i+1, i1+1, i2+1)
	}

}
