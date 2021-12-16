package main

import (
	"fmt"
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

	return readPacket(r)
}

func readPacket(r *reader) int {
	version := r.read(3)
	sum := version

	packetType := r.read(3)
	switch packetType {
	case 4:
		readLiteral(r)
	default:
		sum += readOperator(r)
	}

	return sum
}

func readOperator(r *reader) int {
	lengthType := r.read(1)
	switch lengthType {
	case 0:
		return read15BitOperator(r)
	case 1:
		return read11BitOperator(r)
	default:
		panic("Unexpected length type")
	}
}

func read11BitOperator(r *reader) int {
	sum := 0
	subPackets := r.read(11)
	for i := 0; i < subPackets; i++ {
		sum += readPacket(r)
	}

	return sum
}

func read15BitOperator(r *reader) int {
	sum := 0
	length := r.read(15)
	startIndex := r.currentBitIndex
	for r.currentBitIndex < startIndex+length {
		sum += readPacket(r)
	}

	return sum
}

func readLiteral(r *reader) {
	for {
		v := r.read(5)
		if v&0b10000 == 0 {
			return
		}
	}
}
