package main

import (
	"fmt"
)

func main() {
	var l, d, n int
	fmt.Scan(&l, &d, &n)
	dct := make([]string, d) // dictionary
	for di := 0; di < d; di++ {
		fmt.Scanln(&dct[di])
	}
	for ni := 0; ni < n; ni++ {
		var pt string // pattern
		fmt.Scan(&pt)
		paren := false
		pos := 0
		bitmap := make([][26]bool, l)
		for _, c := range pt {
			if paren {
				if c == ')' {
					paren = false
					pos++
				} else {
					bitmap[pos][c-'a'] = true
					// fmt.Printf("[%d|%d]", pos, c)
				}
			} else {
				if c == '(' {
					paren = true
				} else {
					bitmap[pos][c-'a'] = true
					// fmt.Printf("<%d|%d>", pos, c)
					pos++
				}
			}
		}
		count := 0
		for _, s := range dct {
			match := true
			for i, c := range s {
				match = match && bitmap[i][c-'a']
				if !match {
					break
				}
			}
			if match {
				count++
			}
		}
		fmt.Printf("Case #%d: %d\n", ni+1, count)
	}
}
