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

var rolls map[int]int // score -> number of universes

func init() {
	rolls = map[int]int{}
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			for k := 1; k <= 3; k++ {
				rolls[i+j+k]++
			}
		}
	}
}

type state struct {
	score    int
	position int
}

type universeState struct {
	players [2]state
}

func turn(unis map[universeState]int, playerIndex int) (newUnis map[universeState]int, wins int) {
	anotherPlayerIndex := 1 - playerIndex
	newUnis = make(map[universeState]int, len(unis))
	for uniState, uniCount := range unis {
		currentPlayerState := uniState.players[playerIndex]
		for rollScore, rollUni := range rolls {
			position := (currentPlayerState.position-1+rollScore)%10 + 1
			score := currentPlayerState.score + position
			universes := uniCount * rollUni
			if score >= 21 {
				wins += universes
				continue
			}
			newCurrentPlayerState := state{score: score, position: position}
			newPlayersState := [2]state{}
			newPlayersState[playerIndex] = newCurrentPlayerState
			newPlayersState[anotherPlayerIndex] = uniState.players[anotherPlayerIndex]
			newUni := universeState{players: newPlayersState}
			newUnis[newUni] += universes
		}
	}

	return newUnis, wins
}

func process(p1, p2 int) int {
	unis := map[universeState]int{{players: [2]state{
		0: {position: p1},
		1: {position: p2},
	}}: 1}
	wins := [2]int{}

	currentPlayer := 0
	for len(unis) > 0 {
		var w int
		unis, w = turn(unis, currentPlayer)
		wins[currentPlayer] += w

		currentPlayer = 1 - currentPlayer
	}

	max := wins[0]
	if wins[1] > max {
		max = wins[1]
	}

	return max
}
