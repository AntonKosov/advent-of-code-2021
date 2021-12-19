package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2021/aoc"
)

func main() {
	data := read()
	r := process(data, 12)
	fmt.Printf("Answer: %v\n", r)
}

func read() (data [][]aoc.Vector3) {
	lines := aoc.ReadAllInput()

	i := 0
	for i < len(lines) {
		i++
		scanner := []aoc.Vector3{}
		for ; ; i++ {
			line := lines[i]
			if line == "" {
				break
			}
			parts := aoc.StrToInts(line, ",")
			scanner = append(scanner, aoc.NewVector3(parts[0], parts[1], parts[2]))
		}
		data = append(data, scanner)
		i++
	}

	return data
}

func process(data [][]aoc.Vector3, common int) int {
	baseScanner, otherScanners := prepare(data)

	for len(otherScanners) > 0 {
		var found []int

		for id, s := range otherScanners {
			if mergeScanner(baseScanner, s, common) {
				fmt.Printf("Found: %v\n", id)
				found = append(found, id)
			}
		}

		if len(found) == 0 {
			panic("Solution not found")
		}

		for _, f := range found {
			delete(otherScanners, f)
		}

		fmt.Printf("Found: %v\nRemains: %v\n\n", found, len(otherScanners))
	}

	return len(baseScanner)
}

func mergeScanner(baseScanner map[aoc.Vector3]bool, scanner scanner, common int) (merged bool) {
	for _, v := range scanner.variants {
		if mergeVariant(baseScanner, v, common) {
			return true
		}
	}

	return false
}

func mergeVariant(baseScanner map[aoc.Vector3]bool, variant variant, common int) (merged bool) {
	// Filtering by distance would be more efficient
	for bs := range baseScanner {
		for _, av := range variant.data {
			offset := bs.Sub(av)
			sum := 0
			for _, v := range variant.data {
				if baseScanner[v.Add(offset)] {
					sum++
					if sum == common {
						merge(baseScanner, variant, offset)
						return true
					}
				}
			}
		}
	}

	return false
}

func merge(baseScanner map[aoc.Vector3]bool, variant variant, offset aoc.Vector3) {
	for _, v := range variant.data {
		bs := v.Add(offset)
		baseScanner[bs] = true
	}
}

func prepare(data [][]aoc.Vector3) (baseScanner map[aoc.Vector3]bool, otherScanners map[int]scanner) {
	baseScanner = make(map[aoc.Vector3]bool, len(data[0]))
	for _, v := range data[0] {
		baseScanner[v] = true
	}

	otherScanners = make(map[int]scanner, len(data)-1)
	rotations := allRotations()

	for i := 1; i < len(data); i++ {
		var variants []variant
		for _, rotation := range rotations {
			v := rotate(rotation, data[i])
			variants = append(variants, variant{data: v})
		}

		otherScanners[i] = scanner{variants: variants}
	}

	return baseScanner, otherScanners
}

func rotate(rotation aoc.Matrix3x3, vectors []aoc.Vector3) []aoc.Vector3 {
	var result []aoc.Vector3

	for _, v := range vectors {
		result = append(result, rotation.MulVector3(v))
	}

	return result
}

func allRotations() []aoc.Matrix3x3 {
	faceRotations := []aoc.Matrix3x3{
		aoc.NewIdentityMatrix3x3(),
		aoc.RotationMatrixXY(90),
		aoc.RotationMatrixXY(180),
		aoc.RotationMatrixXY(270),
		aoc.RotationMatrixYZ(90),
		aoc.RotationMatrixYZ(-90),
	}
	topRotations := []aoc.Matrix3x3{
		aoc.NewIdentityMatrix3x3(),
		aoc.RotationMatrixXZ(90),
		aoc.RotationMatrixXZ(180),
		aoc.RotationMatrixXZ(270),
	}

	var result []aoc.Matrix3x3
	for _, fr := range faceRotations {
		for _, tr := range topRotations {
			result = append(result, tr.Mul(fr))
		}
	}

	return result
}

type variant struct {
	data []aoc.Vector3
}

type scanner struct {
	variants []variant
}
