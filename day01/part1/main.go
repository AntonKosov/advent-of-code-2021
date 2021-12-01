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
	sum := 0

	for i := 1; i < len(data); i++ {
		if data[i-1] < data[i] {
			sum++
		}
	}

	return sum
}
