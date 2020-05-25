package main

import (
	"fmt"

	"github.com/fowbi/advent-of-code-2016/tools"
)

type Point struct {
	x float64
	y float64
}

func getNextPoint(p Point, s string) Point {
	nextPoint := p

	switch string(s) {
	case "L":
		nextPoint.x -= 1
	case "R":
		nextPoint.x += 1
	case "D":
		nextPoint.y -= 1
	default: // U
		nextPoint.y += 1
	}

	if outOfBounds(nextPoint) {
		return p
	}

	return nextPoint
}

func outOfBounds(p Point) bool {
	return (p.x > 1 || p.x < -1 || p.y > 1 || p.y < -1)
}

func determineKey(p Point) int {
	if p.x == -1 && p.y == 1 {
		return 1
	}

	if p.x == 0 && p.y == 1 {
		return 2
	}

	if p.x == 1 && p.y == 1 {
		return 3
	}

	if p.x == -1 && p.y == 0 {
		return 4
	}

	if p.x == 0 && p.y == 0 {
		return 5
	}

	if p.x == 1 && p.y == 0 {
		return 6
	}

	if p.x == -1 && p.y == -1 {
		return 7
	}

	if p.x == 0 && p.y == -1 {
		return 8
	}

	return 9
}

func main() {
	lines := tools.ReadLines("input.txt")

	p := Point{0, 0}

	for _, line := range lines {
		for _, char := range line {
			p = getNextPoint(p, string(char))
		}
		fmt.Println(p, determineKey(p))
	}
}
