package main

import (
	"fmt"
	"sort"

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

var openBrackets map[rune]bool
var closeBrackets map[rune]rune
var points map[rune]int

func init() {
	openBrackets = map[rune]bool{'(': true, '{': true, '[': true, '<': true}
	closeBrackets = map[rune]rune{'(': ')', '{': '}', '[': ']', '<': '>'}
	points = map[rune]int{')': 1, ']': 2, '}': 3, '>': 4}
}

func process(data []string) int {
	var scores []int

nextLine:
	for _, line := range data {
		var stack []rune
		for _, b := range line {
			if openBrackets[b] {
				stack = append(stack, b)
				continue
			}
			expectedCloseBracket := closeBrackets[stack[len(stack)-1]]
			if b != expectedCloseBracket {
				continue nextLine
			}
			stack = stack[:len(stack)-1]
		}
		score := 0
		for len(stack) > 0 {
			ob := stack[len(stack)-1]
			cb := closeBrackets[ob]
			score = score*5 + points[cb]
			stack = stack[:len(stack)-1]
		}
		scores = append(scores, score)
	}

	sort.Ints(scores)

	return scores[len(scores)/2]
}
