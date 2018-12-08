package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

type node struct {
	children []*node
	metadata []byte
}

func main() {
	input, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}

	data := make([]byte, 0)

	for _, raw := range bytes.Split(input, []byte{' '}) {
		num, err := strconv.Atoi(string(raw))
		if err != nil {
			panic(err)
		}

		data = append(data, byte(num))
	}

	n, _ := read(data)
	fmt.Println("Result:", sum(n))
}

func read(data []byte) (n *node, size int) {
	children := int(data[0])
	metadata := int(data[1])
	n = &node{
		children: make([]*node, children),
		metadata: make([]byte, metadata),
	}

	data = data[2:]
	size += 2

	for i := 0; i < children; i++ {
		child, childSize := read(data)
		n.children[i] = child

		data = data[childSize:]
		size += childSize
	}

	n.metadata = data[:metadata]
	size += metadata

	return
}

func sum(n *node) (total int) {
	for _, c := range n.children {
		total += sum(c)
	}

	for _, d := range n.metadata {
		total += int(d)
	}

	return total
}
