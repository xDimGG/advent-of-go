package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

func main() {
	input, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}

	values := bytes.Split(input, []byte{'\n'})

	fmt.Println("Duplicate:", firstDuplicate(values))
}

func firstDuplicate(values [][]byte) int {
	seen := make(map[int]struct{})
	acc := 0

	for {
		for _, v := range values {
			i, err := strconv.Atoi(string(v))
			if err != nil {
				panic(err)
			}

			acc += i

			if _, ok := seen[acc]; ok {
				return acc
			} else {
				seen[acc] = struct{}{}
			}
		}
	}
}
