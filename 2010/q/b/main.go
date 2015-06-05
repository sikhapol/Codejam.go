package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	n := 0
	fmt.Scan(&n)
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < n && scanner.Scan(); i++ {
		line := scanner.Text()
		tokens := strings.Split(line, " ")
		fmt.Printf("Case #%d:", i+1)
		for j := len(tokens) - 1; j >= 0; j-- {
			fmt.Print(" ", tokens[j])
		}
		fmt.Print("\n")
	}
}
