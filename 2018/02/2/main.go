package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

func main() {
	input, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}

	ids := bytes.Split(input, []byte{'\n'})
	finalMatchCount := 0
	var finalA, finalB []byte

	for i, a := range ids {
		for _, b := range ids[i:] {
			if bytes.Equal(a, b) {
				continue
			}

			if m := matchCount(a, b); m > finalMatchCount {
				finalMatchCount = m
				finalA = a
				finalB = b
			}
		}
	}

	if finalMatchCount == 0 {
		fmt.Println("No match found.")
	} else {
		str := ""

		for i, char := range finalA {
			if char == finalB[i] {
				str += string(char)
			}
		}

		fmt.Println("Match:", str)
	}
}

func matchCount(a, b []byte) (count int) {
	for i, char := range a {
		if char == b[i] {
			count++
		}
	}

	return
}
