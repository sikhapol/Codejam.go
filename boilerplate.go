package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func split(data []byte, atEOF bool) (advance int, token []byte, err error) {
	return bufio.ScanLines(data, atEOF)
}

func mapProb(chunk []byte) interface{} {
	return chunk
}

func solveProb(problem interface{}) string {
	return string(problem.([]byte))
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	numProbs, _ := strconv.Atoi(scanner.Text())

	log.Println("problem size:", numProbs)

	scanner.Split(split)
	for i := 1; scanner.Scan(); i++ {
		probBytes := scanner.Bytes()
		prob := mapProb(probBytes)
		fmt.Println("Case #"+strconv.Itoa(i), solveProb(prob))
	}
}
