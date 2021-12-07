package aoc

import "math"

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func Abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func MinMax(data []int) (min, max int) {
	if len(data) == 0 {
		panic("No data")
	}
	min = math.MaxInt
	max = math.MinInt
	for _, v := range data {
		min = Min(min, v)
		max = Max(max, v)
	}
	return min, max
}
