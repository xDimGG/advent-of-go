package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(input)
	all := make(map[int]int)

	for scanner.Scan() {
		chars := make(map[byte]int)
		for _, char := range scanner.Bytes() {
			chars[char]++
		}

		counts := make(map[int]struct{})
		for _, count := range chars {
			counts[count] = struct{}{}
		}

		for count := range counts {
			if count <= 1 {
				continue
			}

			all[count]++
		}
	}

	product := 1
	for _, num := range all {
		product *= num
	}

	fmt.Println("Product:", product)
}
