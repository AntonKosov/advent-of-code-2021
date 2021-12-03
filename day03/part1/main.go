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

func read() (data []string) {
	lines := aoc.ReadAllInput()

	return lines[:len(lines)-1]
}

func process(data []string) int {
	gamma := 0
	epsilon := 0
	for bit := 0; bit < len(data[0]); bit++ {
		ones := 0
		for _, l := range data {
			if l[bit] == '1' {
				ones++
			}
		}
		if ones > len(data)/2 {
			gamma = (gamma << 1) | 1
			epsilon <<= 1
		} else {
			gamma <<= 1
			epsilon = (epsilon << 1) | 1
		}
	}

	return gamma * epsilon
}
