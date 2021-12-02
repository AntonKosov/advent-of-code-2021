package main

import (
	"fmt"
	"strings"

	"github.com/AntonKosov/advent-of-code-2021/aoc"
)

type command struct {
	dir   string
	units int
}

func main() {
	data := read()
	r := process(data)
	fmt.Printf("Answer: %v\n", r)
}

func read() (data []command) {
	lines := aoc.ReadAllInput()

	for _, line := range lines {
		if line == "" {
			continue
		}
		sp := strings.Split(line, " ")
		data = append(data, command{dir: sp[0], units: aoc.StrToInt(sp[1])})
	}

	return data
}

func process(data []command) int {
	distance := 0
	depth := 0
	aim := 0

	for _, c := range data {
		switch c.dir {
		case "forward":
			distance += c.units
			depth += aim * c.units
		case "down":
			aim += c.units
		case "up":
			aim -= c.units
		default:
			panic("Unknown command")
		}
	}

	return distance * depth
}
