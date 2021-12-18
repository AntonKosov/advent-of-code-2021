package aoc

import (
	"os"
	"strings"
)

func ReadAllInput() []string {
	if len(os.Args) != 2 {
		panic("wrong arguments")
	}
	return ReadAllInputFromFile(os.Args[1])
}

func ReadAllInputFromFile(name string) []string {
	bytes, err := os.ReadFile(name)
	if err != nil {
		panic(err.Error())
	}

	lines := strings.Split(string(bytes), "\n")

	return lines
}
