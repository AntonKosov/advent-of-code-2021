package main

import (
	"fmt"
	"math"

	"github.com/AntonKosov/advent-of-code-2021/aoc"
)

func main() {
	data := read()
	r := process(data)
	fmt.Printf("Answer: %v\n", r)
}

func read() (data []int) {
	lines := aoc.ReadAllInput()

	return aoc.StrToInts(lines[0], ",")
}

func process(data []int) int {
	minFuel := math.MaxInt
	min, max := aoc.MinMax(data)
	for i := min; i <= max; i++ {
		fuel := 0
		for _, j := range data {
			distance := aoc.Abs(i - j)
			fuel += distance
		}
		minFuel = aoc.Min(minFuel, fuel)
	}

	return minFuel
}
