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

var openBrackets map[rune]bool
var closeBrackets map[rune]rune
var points map[rune]int

func init() {
	openBrackets = map[rune]bool{'(': true, '{': true, '[': true, '<': true}
	closeBrackets = map[rune]rune{'(': ')', '{': '}', '[': ']', '<': '>'}
	points = map[rune]int{')': 3, ']': 57, '}': 1197, '>': 25137}
}

func process(data []string) int {
	score := 0

	for _, line := range data {
		var stack []rune
		for _, b := range line {
			if openBrackets[b] {
				stack = append(stack, b)
				continue
			}
			expectedCloseBracket := closeBrackets[stack[len(stack)-1]]
			if b != expectedCloseBracket {
				score += points[b]
				break
			}
			stack = stack[:len(stack)-1]
		}
	}

	return score
}
