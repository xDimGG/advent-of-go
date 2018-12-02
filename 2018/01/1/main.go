package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(input)
	acc := 0

	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}

		acc += i
	}

	fmt.Println("Sum:", acc)
}
