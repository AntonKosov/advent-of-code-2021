package aoc

import (
	"strconv"
	"strings"
)

func StrToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err.Error())
	}
	return i
}

func StrToInts(s string, sep string) []int {
	var result []int
	parts := strings.Split(s, sep)
	for _, p := range parts {
		if p != "" {
			result = append(result, StrToInt(p))
		}
	}
	return result
}
