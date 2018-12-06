package main

import (
	"bufio"
	"fmt"
	"os"
)

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

	gridSize := w * h
	grid := make([]int, gridSize)
	infinite := make([]bool, len(points))

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			offset := (y * w) + x

			closestDistance := gridSize
			closestTwice := false
			var closest *point

			for _, p := range points {
				distance := manhattanDistance(x, p.x, y, p.y)

				if distance < closestDistance {
					closestTwice = false
					closestDistance = distance
					closest = p

					if x == 0 || y == 0 || x == w-1 || y == h-1 {
						infinite[closest.id] = true
					}
				} else if distance == closestDistance {
					closestTwice = true
				}
			}

			if closestTwice {
				grid[offset] = -1
			} else {
				grid[offset] = closest.id
			}
		}
	}

	// for xy, id := range grid {
	// 	if id == -1 {
	// 		id = '.'
	// 	} else {
	// 		id += 'a'
	// 	}

	// 	if xy%w == 0 && xy != 0 {
	// 		fmt.Printf("\n")
	// 	}

	// 	fmt.Printf("%c", id)
	// }

	counts := make(map[int]int)

	for _, id := range grid {
		if id == -1 {
			continue
		}

		if !infinite[id] {
			counts[id]++
		}
	}

	highestCount := 0

	for _, count := range counts {
		if count > highestCount {
			highestCount = count
		}
	}

	fmt.Println("Result:", highestCount)
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
