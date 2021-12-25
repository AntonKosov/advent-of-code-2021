package main

import (
	"fmt"
	"math"

	"github.com/AntonKosov/advent-of-code-2021/aoc"
)

func main() {
	data := read()
	r := process(data)
	fmt.Printf("Answer: %v\n", r)
}

const cells = 23

// #################
// #01. 2 .3. 4. 56#
// ###10#14#18#22###
//   # 9#13#17#21#
//   # 8#12#16#20#
//   # 7#11#15#19#
//   #############
type burrow struct {
	state [cells]int
}

func (b *burrow) isPathFree(p path) bool {
	for _, i := range p.pointsBetween {
		if b.state[i] != 0 {
			return false
		}
	}
	return true
}

func (b *burrow) hall(i int) (animal int, found bool) {
	animal = b.state[i]
	return animal, animal != 0
}

func (b *burrow) isFree(hallIndex int) bool {
	return b.state[hallIndex] == 0
}

func (b *burrow) homeAvailable(animal int) (index int, found bool) {
	bi, ti := roomIndexes(animal)
	for i := bi; i <= ti; i++ {
		a := b.state[i]
		if a == 0 {
			return i, true
		}
		if a != animal {
			return -1, false
		}
	}
	return -1, false
}

func (b *burrow) swap(from, to int) {
	b.state[from], b.state[to] = b.state[to], b.state[from]
}

func (b *burrow) moveToHome(hallIndex int) (spentEnergy int) {
	animal, found := b.hall(hallIndex)
	if !found {
		return 0
	}

	homeCell, found := b.homeAvailable(animal)
	if !found {
		return 0
	}

	p := paths[hallIndex][homeCell]
	if !b.isPathFree(p) {
		return 0
	}

	b.state[homeCell], b.state[hallIndex] = b.state[hallIndex], b.state[homeCell]

	return energy[animal] * p.length
}

func (b *burrow) topAnimalInHomeToMoveOut(home int) (index int, animal int, found bool) {
	bi, ti := roomIndexes(home)
	for i := ti; i >= bi; i-- {
		a := b.state[i]
		if a == 0 {
			continue
		}
		if a != home {
			return i, a, true
		}
		for j := i - 1; j >= bi; j-- {
			if b.state[j] != home {
				return i, a, true
			}
		}
		return -1, -1, false
	}
	return -1, -1, false
}

func roomIndexes(animal int) (bottomIndex, topIndex int) {
	bottomIndex = 3 + animal*4
	return bottomIndex, bottomIndex + 3
}

const lastHallIndex = 6

type path struct {
	pointsBetween []int
	length        int
}

var paths []map[int]path

var energy map[int]int

func init() {
	paths = make([]map[int]path, cells)
	for i := range paths {
		paths[i] = make(map[int]path, 8)
	}
	addPath := func(p []int, length int) {
		from := p[0]
		to := p[len(p)-1]
		p = p[1 : len(p)-1]
		paths[from][to] = path{
			pointsBetween: p,
			length:        length,
		}
		rev := make([]int, 0, len(p))
		for i := len(p) - 1; i >= 0; i-- {
			rev = append(rev, p[i])
		}
		paths[to][from] = path{
			pointsBetween: rev,
			length:        length,
		}
	}

	addPathsToRoom := func(p []int, length int) {
		for i := 0; i < 4; i++ {
			addPath(p, length)
			p = append(p, p[len(p)-1]-1)
			length++
		}
	}
	// #################
	// #01. 2 .3. 4. 56#
	// ###10#14#18#22###
	//   # 9#13#17#21#
	//   # 8#12#16#20#
	//   # 7#11#15#19#
	//   #############
	addPathsToRoom([]int{0, 1, 10}, 3)
	addPathsToRoom([]int{0, 1, 2, 14}, 5)
	addPathsToRoom([]int{0, 1, 2, 3, 18}, 7)
	addPathsToRoom([]int{0, 1, 2, 3, 4, 22}, 9)
	addPathsToRoom([]int{1, 10}, 2)
	addPathsToRoom([]int{1, 2, 14}, 4)
	addPathsToRoom([]int{1, 2, 3, 18}, 6)
	addPathsToRoom([]int{1, 2, 3, 4, 22}, 8)
	addPathsToRoom([]int{2, 10}, 2)
	addPathsToRoom([]int{2, 14}, 2)
	addPathsToRoom([]int{2, 3, 18}, 4)
	addPathsToRoom([]int{2, 3, 4, 22}, 6)
	addPathsToRoom([]int{3, 2, 10}, 4)
	addPathsToRoom([]int{3, 14}, 2)
	addPathsToRoom([]int{3, 18}, 2)
	addPathsToRoom([]int{3, 4, 22}, 4)
	addPathsToRoom([]int{4, 3, 2, 10}, 6)
	addPathsToRoom([]int{4, 3, 14}, 4)
	addPathsToRoom([]int{4, 18}, 2)
	addPathsToRoom([]int{4, 22}, 2)
	addPathsToRoom([]int{5, 4, 3, 2, 10}, 8)
	addPathsToRoom([]int{5, 4, 3, 14}, 6)
	addPathsToRoom([]int{5, 4, 18}, 4)
	addPathsToRoom([]int{5, 22}, 2)
	addPathsToRoom([]int{6, 5, 4, 3, 2, 10}, 9)
	addPathsToRoom([]int{6, 5, 4, 3, 14}, 7)
	addPathsToRoom([]int{6, 5, 4, 18}, 5)
	addPathsToRoom([]int{6, 5, 22}, 3)

	energy = map[int]int{
		1: 1,
		2: 10,
		3: 100,
		4: 1000,
	}
}

