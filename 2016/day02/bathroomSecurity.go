package main

import (
	"fmt"

	"github.com/fowbi/advent-of-code-2016/tools"
)

type Point struct {
	x int
	y int
}

type Keypad struct {
	mapping [][]string
	max     int
}

func getNextPoint(p Point, s string, keypad Keypad) Point {
	nextPoint := p

	switch string(s) {
	case "L":
		nextPoint.y -= 1
	case "R":
		nextPoint.y += 1
	case "D":
		nextPoint.x += 1
	default: // U
		nextPoint.x -= 1
	}

	if outOfBounds(nextPoint, keypad.max) {
		return p
	}

	key := determineKey(nextPoint, keypad)
	if key == "" {
		return p
	}

	return nextPoint
}

func outOfBounds(p Point, max int) bool {
	return (p.x > max || p.x < 0 || p.y > max || p.y < 0)
}

func determineKey(p Point, keypad Keypad) string {
	a := keypad.mapping[p.x][p.y]
	return a
}

func main() {
	lines := tools.ReadLines("input.txt")

	p := Point{1, 1}

	keypad := Keypad{
		[][]string{
			{"1", "2", "3"},
			{"4", "5", "6"},
			{"7", "8", "9"},
		},
		2,
	}

	codePartOne := ""
	for _, line := range lines {
		for _, char := range line {
			p = getNextPoint(p, string(char), keypad)
		}
		codePartOne += determineKey(p, keypad)
	}
	fmt.Println("Solution part 1 :", codePartOne)

	p = Point{2, 0}

	keypad = Keypad{
		[][]string{
			{"", "", "1", "", ""},
			{"", "2", "3", "4", ""},
			{"5", "6", "7", "8", "9"},
			{"", "A", "B", "C", ""},
			{"", "", "D", "", ""},
		},
		4,
	}

	codePartTwo := ""
	for _, line := range lines {
		for _, char := range line {
			p = getNextPoint(p, string(char), keypad)
		}
		codePartTwo += determineKey(p, keypad)
	}
	fmt.Println("Solution part 2 :", codePartTwo)
}
