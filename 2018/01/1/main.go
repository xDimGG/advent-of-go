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

	values := bufio.NewScanner(input)
	acc := 0

	for values.Scan() {
		i, err := strconv.Atoi(values.Text())
		if err != nil {
			panic(err)
		}

		acc += i
	}

	fmt.Println("Sum:", acc)
}
