package aoc

import (
	"os"
	"strings"
)

func ReadAllInput() []string {
	if len(os.Args) != 2 {
		panic("wrong arguments")
	}
	bytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err.Error())
	}

	lines := strings.Split(string(bytes), "\n")

	return lines
}
