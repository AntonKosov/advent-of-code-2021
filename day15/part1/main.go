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
		row := make([]int, len(line))
		for i, v := range line {
			row[i] = int(v - '0')
		}
		data = append(data, row)
	}

	return data
}

type pos struct {
	row, col int
}

func process(data [][]int) int {
	queue := []pos{{0, 0}}
	width := len(data[0])
	height := len(data)

	m := make([][]int, height)
	for i := range m {
		m[i] = make([]int, width)
	}

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		currentTotalLevel := m[p.row][p.col]
		checkPosition := func(row, col int) {
			level := data[row][col]
			previousTotalLevel := m[row][col]
			if previousTotalLevel == 0 || currentTotalLevel+level < previousTotalLevel {
				m[row][col] = currentTotalLevel + level
				queue = append(queue, pos{row, col})
			}
		}
		if p.row > 0 {
			checkPosition(p.row-1, p.col)
		}
		if p.row < height-1 {
			checkPosition(p.row+1, p.col)
		}
		if p.col > 0 {
			checkPosition(p.row, p.col-1)
		}
		if p.col < width-1 {
			checkPosition(p.row, p.col+1)
		}
	}

	return m[height-1][width-1]
}
