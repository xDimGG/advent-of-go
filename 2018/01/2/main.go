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

	i, err := firstDuplicate(bytes.Split(input, []byte{'\n'}))
	if err != nil {
		panic(err)
	}

	fmt.Println("Duplicate:", i)
}

func firstDuplicate(values [][]byte) (i int, err error) {
	seen := make(map[int]struct{})
	acc := 0

	for {
		for _, v := range values {
			i, err = strconv.Atoi(string(v))
			if err != nil {
				return
			}

			acc += i

			if _, ok := seen[acc]; ok {
				return acc, nil
			}

			seen[acc] = struct{}{}
		}
	}
}
