package main

import (
	"fmt"
	"strings"

	"github.com/AntonKosov/advent-of-code-2021/aoc"
	// "github.com/AntonKosov/advent-of-code-2021/day08/part2/bruteforce"
)

type entry struct {
	digits []string
	value  []string
}

func main() {
	data := read()
	r := process(data)
	fmt.Printf("Answer: %v\n", r)
}

func read() (data []entry) {
	lines := aoc.ReadAllInput()

	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " | ")
		data = append(data, entry{
			digits: strings.Split(parts[0], " "),
			value:  strings.Split(parts[1], " "),
		})
	}

	return data
}

func process(data []entry) int {
	sum := 0
	for _, d := range data {
		// sum += bruteforce.Process(d.digits, d.value)
		sum += calc(d.digits, d.value)
	}
	return sum
}

func calc(digits []string, value []string) int {
	mask := bitMask(digits)
	v := 0
	for _, digit := range value {
		var d byte
		for _, segment := range digit {
			d |= mask[segment]
		}
		v = v*10 + int(imageValues[d])
	}
	return v
}

var imageValues map[byte]byte

func init() {
	//  aaaa
	// b    c
	// b    c
	//  dddd
	// e    f
	// e    f
	//  gggg
	imageValues = map[byte]byte{
		//abcdefg
		0b1110111: 0,
		0b0010010: 1,
		0b1011101: 2,
		0b1011011: 3,
		0b0111010: 4,
		0b1101011: 5,
		0b1101111: 6,
		0b1010010: 7,
		0b1111111: 8,
		0b1111011: 9,
	}
}

func bitMask(digits []string) map[rune]byte {
	segments := make(map[rune]int, 10)
	var one []rune
	var four []rune
	for _, digit := range digits {
		switch len(digit) {
		case 2:
			one = []rune(digit)
		case 4:
			four = []rune(digit)
		}
		for _, segment := range digit {
			segments[segment]++
		}
	}

	count := make(map[int][]rune)
	for r, c := range segments {
		count[c] = append(count[c], r)
	}

	mask := make(map[rune]byte, 10)

	b := count[6][0]
	e := count[4][0]
	f := count[9][0]
	c := another(one, f)
	a := findExtra(count[8], one...)
	d := findExtra(four, b, c, f)
	g := another(count[7], d)

	//          abcdefg
	mask[a] = 0b1000000
	mask[b] = 0b0100000
	mask[c] = 0b0010000
	mask[d] = 0b0001000
	mask[e] = 0b0000100
	mask[f] = 0b0000010
	mask[g] = 0b0000001

	return mask
}

func another(l []rune, ex rune) rune {
	if l[0] == ex {
		return l[1]
	}
	return l[0]
}

func findExtra(extra []rune, exclude ...rune) rune {
next:
	for _, i := range extra {
		for _, j := range exclude {
			if i == j {
				continue next
			}
		}
		return i
	}

	panic("Extra rune not found")
}
