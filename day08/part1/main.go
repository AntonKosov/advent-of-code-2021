package main

import (
	"fmt"
	"strings"

	"github.com/AntonKosov/advent-of-code-2021/aoc"
)

func main() {
	data := read()
	r := process(data)
	fmt.Printf("Answer: %v\n", r)
}

func read() (data []string) {
	lines := aoc.ReadAllInput()

	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " | ")
		data = append(data, strings.Split(parts[1], " ")...)
	}

	return data
}

func process(data []string) int {
	sum := 0
	for _, d := range data {
		l := len(d)
		if l == 2 || l == 4 || l == 3 || l == 7 {
			sum++
		}
	}
	return sum
}
