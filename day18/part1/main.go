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

func read() (data []string) {
	lines := aoc.ReadAllInput()

	return lines[:len(lines)-1]
}

const maxValue = 9

func process(data []string) int {
	expressions := parseExpressions(data)
	s := sum(expressions)
	m := s.magnitude()

	return m
}

func sum(expressions []*pair) *pair {
	sum := expressions[0]
	for i := 1; i < len(expressions); i++ {
		sum = sum.add(expressions[i])
	}
	return sum
}

func parseExpressions(expressions []string) (parsed []*pair) {
	for _, e := range expressions {
		parsed = append(parsed, parseExpression(e))
	}

	return parsed
}

type pair struct {
	parent *pair

	leftValue *int
	leftPair  *pair

	rightValue *int
	rightPair  *pair
}

func splitExpression(e string) (left, right string) {
	p := e[1 : len(e)-1]
	ob := 0
	for i, r := range p {
		switch r {
		case '[':
			ob++
		case ']':
			ob--
		case ',':
			if ob == 0 {
				return p[:i], p[i+1:]
			}
		}
	}

	panic("Failed to split an expression")
}

func parseExpression(s string) *pair {
	left, right := splitExpression(s)
	p := &pair{}
	if left[0] == '[' {
		p.setLeftPair(parseExpression(left))
	} else {
		p.setLeftValue(aoc.StrToInt(left))
	}
	if right[0] == '[' {
		p.setRightPair(parseExpression(right))
	} else {
		p.setRightValue(aoc.StrToInt(right))
	}

	return p
}

func (p *pair) add(rightPair *pair) *pair {
	// fmt.Printf(" %v\n+%v\n", p, rightPair)
	result := &pair{}
	result.setLeftPair(p)
	result.setRightPair(rightPair)

	// fmt.Printf("=%v\n", result)

	for {
		if p := result.findPairByLevel(0, 4); p != nil {
			p.explode()
			// fmt.Printf("e%v\n", result)
			continue
		}
		if p := result.findTooBig(); p != nil {
			p.split()
			// fmt.Printf("s%v\n", result)
			continue
		}
		break
	}
	// fmt.Printf("Done\n\n")

	return result
}

func (p *pair) findPairByLevel(currentLevel, requiredLevel int) *pair {
	if currentLevel == requiredLevel {
		return p
	}

	if p.leftPair != nil {
		if lp := p.leftPair.findPairByLevel(currentLevel+1, requiredLevel); lp != nil {
			return lp
		}
	}

	if p.rightPair != nil {
		if rp := p.rightPair.findPairByLevel(currentLevel+1, requiredLevel); rp != nil {
			return rp
		}
	}

	return nil
}

func (p *pair) findLeftValue() **int {
	t := p.parent
	prev := p
	for t != nil {
		if t.leftPair == nil {
			return &t.leftValue
		}
		if t.leftPair != prev {
			t = t.leftPair
			for t.rightPair != nil {
				t = t.rightPair
			}
			return &t.rightValue
		}
		prev = t
		t = t.parent
	}

	return nil
}

func (p *pair) findRightValue() **int {
	t := p.parent
	prev := p
	for t != nil {
		if t.rightPair == nil {
			return &t.rightValue
		}
		if t.rightPair != prev {
			t = t.rightPair
			for t.leftPair != nil {
				t = t.leftPair
			}
			return &t.leftValue
		}
		prev = t
		t = t.parent
	}

	return nil
}

func (p *pair) isLeft() bool {
	return p.parent.leftPair == p
}

func (p *pair) explode() {
	leftValuePP := p.findLeftValue()
	rightValuePP := p.findRightValue()

	if leftValuePP == nil {
		*rightValuePP = ref(*p.rightValue + **rightValuePP)
		p.parent.setLeftValue(0)
		return
	}

	if rightValuePP == nil {
		np := &pair{}
		np.setRightValue(0)
		np.setLeftValue(*p.leftValue + **leftValuePP)
		p.parent.parent.setRightPair(np)
		return
	}

	*leftValuePP = ref(*p.leftValue + **leftValuePP)
	*rightValuePP = ref(*p.rightValue + **rightValuePP)

	if p.isLeft() {
		p.parent.setLeftValue(0)
	} else {
		p.parent.setRightValue(0)
	}
}

func (p *pair) findTooBig() *pair {
	if p.leftPair != nil {
		if lp := p.leftPair.findTooBig(); lp != nil {
			return lp
		}
	} else if *p.leftValue > maxValue {
		return p
	}

	if p.rightPair != nil {
		if rp := p.rightPair.findTooBig(); rp != nil {
			return rp
		}
	} else if *p.rightValue > maxValue {
		return p
	}

	return nil
}

func (p *pair) split() {
	createPair := func(v int) *pair {
		np := &pair{}
		np.setLeftValue(v / 2)
		np.setRightValue(v/2 + v%2)
		return np
	}

	if p.leftValue != nil && *p.leftValue > maxValue {
		p.setLeftPair(createPair(*p.leftValue))
		return
	}

	p.setRightPair(createPair(*p.rightValue))
}

func (p *pair) setLeftValue(v int) {
	if p.leftPair != nil {
		p.leftPair.parent = nil
	}
	p.leftPair = nil
	p.leftValue = &v
}

func (p *pair) setRightValue(v int) {
	if p.rightPair != nil {
		p.rightPair.parent = nil
	}
	p.rightPair = nil
	p.rightValue = &v
}

func (p *pair) setLeftPair(lp *pair) {
	lp.parent = p
	p.leftValue = nil
	p.leftPair = lp
}

func (p *pair) setRightPair(rp *pair) {
	rp.parent = p
	p.rightValue = nil
	p.rightPair = rp
}

func (p *pair) magnitude() int {
	var left int
	var right int

	if p.leftValue != nil {
		left = *p.leftValue
	} else {
		left = p.leftPair.magnitude()
	}

	if p.rightValue != nil {
		right = *p.rightValue
	} else {
		right = p.rightPair.magnitude()
	}

	return 3*left + 2*right
}

func (p *pair) String() string {
	var sb strings.Builder
	sb.WriteRune('[')
	if p.leftValue != nil {
		sb.WriteString(strconv.Itoa(*p.leftValue))
	} else {
		sb.WriteString(p.leftPair.String())
	}
	sb.WriteRune(',')
	if p.rightValue != nil {
		sb.WriteString(strconv.Itoa(*p.rightValue))
	} else {
		sb.WriteString(p.rightPair.String())
	}
	sb.WriteRune(']')

	return sb.String()
}

func ref(v int) *int {
	return &v
}
