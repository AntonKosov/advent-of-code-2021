package main

import (
	"strings"
	"testing"
)

func testProcess(t *testing.T, m string, expected int) {
	b := parse(m)
	actual := process(b)
	if actual != expected {
		t.Fatalf("Expected: %v, actual: %v", expected, actual)
	}
}

func parse(t string) burrow {
	lines := strings.Split(t, "\n")[1:]
	var b burrow
	set := func(i int, v byte) {
		if v != byte('.') {
			b[i] = 1 + int(v-'A')
		}
	}
	set(0, lines[1][1])
	set(1, lines[1][2])
	set(2, lines[1][4])
	set(3, lines[1][6])
	set(4, lines[1][8])
	set(5, lines[1][10])
	set(6, lines[1][11])
	set(7, lines[3][3])
	set(8, lines[2][3])
	set(9, lines[3][5])
	set(10, lines[2][5])
	set(11, lines[3][7])
	set(12, lines[2][7])
	set(13, lines[3][9])
	set(14, lines[2][9])

	return b
}

func TestLastMovement(t *testing.T) {
	testProcess(
		t,
		`
#############
#.........A.#
###.#B#C#D###
  #A#B#C#D#
  #########`,
		8,
	)
}

func TestThreeMovement(t *testing.T) {
	testProcess(
		t,
		`
#############
#.....D.D.A.#
###.#B#C#.###
  #A#B#C#.#
  #########`,
		7000+8,
	)
}

func TestGiveWayToBothDMovement(t *testing.T) {
	testProcess(
		t,
		`
#############
#.....D.....#
###.#B#C#D###
  #A#B#C#A#
  #########
`,
		2003+7000+8,
	)
}

func TestBDADDA(t *testing.T) {
	testProcess(
		t,
		`
#############
#.....D.....#
###B#.#C#D###
  #A#B#C#A#
  #########
`,
		40+2003+7000+8,
	)
}

func TestDBBDADDA(t *testing.T) {
	testProcess(
		t,
		`
#############
#...B.......#
###B#.#C#D###
  #A#D#C#A#
  #########
`,
		3000+30+40+2003+7000+8,
	)
}

func TestFullExample(t *testing.T) {
	testProcess(
		t,
		`
#############
#...........#
###B#C#B#D###
  #A#D#C#A#
  #########
`,
		12521,
	)
}
