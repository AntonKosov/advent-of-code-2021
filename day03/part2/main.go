package main

import (
	"fmt"
	"strconv"

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

type filterRule func(ones, zeros int) rune

func filter(list []string, bit int, rule filterRule) []string {
	if len(list) == 1 {
		return list
	}

	ones, zeros := 0, 0
	for _, l := range list {
		if l[bit] == '1' {
			ones++
		} else {
			zeros++
		}
	}

	keep := rule(ones, zeros)

	var result []string
	for _, l := range list {
		if l[bit] == byte(keep) {
			result = append(result, l)
		}
	}

	return result
}

func process(data []string) int64 {
	oxygenValues := make([]string, len(data))
	co2Values := make([]string, len(data))
	copy(oxygenValues, data)
	copy(co2Values, data)
	for bit := 0; bit < len(data[0]); bit++ {
		oxygenValues = filter(oxygenValues, bit, func(ones, zeros int) rune {
			if ones >= zeros {
				return '1'
			}
			return '0'
		})
		co2Values = filter(co2Values, bit, func(ones, zeros int) rune {
			if zeros <= ones {
				return '0'
			}
			return '1'
		})
	}
	oxygen, err := strconv.ParseInt(oxygenValues[0], 2, 64)
	aoc.PanicIfError(err)
	co2, err := strconv.ParseInt(co2Values[0], 2, 64)
	aoc.PanicIfError(err)

	return oxygen * co2
}
