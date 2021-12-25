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

// #################
// #01. 2 .3. 4. 56#
// ### 8#10#12#14###
//   # 7# 9#11#13#
//   #############
type burrow [15]int

const lastHallIndex = 6

type path struct {
	pointsBetween []int
	length        int
}

var paths []map[int]path

var energy map[int]int

func init() {
	paths = make([]map[int]path, 15)
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
		p = append(p, p[len(p)-1]-1)
		length++
		addPath(p, length)
		addPath(p[:len(p)-1], length-1)
	}
	// #################
	// #01. 2 .3. 4. 56#
	// ### 8#10#12#14###
	//   # 7# 9#11#13#
	//   #############
	addPathsToRoom([]int{0, 1, 8}, 3)
	addPathsToRoom([]int{0, 1, 2, 10}, 5)
	addPathsToRoom([]int{0, 1, 2, 3, 12}, 7)
	addPathsToRoom([]int{0, 1, 2, 3, 4, 14}, 9)
	addPathsToRoom([]int{1, 8}, 2)
	addPathsToRoom([]int{1, 2, 10}, 4)
	addPathsToRoom([]int{1, 2, 3, 12}, 6)
	addPathsToRoom([]int{1, 2, 3, 4, 14}, 8)
	addPathsToRoom([]int{2, 8}, 2)
	addPathsToRoom([]int{2, 10}, 2)
	addPathsToRoom([]int{2, 3, 12}, 4)
	addPathsToRoom([]int{2, 3, 4, 14}, 6)
	addPathsToRoom([]int{3, 2, 8}, 4)
	addPathsToRoom([]int{3, 10}, 2)
	addPathsToRoom([]int{3, 12}, 2)
	addPathsToRoom([]int{3, 4, 14}, 4)
	addPathsToRoom([]int{4, 3, 2, 8}, 6)
	addPathsToRoom([]int{4, 3, 10}, 4)
	addPathsToRoom([]int{4, 12}, 2)
	addPathsToRoom([]int{4, 14}, 2)
	addPathsToRoom([]int{5, 4, 3, 2, 8}, 8)
	addPathsToRoom([]int{5, 4, 3, 10}, 6)
	addPathsToRoom([]int{5, 4, 12}, 4)
	addPathsToRoom([]int{5, 14}, 2)
	addPathsToRoom([]int{6, 5, 4, 3, 2, 8}, 9)
	addPathsToRoom([]int{6, 5, 4, 3, 10}, 7)
	addPathsToRoom([]int{6, 5, 4, 12}, 5)
	addPathsToRoom([]int{6, 5, 14}, 3)

	energy = map[int]int{
		1: 1,
		2: 10,
		3: 100,
		4: 1000,
	}
}

func read() (data burrow) {
	lines := aoc.ReadAllInput()

	set := func(index int, a byte) {
		data[index] = int(a-byte('A')) + 1
	}
	for i := 0; i < 4; i++ {
		set(i*2+7, lines[3][i*2+3])
		set(i*2+8, lines[2][i*2+3])
	}

	return data
}

func process(data burrow) int {
	cache := map[burrow]int{{0, 0, 0, 0, 0, 0, 0, 1, 1, 2, 2, 3, 3, 4, 4}: 0}
	min, solved := minEnergy(data, cache)
	if !solved {
		panic("No solution")
	}
	return min
}

// #################
// #01. 2 .3. 4. 56#
// ### 8#10#12#14###
//   # 7# 9#11#13#
//   #############
func minEnergy(b burrow, cache map[burrow]int) (min int, solved bool) {
	if v, ok := cache[b]; ok {
		return v, true
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
	for animal := 1; animal <= 4; animal++ {
		bi, ti := animalHome(animal)
		var from int
		if b[ti] != 0 {
			if b[ti] != animal || b[bi] != animal {
				from = ti
			} else {
				continue
			}
		} else if b[ti] == 0 && b[bi] != 0 && b[bi] != animal {
			from = bi
		} else {
			continue
		}
		for to := 0; to <= lastHallIndex; to++ {
			if b[to] != 0 {
				continue
			}
			p := paths[from][to]
			if !isPathFree(b, p) {
				continue
			}
			movedAnimal := b[from]
			movingEnergy := energy[movedAnimal] * p.length
			if movingEnergy < minVariant {
				b[from], b[to] = b[to], b[from]
				if en, ok := minEnergy(b, cache); ok {
					solved = true
					minVariant = aoc.Min(minVariant, en+movingEnergy)
				}
				b[to], b[from] = b[from], b[to]
			}
		}
	}

	if !solved {
		return 0, false
	}

	min += minVariant
	cache[originalBurrow] = min

	return min, true
}

func animalHome(animal int) (bottomIndex, topIndex int) {
	bottomIndex = 5 + animal*2
	return bottomIndex, bottomIndex + 1
}

func isPathFree(b burrow, p path) bool {
	for _, i := range p.pointsBetween {
		if b[i] != 0 {
			return false
		}
	}
	return true
}

func moveToHome(b *burrow) int {
	sum := 0
	for from := 0; from <= lastHallIndex; from++ {
		animal := b[from]
		if animal == 0 {
			continue
		}

		bi, ti := animalHome(animal)
		if b[ti] != 0 || (b[bi] != 0 && b[bi] != animal) {
			continue
		}

		p := paths[from][bi]
		if !isPathFree(*b, p) {
			continue
		}

		if b[bi] == 0 {
			b[bi], b[from] = b[from], b[bi]
			sum += energy[animal] * p.length
			continue
		}

		b[ti], b[from] = b[from], b[ti]
		sum += energy[animal] * paths[from][ti].length
	}

	return sum
}
