package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func split(data []byte, atEOF bool) (advance int, token []byte, err error) {
	return bufio.ScanLines(data, atEOF)
}

func mapProb(chunk []byte) interface{} {
	return chunk
}

func solveProb(problem interface{}) string {
	shynesses := strings.Split(string(problem.([]byte)), " ")[1]
	var numRise, numInvite int
	for i, c := range shynesses {
		inv := int(math.Max(0, float64(i-numRise)))
		numInvite += inv
		numRise += inv + int(c-'0')
	}
	return strconv.Itoa(numInvite)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	numProbs, _ := strconv.Atoi(scanner.Text())

	log.Println("problem size:", numProbs)

	scanner.Split(split)
	var i int
	for i = 1; scanner.Scan(); i++ {
		probBytes := scanner.Bytes()
		prob := mapProb(probBytes)
		fmt.Println("Case #"+strconv.Itoa(i), solveProb(prob))
	}

	if i-1 != numProbs {
		log.Println("number of problems and solutions not match")
	}
}
