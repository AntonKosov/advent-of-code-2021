package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2021/aoc"
)

func main() {
	data := read()
	r := process(data)
	fmt.Printf("Answer: %v\n", r)
}

func read() (data map[aoc.Vector2]int) {
	lines := aoc.ReadAllInput()
	data = make(map[aoc.Vector2]int, 100)

	for y := 0; y < 10; y++ {
		line := lines[y]
		for x, v := range line {
			data[aoc.NewVector2(x, y)] = int(v - '0')
		}
	}

	return data
}

func process(data map[aoc.Vector2]int) int {
	for i := 0; ; i++ {
		queue := make(map[aoc.Vector2]bool)
		flashed := make(map[aoc.Vector2]bool)
		charge := func(c aoc.Vector2) {
			if flashed[c] {
				return
			}
			data[c]++
			if data[c] == 10 {
				queue[c] = true
			}
		}

		for c := range data {
			charge(c)
		}

		flashes := 0
		for len(queue) > 0 {
			q := queue
			queue = make(map[aoc.Vector2]bool)
			for c := range q {
				flashes++
				flashed[c] = true
				data[c] = 0
				for _, a := range c.Adjacent() {
					if _, ok := data[a]; ok {
						charge(a)
					}
				}
			}
		}
		if flashes == 100 {
			return i + 1
		}
	}
}
