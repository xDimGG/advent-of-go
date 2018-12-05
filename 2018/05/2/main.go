package main

import (
	"fmt"
	"io/ioutil"
)

const alphabetSize = 26

func main() {
	input, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}

	var results [alphabetSize]int

	for c := 0; c < alphabetSize; c++ {
		inputCopy := make([]byte, len(input))
		copy(inputCopy, input)

		for i := 0; i < len(inputCopy)-1; i++ {
			if int(inputCopy[i]) == c+'A' || int(inputCopy[i]) == c+'a' {
				inputCopy = append(inputCopy[:i], inputCopy[i+1:]...)
			} else if isOppCase(inputCopy[i], inputCopy[i+1]) {
				inputCopy = append(inputCopy[:i], inputCopy[i+2:]...)
			} else {
				continue
			}

			i -= 2
			if i < -1 {
				i = -1
			}
		}

		results[c] = len(inputCopy)
	}

	lowestResult := results[0]

	for _, result := range results {
		if result < lowestResult {
			lowestResult = result
		}
	}

	fmt.Println("Result:", lowestResult)
}

func isOppCase(a, b byte) bool {
	if a >= 'A' && a <= 'Z' {
		return a+32 == b
	}

	return b+32 == a
}
