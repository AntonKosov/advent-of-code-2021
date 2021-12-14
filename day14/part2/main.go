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

type cacheItem struct {
	pair  string
	steps int
}

func process(data input) int {
	count := make(map[rune]int)
	cache := make(map[cacheItem]map[rune]int)
	for j := 0; j < len(data.pol)-1; j++ {
		pair := data.pol[j : j+2]
		r := countPairs(pair, 40, data.rules, cache)
		merge(count, r)
	}
	count[rune(data.pol[len(data.pol)-1])]++

	min, max := math.MaxInt, math.MinInt
	for _, c := range count {
		min = aoc.Min(min, c)
		max = aoc.Max(max, c)
	}

	return max - min
}

func merge(dst map[rune]int, src map[rune]int) {
	for r, c := range src {
		dst[r] += c
	}
}

func countPairs(pair string, stepsLeft int, rules map[string]rune,
	cache map[cacheItem]map[rune]int,
) map[rune]int {
	result := make(map[rune]int)
	countPair := func(p string) {
		if stepsLeft == 1 {
			result[rune(p[0])]++
			return
		}
		ci := cacheItem{pair: p, steps: stepsLeft}
		if cached, ok := cache[ci]; ok {
			merge(result, cached)
			return
		}
		m := countPairs(p, stepsLeft-1, rules, cache)
		cache[ci] = m
		merge(result, m)
	}

	m := rules[pair]
	countPair(string(pair[0]) + string(m))
	countPair(string(m) + string(pair[1]))

	return result
}
