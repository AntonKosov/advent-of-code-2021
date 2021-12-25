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
			b.state[i] = 1 + int(v-'A')
		}
	}
	// #################
	// #01. 2 .3. 4. 56#
	// ###10#14#18#22###
	//   # 9#13#17#21#
	//   # 8#12#16#20#
	//   # 7#11#15#19#
	//   #############
	setHome := func(homeIndex int) {
		bi, ti := roomIndexes(homeIndex)
		for i := bi; i <= ti; i++ {
			off := i - bi
			set(i, lines[5-off][homeIndex*2+1])
		}
	}
	set(0, lines[1][1])
	set(1, lines[1][2])
	set(2, lines[1][4])
	set(3, lines[1][6])
	set(4, lines[1][8])
	set(5, lines[1][10])
	set(6, lines[1][11])
	for i := 1; i <= 4; i++ {
		setHome(i)
	}

	return b
}

func TestLastMovement(t *testing.T) {
	testProcess(
		t,
		`
#############
#..........D#
###A#B#C#.###
  #A#B#C#D#
  #A#B#C#D#
  #A#B#C#D#
  #########`,
		3000,
	)
}

func TestADMovement(t *testing.T) {
	testProcess(
		t,
		`
#############
#.........AD#
###.#B#C#.###
  #A#B#C#D#
  #A#B#C#D#
  #A#B#C#D#
  #########`,
		8+3000,
	)
}

func TestDADMovement(t *testing.T) {
	testProcess(
		t,
		`
#############
#...D.....AD#
###.#B#C#.###
  #A#B#C#.#
  #A#B#C#D#
  #A#B#C#D#
  #########`,
		7000+8+3000,
	)
}

func TestADADMovement(t *testing.T) {
	testProcess(
		t,
		`
#############
#A..D.....AD#
###.#B#C#.###
  #.#B#C#.#
  #A#B#C#D#
  #A#B#C#D#
  #########`,
		4+7000+8+3000,
	)
}

func TestAADADMovement(t *testing.T) {
	testProcess(
		t,
		`
#############
#AA.D.....AD#
###.#B#C#.###
  #.#B#C#.#
  #.#B#C#D#
  #A#B#C#D#
  #########`,
		4+4+7000+8+3000,
	)
}

func TestFirstMovementOut(t *testing.T) {
	testProcess(
		t,
		`
#############
#AA.......AD#
###.#B#C#.###
  #.#B#C#.#
  #D#B#C#D#
  #A#B#C#D#
  #########`,
		4000+4+4+7000+8+3000,
	)
}
