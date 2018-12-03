package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	land := make(map[int32]int)

	for scanner.Scan() {
		var id, x, y, w, h int32
		fmt.Sscanf(scanner.Text(), "#%d @ %d,%d: %dx%d", &id, &x, &y, &w, &h)

		for hi := int32(0); hi < h; hi++ {
			for wi := int32(0); wi < w; wi++ {
				land[(y+hi)<<16|(x+wi)]++
			}
		}
	}

	total := 0

	for _, count := range land {
		if count >= 2 {
			total++
		}
	}

	fmt.Println("Count:", total)
}
