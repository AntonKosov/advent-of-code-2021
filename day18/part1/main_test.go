package main

import (
	"testing"

	"github.com/AntonKosov/advent-of-code-2021/aoc"
)

func testSum(t *testing.T, testName string, expected string, input []string) {
	expressions := parseExpressions(input)
	actual := sum(expressions).String()
	if actual != expected {
		t.Fatalf("%v: expected %v, actual %v", testName, expected, actual)
	}

}

func TestFirstExample(t *testing.T) {
	testSum(t, "TestFirstExample", "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]", []string{
		"[[[[4,3],4],4],[7,[[8,4],9]]]",
		"[1,1]",
	})
}

func Test1234(t *testing.T) {
	testSum(t, "Test1234", "[[[[1,1],[2,2]],[3,3]],[4,4]]", []string{
		"[1,1]",
		"[2,2]",
		"[3,3]",
		"[4,4]",
	})
}

func Test12345(t *testing.T) {
	testSum(t, "Test12345", "[[[[3,0],[5,3]],[4,4]],[5,5]]", []string{
		"[1,1]",
		"[2,2]",
		"[3,3]",
		"[4,4]",
		"[5,5]",
	})
}

func Test123456(t *testing.T) {
	testSum(t, "Test123456", "[[[[5,0],[7,4]],[5,5]],[6,6]]", []string{
		"[1,1]",
		"[2,2]",
		"[3,3]",
		"[4,4]",
		"[5,5]",
		"[6,6]",
	})
}

func TestSlightlyLarger(t *testing.T) {
	testSum(t, "TestSlightlyLarger", "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]", []string{
		"[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]",
		"[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]",
		"[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]",
		"[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]",
		"[7,[5,[[3,8],[1,4]]]]",
		"[[2,[2,2]],[8,[8,1]]]",
		"[2,9]",
		"[1,[[[9,3],9],[[9,0],[0,7]]]]",
		"[[[5,[7,4]],7],1]",
		"[[[[4,2],2],6],[8,7]]",
	})
}

func TestInput(t *testing.T) {
	lines := aoc.ReadAllInputFromFile("input.txt")
	actual := process(lines[:len(lines)-1])
	expected := 4145
	if actual != expected {
		t.Fatalf("TestInput: expected %v, actual %v", expected, actual)
	}
}

func TestMagnitude(t *testing.T) {
	exp := parseExpression("[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]")
	magnitude := exp.magnitude()
	if magnitude != 3488 {
		t.Fatalf("TestMagnitude: expected %v, actual %v", 2488, magnitude)
	}
}
