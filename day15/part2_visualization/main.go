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

func read() (data [][]int) {
	lines := aoc.ReadAllInput()

	var tile [][]int
	for _, line := range lines {
		if line == "" {
			continue
		}
		row := make([]int, len(line))
		for i, v := range line {
			row[i] = int(v - '0')
		}
		tile = append(tile, row)
	}

	tileW := len(tile[0])
	tileH := len(tile)
	data = make([][]int, tileH*5)
	for i := range data {
		data[i] = make([]int, tileW*5)
	}

	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			for tr := 0; tr < tileH; tr++ {
				for tc := 0; tc < tileW; tc++ {
					v := tile[tr][tc] + row + col
					if v > 9 {
						v = v%10 + 1
					}
					data[row*tileH+tr][col*tileW+tc] = v
				}
			}
		}
	}

	return data
}

var styles map[int]tcell.Style

func init() {
	styles = make(map[int]tcell.Style, 256)
	for i := 0; i < 256; i++ {
		var color int32 = int32(i)
		styles[i] = tcell.StyleDefault.Background(tcell.NewRGBColor(color, color, color))
	}
}

func output(row, col, value int, screen tcell.Screen) {
	style := styles[value]
	screen.SetContent(col*2, row, ' ', nil, style)
	screen.SetContent(col*2+1, row, ' ', nil, style)
}

type pos struct {
	row, col int
}

func process(data [][]int, screen tcell.Screen) {
	draw := func() {
		screen.Show()
		time.Sleep(time.Millisecond * 5)
	}
	queue := []pos{{0, 0}}
	width := len(data[0])
	height := len(data)

	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			output(r, c, 0, screen)
		}
	}
	draw()

	m := make([][]int, height)
	for i := range m {
		m[i] = make([]int, width)
	}

	for len(queue) > 0 {
		p := queue[0]
		output(p.row, p.col, 0, screen)
		queue = queue[1:]
		currentTotalLevel := m[p.row][p.col]
		checkPosition := func(row, col int) {
			level := data[row][col]
			previousTotalLevel := m[row][col]
			if previousTotalLevel == 0 || currentTotalLevel+level < previousTotalLevel {
				m[row][col] = currentTotalLevel + level
				queue = append(queue, pos{row, col})
			}
		}
		if p.row > 0 {
			checkPosition(p.row-1, p.col)
		}
		if p.col > 0 {
			checkPosition(p.row, p.col-1)
		}
		if p.col < width-1 {
			checkPosition(p.row, p.col+1)
		}
		if p.row < height-1 {
			checkPosition(p.row+1, p.col)
		}

		for i, v := range queue {
			color := 255 / (len(queue) / (len(queue) - i))
			output(v.row, v.col, color, screen)
		}

		draw()
	}
}
