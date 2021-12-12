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

const start = "start"
const end = "end"

type cave struct {
	id      string
	isSmall bool
}

func read() (data map[string][]cave) {
	lines := aoc.ReadAllInput()
	data = make(map[string][]cave)

	addCon := func(from, to string) {
		data[from] = append(data[from], cave{
			id:      to,
			isSmall: to != strings.ToUpper(to),
		})
	}

	for _, line := range lines {
		if line == "" {
			continue
		}
		connections := strings.Split(line, "-")
		con0 := connections[0]
		con1 := connections[1]
		if con0 == start || con1 == end {
			addCon(con0, con1)
		} else if con1 == start || con0 == end {
			addCon(con1, con0)
		} else {
			addCon(con0, con1)
			addCon(con1, con0)
		}
	}

	return data
}

func process(data map[string][]cave) int {
	visitedSmall := make(map[string]bool)
	c := count(data, visitedSmall, start)

	return c
}

func count(data map[string][]cave, visitedSmall map[string]bool, current string) int {
	sum := 0

	for _, c := range data[current] {
		if visitedSmall[c.id] {
			continue
		}
		if c.id == end {
			if len(visitedSmall) > 0 {
				sum++
			}
			continue
		}
		if c.isSmall {
			visitedSmall[c.id] = true
		}
		sum += count(data, visitedSmall, c.id)

		delete(visitedSmall, c.id)
	}

	return sum
}
