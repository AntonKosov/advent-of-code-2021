package main

import (
	"fmt"
	"strings"

	"github.com/AntonKosov/advent-of-code-2021/aoc"
)

/*
All algorithms for all digits are similar. There are differences in three variables only and z which
is the only value which goes outside. This algorithm may be simplified.

inp w         | w = [1..9]
mul x 0       | x = 0
add x z       | x = x + z
mod x 26      | x = x % 26
div z 1  <- a | z = z / a
add x 14 <- b | x = x + b
eql x w       | x = x == w ? 1 : 0
eql x 0       | x = 1 - x // ("not" x)
mul y 0       | y = 0
add y 25      | y = y + 25
mul y x       | y = y * x
add y 1       | y = y + 1
mul z y       | z = z * y
mul y 0       | y = 0
add y w       | y = y + w
add y 0  <- c | y = y + c
mul y x       | y = y * x
add z y       | z = z * y

So, here is a simplified version:

z = <the result of previous iteration>
w = <digit>
x = z%26 + b
if x == w {
	z = z/a
} else {
	z = z/a*26 + w + c
}
*/
func main() {
	data := read()
	r := process(data)
	fmt.Printf("Answer: %v\n", r)
}

const digits = 14

type args struct {
	a, b, c int
}

func read() (arguments [digits]args) {
	lines := aoc.ReadAllInput()
	for i := 0; i < digits; i++ {
		a := aoc.StrToInt(strings.Split(lines[i*18+4], " ")[2])
		b := aoc.StrToInt(strings.Split(lines[i*18+5], " ")[2])
		c := aoc.StrToInt(strings.Split(lines[i*18+15], " ")[2])
		arguments[i] = args{a: a, b: b, c: c}
	}

	return arguments
}

func process(data [digits]args) int64 {
	solvedVariants := map[variant][]int64{}
	variants := find(0, &data, 0, solvedVariants)
	var max int64
	for _, v := range variants {
		if max < v {
			max = v
		}
	}

	return max
}

type variant struct {
	digitIndex int
	z          int
}

var pows [digits]int64

func init() {
	pows[digits-1] = 1
	for i := digits - 2; i >= 0; i-- {
		pows[i] = pows[i+1] * 10
	}
}

func find(inputZ int, args *[digits]args, digitIndex int, solvedVariants map[variant][]int64) []int64 {
	startVariant := variant{digitIndex: digitIndex, z: inputZ}
	if variants, ok := solvedVariants[startVariant]; ok {
		return variants
	}

	variants := []int64{}
	arguments := args[digitIndex]
	for w := 1; w <= 9; w++ {
		if digitIndex < 2 {
			fmt.Printf("digitIndex: %v, d=%v, cache=%v\n", digitIndex, w, len(solvedVariants))
		}

		z := inputZ
		x := z%26 + arguments.b
		if x == w {
			z = z / arguments.a
		} else {
			z = z/arguments.a*26 + w + arguments.c
		}

		if digitIndex == digits-1 {
			if z == 0 {
				variants = append(variants, int64(w))
			}
			continue
		}

		sv := find(z, args, digitIndex+1, solvedVariants)
		firstDigit := pows[digitIndex] * int64(w)
		for _, v := range sv {
			variants = append(variants, firstDigit+v)
		}

		// Don't need to look for other values
		if digitIndex == 0 && len(variants) > 0 {
			return variants
		}
	}

	solvedVariants[startVariant] = variants

	return variants
}
