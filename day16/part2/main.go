package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/AntonKosov/advent-of-code-2021/aoc"
)

func main() {
	data := read()
	r := process(data)
	fmt.Printf("Answer: %v\n", r)
}

func read() (data string) {
	return aoc.ReadAllInput()[0]
}

type reader struct {
	bin             string
	currentBitIndex int
}

func newReader(hex string) *reader {
	var sb strings.Builder
	for _, h := range hex {
		i, err := strconv.ParseInt(string(h), 16, 64)
		aoc.PanicIfError(err)
		b := strconv.FormatInt(int64(i), 2)
		for j := len(b); j < 4; j++ {
			sb.WriteRune('0')
		}
		sb.WriteString(b)
	}
	return &reader{bin: sb.String()}
}

func (r *reader) read(bits int) int {
	sr := r.bin[r.currentBitIndex : r.currentBitIndex+bits]
	r.currentBitIndex += bits
	v, err := strconv.ParseInt(sr, 2, 64)
	aoc.PanicIfError(err)
	return int(v)
}

func process(data string) int {
	r := newReader(data)
	value := readPacket(r)

	return value
}

var operators map[int]func([]int) int

func init() {
	operators = map[int]func([]int) int{
		0: sum,
		1: product,
		2: min,
		3: max,
		5: greaterThan,
		6: lessThan,
		7: equalTo,
	}
}

func readPacket(r *reader) int {
	r.read(3) // the packet version is ignored in part 2
	packetType := r.read(3)

	switch packetType {
	case 4:
		return readLiteral(r)
	default:
		args := readArguments(r)
		return operators[packetType](args)
	}
}

func sum(args []int) int {
	s := 0
	for _, v := range args {
		s += v
	}
	return s
}

func product(args []int) int {
	p := 1
	for _, v := range args {
		p *= v
	}
	return p
}

func min(args []int) int {
	result := math.MaxInt
	for _, v := range args {
		result = aoc.Min(result, v)
	}
	return result
}

func max(args []int) int {
	result := math.MinInt
	for _, v := range args {
		result = aoc.Max(result, v)
	}
	return result
}

func greaterThan(args []int) int {
	if args[0] > args[1] {
		return 1
	}
	return 0
}

func lessThan(args []int) int {
	if args[0] < args[1] {
		return 1
	}
	return 0
}

func equalTo(args []int) int {
	if args[0] == args[1] {
		return 1
	}
	return 0
}

func readArguments(r *reader) []int {
	lengthType := r.read(1)
	switch lengthType {
	case 0:
		return read15BitArguments(r)
	case 1:
		return read11BitArguments(r)
	default:
		panic("Unexpected length type")
	}
}

func read11BitArguments(r *reader) (args []int) {
	subPackets := r.read(11)
	for i := 0; i < subPackets; i++ {
		args = append(args, readPacket(r))
	}

	return args
}

func read15BitArguments(r *reader) (args []int) {
	length := r.read(15)
	startIndex := r.currentBitIndex
	for r.currentBitIndex < startIndex+length {
		args = append(args, readPacket(r))
	}

	return args
}

func readLiteral(r *reader) int {
	result := 0
	for {
		v := r.read(5)
		result = (result << 4) | (v & 0b1111)
		if v&0b10000 == 0 {
			return result
		}
	}
}