func read() (data burrow) {
	lines := aoc.ReadAllInput()

	// #################
	// #01. 2 .3. 4. 56#
	// ###10#14#18#22###
	//   # 9#13#17#21#
	//   # 8#12#16#20#
	//   # 7#11#15#19#
	//   #############
	set := func(index int, a byte) {
		data.state[index] = int((a - byte('A') + 1))
	}
	for i := 0; i < 4; i++ {
		set(i*4+7, lines[5][i*2+3])
		set(i*4+8, lines[4][i*2+3])
		set(i*4+9, lines[3][i*2+3])
		set(i*4+10, lines[2][i*2+3])
	}

	return data
}

func process(data burrow) int {
	cache := map[burrow]int{
		{
			state: [23]int{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 2, 2, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4},
		}: 0,
	}
	deadEnd := map[burrow]bool{}

	min, solved := minEnergy(data, cache, deadEnd)
	if !solved {
		panic("No solution")
	}

	return min
}

// #################
// #01. 2 .3. 4. 56#
// ###10#14#18#22###
//   # 9#13#17#21#
//   # 8#12#16#20#
//   # 7#11#15#19#
//   #############
func minEnergy(b burrow, cache map[burrow]int, deadEnd map[burrow]bool) (min int, solved bool) {
	if v, ok := cache[b]; ok {
		return v, true
	}
	if deadEnd[b] {
		return 0, false
	}

	originalBurrow := b
	min = 0
	for {
		se := moveToHome(&b)
		if se == 0 {
			break
		}
		min += se
	}

	if v, ok := cache[b]; ok {
		min += v
		cache[originalBurrow] = min
		return min, true
	}

	minVariant := math.MaxInt
	for home := 1; home <= 4; home++ {
		from, animal, found := b.topAnimalInHomeToMoveOut(home)
		if !found {
			continue
		}
		for to := 0; to <= lastHallIndex; to++ {
			if !b.isFree(to) {
				continue
			}
			p := paths[from][to]
			movingEnergy := energy[animal] * p.length
			if movingEnergy >= minVariant {
				continue
			}
			if !b.isPathFree(p) {
				continue
			}
			b.swap(from, to)
			if en, ok := minEnergy(b, cache, deadEnd); ok {
				solved = true
				minVariant = aoc.Min(minVariant, en+movingEnergy)
			}
			b.swap(from, to)
		}
	}

	if !solved {
		deadEnd[originalBurrow] = true
		return 0, false
	}

	min += minVariant
	cache[originalBurrow] = min

	return min, true
}

func moveToHome(b *burrow) int {
	sum := 0
	for i := 0; i <= lastHallIndex; i++ {
		sum += b.moveToHome(i)
	}

	return sum
}
