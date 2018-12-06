package main

import (
	"bufio"
	"fmt"
	"os"
)

const maxDist = 10000

type point struct{ x, y, id int }

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)

	var w, h int
	var points []*point

	for id := 0; scanner.Scan(); id++ {
		p := &point{id: id}
		fmt.Sscanf(scanner.Text(), "%d, %d", &p.x, &p.y)
		points = append(points, p)

		if p.x > w {
			w = p.x
		}

		if p.y > h {
			h = p.y
		}
	}

	w++
	h++

	area := 0

	for y := 0; y < h; y++ {
	loop:
		for x := 0; x < w; x++ {
			sum := 0

			for _, p := range points {
				sum += manhattanDistance(x, p.x, y, p.y)

				if sum >= maxDist {
					continue loop
				}
			}

			area++
		}
	}

	fmt.Println("Result:", area)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func manhattanDistance(x1, y1, x2, y2 int) int {
	return abs(x1-y1) + abs(x2-y2)
}
