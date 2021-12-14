package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/AntonKosov/advent-of-code-2021/aoc"
)

func main() {
	data := read()
	r := process(data)
	fmt.Printf("Answer: %v\n", r)
}

type input struct {
	pol   string
	rules map[string]rune
}

func read() (data input) {
	lines := aoc.ReadAllInput()
	data.pol = lines[0]
	data.rules = make(map[string]rune)

	for i := 2; i < len(lines); i++ {
		line := lines[i]
		if line == "" {
			continue
		}
		s := strings.Split(line, " -> ")
		data.rules[s[0]] = rune(s[1][0])
	}

	return data
}

func process(data input) int {
	count := make(map[rune]int)
	for j := 0; j < len(data.pol)-1; j++ {
		pair := data.pol[j : j+2]
		countPairs(pair, 10, data.rules, count)
	}
	count[rune(data.pol[len(data.pol)-1])]++

	min, max := math.MaxInt, math.MinInt
	for _, c := range count {
		min = aoc.Min(min, c)
		max = aoc.Max(max, c)
	}

	return max - min
}

func countPairs(pair string, stepsLeft int, rules map[string]rune, count map[rune]int) {
	if stepsLeft == 0 {
		count[rune(pair[0])]++
		return
	}

	m := rules[pair]
	firstPair := string(pair[0]) + string(m)
	countPairs(firstPair, stepsLeft-1, rules, count)
	secondPair := string(m) + string(pair[1])
	countPairs(secondPair, stepsLeft-1, rules, count)
}
