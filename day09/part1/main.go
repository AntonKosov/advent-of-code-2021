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

func read() (data [][]int) {
	lines := aoc.ReadAllInput()

	for _, line := range lines {
		if line == "" {
			continue
		}
		row := make([]int, 0, len(line))
		for _, v := range line {
			row = append(row, int(v-'0'))
		}

		data = append(data, row)
	}

	return data
}

func process(data [][]int) int {
	sum := 0
	w := len(data[0])
	h := len(data)

	for ri, row := range data {
		for ci, v := range row {
			if ri > 0 && v >= data[ri-1][ci] {
				continue
			}
			if ri < h-1 && v >= data[ri+1][ci] {
				continue
			}
			if ci > 0 && v >= data[ri][ci-1] {
				continue
			}
			if ci < w-1 && v >= data[ri][ci+1] {
				continue
			}
			sum += 1 + v
		}
	}

	return sum
}
