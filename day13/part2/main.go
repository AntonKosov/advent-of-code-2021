package main

import (
	"fmt"
	"strings"

	"github.com/AntonKosov/advent-of-code-2021/aoc"
)

func main() {
	data := read()
	process(data)
}

type folding struct {
	isVertical bool
	line       int
}

type instructions struct {
	dots     map[aoc.Vector2]bool
	width    int
	height   int
	foldings []folding
}

func read() (data instructions) {
	lines := aoc.ReadAllInput()
	data.dots = make(map[aoc.Vector2]bool)

	i := 0
	for {
		line := lines[i]
		if line == "" {
			break
		}
		s := strings.Split(line, ",")
		c := aoc.NewVector2(aoc.StrToInt(s[0]), aoc.StrToInt(s[1]))
		data.dots[c] = true
		data.width = aoc.Max(data.width, c.X+1)
		data.height = aoc.Max(data.height, c.Y+1)
		i++
	}

	for j := i + 1; j < len(lines); j++ {
		line := lines[j]
		if line == "" {
			continue
		}
		s := strings.Split(line[len("fold along "):], "=")
		isVertical := true
		if s[0] == "y" {
			isVertical = false
		}
		data.foldings = append(data.foldings, folding{isVertical: isVertical, line: aoc.StrToInt(s[1])})
	}

	return data
}

func process(data instructions) {
	for _, f := range data.foldings {
		if f.isVertical {
			for y := 0; y < data.height; y++ {
				delete(data.dots, aoc.NewVector2(f.line, y))
			}
			data.width = f.line
			offset := 2 * f.line
			for c, v := range data.dots {
				if !v || c.X < data.width {
					continue
				}
				data.dots[aoc.NewVector2(offset-c.X, c.Y)] = true
				data.dots[c] = false
			}
		} else {
			for x := 0; x < data.width; x++ {
				delete(data.dots, aoc.NewVector2(x, f.line))
			}
			data.height = f.line
			offset := 2 * f.line
			for c, v := range data.dots {
				if !v || c.Y < data.height {
					continue
				}
				data.dots[aoc.NewVector2(c.X, offset-c.Y)] = true
				data.dots[c] = false
			}
		}
	}

	for y := 0; y < data.height; y++ {
		for x := 0; x < data.width; x++ {
			if data.dots[aoc.NewVector2(x, y)] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
