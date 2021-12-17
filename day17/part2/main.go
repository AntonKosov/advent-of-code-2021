package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/AntonKosov/advent-of-code-2021/aoc"
)

func main() {
	data := read()
	r := process(data)
	fmt.Printf("Answer: %v\n", r)
}

type target struct {
	min aoc.Vector2
	max aoc.Vector2
}

func read() (data target) {
	line := aoc.ReadAllInput()[0]

	s := strings.Split(line[15:], ", y=")
	xs := strings.Split(s[0], "..")
	ys := strings.Split(s[1], "..")

	data.min.X = aoc.StrToInt(xs[0])
	data.min.Y = aoc.StrToInt(ys[1])
	data.max.X = aoc.StrToInt(xs[1])
	data.max.Y = aoc.StrToInt(ys[0])

	return data
}

func process(data target) int {
	minXVelocity := getMinVelocity(data.min.X)
	sum := 0
	const maxVelocity = 200
	for startXV := minXVelocity; startXV <= data.max.X; startXV++ {
		for startYV := data.max.Y; startYV <= maxVelocity; startYV++ {
			p := aoc.NewVector2(startXV, startYV)
			xv := startXV
			yv := startYV
			for p.X <= data.max.X && p.Y >= data.max.Y {
				if p.X >= data.min.X && p.Y <= data.min.Y {
					sum++
					break
				}
				if xv == 0 && (p.X < data.min.X || p.X > data.max.X) {
					break
				}
				xv = aoc.Max(0, xv-1)
				yv -= 1
				p.X += xv
				p.Y += yv
			}
		}
	}

	return sum
}

func getMinVelocity(distance int) int {
	return (-1 + int(math.Sqrt(float64(1+8*distance)))) / 2
}
