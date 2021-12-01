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

func read() []int {
	var floor []int
	lines := aoc.ReadAllInput()
	for _, line := range lines {
		if line != "" {
			floor = append(floor, aoc.StrToInt(line))
		}
	}

	return floor
}

func process(data []int) int {
	sums := make([]int, len(data)-2)
	for i := 0; i < len(sums); i++ {
		sums[i] = data[i] + data[i+1] + data[i+2]
	}

	sum := 0

	for i := 1; i < len(sums); i++ {
		if sums[i-1] < sums[i] {
			sum++
		}
	}

	return sum
}
