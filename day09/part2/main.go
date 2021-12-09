package main

import (
	"fmt"
	"sort"

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
	var basinSizes []int

	for ri, row := range data {
		for ci := range row {
			basinSizes = append(basinSizes, findSize(data, ri, ci))
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(basinSizes)))

	return basinSizes[0] * basinSizes[1] * basinSizes[2]
}

func findSize(data [][]int, row, column int) int {
	if data[row][column] == 9 {
		return 0
	}

	data[row][column] = 9
	sum := 1

	if row > 0 {
		sum += findSize(data, row-1, column)
	}
	if row < len(data)-1 {
		sum += findSize(data, row+1, column)
	}
	if column > 0 {
		sum += findSize(data, row, column-1)
	}
	if column < len(data[0])-1 {
		sum += findSize(data, row, column+1)
	}

	return sum
}
