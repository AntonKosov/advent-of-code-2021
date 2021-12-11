package main

import (
	"log"
	"time"

	"github.com/AntonKosov/advent-of-code-2021/aoc"
	"github.com/gdamore/tcell"
)

func main() {
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("Cannot initialize terminal: %v", err.Error())
	}
	if err := screen.Init(); err != nil {
		log.Fatalf("Cannot initialize screen: %v", err.Error())
	}
	defer screen.Fini()

	data := read()
	process(data, screen)
}

func read() (data map[aoc.Vector2]int) {
	lines := aoc.ReadAllInput()
	data = make(map[aoc.Vector2]int, 100)

	for y := 0; y < 10; y++ {
		line := lines[y]
		for x, v := range line {
			data[aoc.NewVector2(x, y)] = int(v - '0')
		}
	}

	return data
}

var styles map[int]tcell.Style

func init() {
	styles = make(map[int]tcell.Style, 10)
	for i := 0; i < 11; i++ {
		var color int32 = int32(i * 25)
		styles[i] = tcell.StyleDefault.Background(tcell.NewRGBColor(color, color, color))
	}
}

func output(c aoc.Vector2, value int, screen tcell.Screen) {
	w, h := screen.Size()
	sx := (w - 10) / 2
	sy := (h - 10) / 2
	style := styles[value]
	screen.SetContent(sx+c.X*2, sy+c.Y, ' ', nil, style)
	screen.SetContent(sx+c.X*2+1, sy+c.Y, ' ', nil, style)
}

func process(data map[aoc.Vector2]int, screen tcell.Screen) {
	draw := func() {
		screen.Show()
		time.Sleep(time.Millisecond * 100)
	}
	for c, v := range data {
		output(c, v, screen)
	}
	draw()
	for i := 0; ; i++ {
		queue := make(map[aoc.Vector2]bool)
		flashed := make(map[aoc.Vector2]bool)
		charge := func(c aoc.Vector2) {
			if flashed[c] {
				return
			}
			data[c]++
			output(c, data[c], screen)
			if data[c] == 10 {
				queue[c] = true
			}
		}

		for c := range data {
			charge(c)
		}

		flashes := 0
		for len(queue) > 0 {
			q := queue
			queue = make(map[aoc.Vector2]bool)
			for c := range q {
				flashes++
				flashed[c] = true
				data[c] = 0
				output(c, data[c], screen)
				for _, a := range c.Adjacent() {
					if _, ok := data[a]; ok {
						charge(a)
					}
				}
			}
		}
		draw()
		if flashes == 100 {
			time.Sleep(time.Second * 2)
			return
		}
	}
}
