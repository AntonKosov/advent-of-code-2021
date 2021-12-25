package main

import (
	"fmt"
	"strings"

	"github.com/AntonKosov/advent-of-code-2021/aoc"
)

func main() {
	data := read()
	r := process(data)
	fmt.Printf("Answer: %v\n", r)
}

type registers [4]int

const regZ = 3
const digits = 14

type command struct {
	isInput          bool
	targetIndex      int
	argRegisterIndex *int
	argValue         *int
	operation        func(target *int, argValue int) (ok bool)
}

func (c command) exec(r *registers) bool {
	arg := 0
	if c.argRegisterIndex != nil {
		arg = r[*c.argRegisterIndex]
	} else if c.argValue != nil {
		arg = *c.argValue
	}
	return c.operation(&r[c.targetIndex], arg)
}

func (c command) setInput(r *registers, value int) {
	r[c.targetIndex] = value
}

func read() [digits][]command {
	lines := aoc.ReadAllInput()
	return readCommands(lines)
}

func readCommands(input []string) (data [digits][]command) {
	add := func(t *int, arg int) bool {
		*t = *t + arg
		return true
	}
	mul := func(t *int, arg int) bool {
		*t = *t * arg
		return true
	}
	div := func(t *int, arg int) bool {
		if arg == 0 {
			return false
		}
		*t = *t / arg
		return true
	}
	mod := func(t *int, arg int) bool {
		if *t < 0 || arg <= 0 {
			return false
		}
		*t = *t % arg
		return true
	}
	eql := func(t *int, arg int) bool {
		if *t == arg {
			*t = 1
		} else {
			*t = 0
		}
		return true
	}

	regIndex := func(s string) int {
		return int(s[0] - byte('w'))
	}

	for i := range data {
		for j := 0; j < 18; j++ {
			p := strings.Split(input[i*18+j], " ")
			c := command{targetIndex: regIndex(p[1])}
			switch p[0] {
			case "inp":
				c.isInput = true
			case "add":
				c.operation = add
			case "mul":
				c.operation = mul
			case "div":
				c.operation = div
			case "mod":
				c.operation = mod
			case "eql":
				c.operation = eql
			default:
				panic("Unknown operation")
			}

			if len(p) > 2 {
				arg := p[2]
				if arg[0] >= byte('w') {
					index := regIndex(arg)
					c.argRegisterIndex = &index
				} else {
					v := aoc.StrToInt(arg)
					c.argValue = &v
				}
			}

			data[i] = append(data[i], c)
		}
	}

	return data
}

func process(data [digits][]command) int64 {
	solvedVariants := map[variant][]int64{}
	variants := find(0, data, 0, solvedVariants)
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

func find(z int, allCommands [digits][]command, digitIndex int, solvedVariants map[variant][]int64) []int64 {
	startVariant := variant{digitIndex: digitIndex, z: z}
	if variants, ok := solvedVariants[startVariant]; ok {
		return variants
	}

	variants := []int64{}
	commands := allCommands[digitIndex]
nextDigit:
	for d := 9; d > 0; d-- {
		if digitIndex < 2 {
			fmt.Printf("digitIndex: %v, d=%v, cache=%v\n", digitIndex, d, len(solvedVariants))
		}
		reg := registers{}
		reg[regZ] = z
		commands[0].setInput(&reg, d)

		for i := 1; i < len(commands); i++ {
			if !commands[i].exec(&reg) {
				continue nextDigit
			}
		}

		currentZ := reg[regZ]
		if digitIndex == digits-1 {
			if currentZ == 0 {
				variants = append(variants, int64(d))
			}
			continue
		}

		sv := find(currentZ, allCommands, digitIndex+1, solvedVariants)
		firstDigit := pows[digitIndex] * int64(d)
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
