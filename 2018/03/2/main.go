package main

import (
	"bufio"
	"fmt"
	"os"
)

type claim struct {
	id, x, y, w, h int
}

func (c *claim) touches(o *claim) bool {
	return c.x < o.x+o.w && c.x+c.w > o.x && c.y < o.y+o.h && c.y+c.h > o.y
}

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	claims := make([]*claim, 0)

	for scanner.Scan() {
		var id, x, y, w, h int
		fmt.Sscanf(scanner.Text(), "#%d @ %d,%d: %dx%d", &id, &x, &y, &w, &h)
		claims = append(claims, &claim{id, x, y, w, h})
	}

loop:
	for _, a := range claims {
		for _, b := range claims {
			if a.id == b.id {
				continue
			}

			if a.touches(b) {
				continue loop
			}
		}

		fmt.Println("Untouched:", a.id)
	}
}
