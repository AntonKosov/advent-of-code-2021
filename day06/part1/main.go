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

func read() (data []int) {
	lines := aoc.ReadAllInput()

	return aoc.StrToInts(lines[0], ",")
}

func process(data []int) int {
	timers := [9]int{}
	population := len(data)
	for _, timer := range data {
		timers[timer]++
	}

	for i := 0; i < 80; i++ {
		giveBirth := timers[0]
		copy(timers[:len(timers)-1], timers[1:])
		timers[6] += giveBirth
		timers[8] = giveBirth
		population += giveBirth
	}

	return population
}
