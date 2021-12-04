package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2021/aoc"
)

type board struct {
	positions map[int]aoc.Vector2
	values    map[aoc.Vector2]int
}

func (b board) isWinner(processed map[int]bool, lastNumber int) (won bool) {
	lp, ok := b.positions[lastNumber]
	if !ok {
		return false
	}

	v := true
	h := true
	for i := 0; i < 5; i++ {
		hv := b.values[aoc.NewVector2(i, lp.Y)]
		h = h && processed[hv]
		vv := b.values[aoc.NewVector2(lp.X, i)]
		v = v && processed[vv]
		if !h && !v {
			return false
		}
	}

	return true
}

func (b board) sumUnmarked(processed map[int]bool) int {
	s := 0
	for v := range b.positions {
		if !processed[v] {
			s += v
		}
	}
	return s
}

type game struct {
	boards  []board
	numbers []int
}

func main() {
	data := read()
	r := process(data)
	fmt.Printf("Answer: %v\n", r)
}

func read() (data game) {
	lines := aoc.ReadAllInput()

	data.numbers = aoc.StrToInts(lines[0], ",")
	for bs := 2; bs < len(lines); bs += 6 {
		b := board{
			positions: make(map[int]aoc.Vector2, 25),
			values:    make(map[aoc.Vector2]int, 25),
		}
		for r := 0; r < 5; r++ {
			row := aoc.StrToInts(lines[bs+r], " ")
			for c, v := range row {
				p := aoc.NewVector2(c, r)
				b.positions[v] = p
				b.values[p] = v
			}
		}

		data.boards = append(data.boards, b)
	}

	return data
}

func process(data game) int {
	processed := make(map[int]bool, len(data.numbers))
	for _, n := range data.numbers {
		processed[n] = true
		for _, b := range data.boards {
			if b.isWinner(processed, n) {
				return b.sumUnmarked(processed) * n
			}
		}
	}
	panic("No winners")
}
