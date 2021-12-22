package main

import (
	"fmt"
	"strings"

	"github.com/AntonKosov/advent-of-code-2021/aoc"
)

func main() {
	player1, player2 := read()
	r := process(player1, player2)
	fmt.Printf("Answer: %v\n", r)
}

func read() (player1, player2 int) {
	lines := aoc.ReadAllInput()

	s1 := strings.Split(lines[0], " ")
	player1 = aoc.StrToInt(s1[len(s1)-1])

	s2 := strings.Split(lines[1], " ")
	player2 = aoc.StrToInt(s2[len(s2)-1])

	return player1, player2
}

type player struct {
	score    int
	position int
}

func process(p1, p2 int) int {
	cp := &player{position: p1}
	np := &player{position: p2}
	rolls := 0
	roll := func() int {
		s := 0
		for i := 0; i < 3; i++ {
			rolls++
			s += (rolls-1)%100 + 1
		}
		return s
	}
	for {
		s := roll()
		cp.position = (cp.position-1+s)%10 + 1
		cp.score += cp.position
		if cp.score >= 1000 {
			return rolls * np.score
		}
		cp, np = np, cp
	}
}
