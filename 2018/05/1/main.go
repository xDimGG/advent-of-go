package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	input, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(input)-1; i++ {
		if isOppCase(input[i], input[i+1]) {
			input = append(input[:i], input[i+2:]...)
			i -= 2
			if i < -1 {
				i = -1
			}
		}
	}

	fmt.Println("Result:", len(input))
}

func isOppCase(a, b byte) bool {
	if a >= 65 && a <= 90 {
		return a+32 == b
	}

	return b+32 == a
}
