package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2021/aoc"
)

type board struct {
	numbersLeft map[int]aoc.Vector2
	columns     [5]int
	rows        [5]int
}

func (b *board) isWinner(number int) (won bool) {
	lp, ok := b.numbersLeft[number]
	if !ok {
		return false
	}

	b.columns[lp.X]++
	b.rows[lp.Y]++
	delete(b.numbersLeft, number)

	return b.columns[lp.X] == 5 || b.rows[lp.Y] == 5
}

func (b *board) sumUnmarked() int {
	s := 0
	for v := range b.numbersLeft {
		s += v
	}

	return s
}

type game struct {
	boards  []*board
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
			numbersLeft: make(map[int]aoc.Vector2, 25),
		}
		for r := 0; r < 5; r++ {
			row := aoc.StrToInts(lines[bs+r], " ")
			for c, v := range row {
				p := aoc.NewVector2(c, r)
				b.numbersLeft[v] = p
			}
		}

		data.boards = append(data.boards, &b)
	}

	return data
}

func process(data game) int {
	leftBoards := make(map[int]*board, len(data.boards))
	for i, b := range data.boards {
		leftBoards[i] = b
	}

	for _, n := range data.numbers {
		var lastBoard *board
		for i, b := range leftBoards {
			if b.isWinner(n) {
				lastBoard = b
				delete(leftBoards, i)
			}
		}
		if len(leftBoards) == 0 {
			return lastBoard.sumUnmarked() * n
		}
	}

	panic("Oops")
}
