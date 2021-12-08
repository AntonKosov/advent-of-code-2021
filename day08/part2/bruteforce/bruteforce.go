package bruteforce

var imageValues map[byte]byte

func init() {
	//  aaaa
	// b    c
	// b    c
	//  dddd
	// e    f
	// e    f
	//  gggg
	imageValues = map[byte]byte{
		//abcdefg
		0b1110111: 0,
		0b0010010: 1,
		0b1011101: 2,
		0b1011011: 3,
		0b0111010: 4,
		0b1101011: 5,
		0b1101111: 6,
		0b1010010: 7,
		0b1111111: 8,
		0b1111011: 9,
	}
}

func Process(digits, value []string) int {
	segments := findSegments(digits)
	return readValue(value, segments)
}

func readValue(digits []string, segments map[rune]byte) int {
	value := 0
	for _, digit := range digits {
		var image byte
		for _, r := range digit {
			image |= segments[r]
		}
		value = 10*value + int(imageValues[image])
	}
	return value
}

func findSegments(digits []string) map[rune]byte { // rune -> segment (bit)
	runes := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g'}
	availablePositions := make(map[int]struct{}, 10)
	for i := 0; i < 7; i++ {
		availablePositions[i] = struct{}{}
	}
	segments := make(map[rune]byte, 10)
	if !solve(digits, runes, 0, availablePositions, segments) {
		panic("No solution")
	}
	return segments
}

func solve(digits []string, runes []rune, runeIndex int, availablePositions map[int]struct{},
	segments map[rune]byte,
) bool {
	currentRune := runes[runeIndex]
	positions := make([]int, 0, len(availablePositions))
	for p := range availablePositions {
		positions = append(positions, p)
	}
	isLastSegment := len(availablePositions) == 1
	for _, position := range positions {
		delete(availablePositions, position)
		segments[currentRune] = byte(1 << position)
		if isLastSegment {
			if isValid(digits, segments) {
				return true
			}
		} else if solve(digits, runes, runeIndex+1, availablePositions, segments) {
			return true
		}
		availablePositions[position] = struct{}{}
	}
	delete(segments, currentRune)
	return false
}

func isValid(digits []string, segments map[rune]byte) bool {
	for _, digit := range digits {
		var v byte
		for _, segment := range digit {
			v |= segments[segment]
		}
		if _, ok := imageValues[v]; !ok {
			return false
		}
	}
	return true
}
