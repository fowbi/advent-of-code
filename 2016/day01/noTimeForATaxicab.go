package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/fowbi/advent-of-code-2016/tools"
)

type Point struct {
	x float64
	y float64
}

type Line struct {
	Start Point
	End   Point
}

func determineNextCardinalPoint(currentCardinalPoint string, direction string) string {
	switch currentCardinalPoint {
	case "N":
		if direction == "L" {
			return "W"
		} else {
			return "E"
		}
	case "E":
		if direction == "L" {
			return "N"
		} else {
			return "S"
		}
	case "S":
		if direction == "L" {
			return "E"
		} else {
			return "W"
		}
	case "W":
		if direction == "L" {
			return "S"
		} else {
			return "N"
		}
	default:
		return currentCardinalPoint
	}
}

func wasPointAlreadyVisited(p Point, points []*Point) bool {
	for _, point := range points {
		if p.x == point.x && p.y == point.y {
			fmt.Println(p.x, p.y, " -- ", point.x, point.y)
			return true
		}
	}
	return false
}

func findIntersection(lineA *Line, lineB *Line) (p Point, err error) {
	denom := ((lineB.End.y - lineB.Start.y) * (lineA.End.x - lineA.Start.x)) - ((lineB.End.x - lineB.Start.x) * (lineA.End.y - lineA.Start.y))

	if denom == 0 {
		return Point{}, errors.New("no intersection")
	}

	a := lineA.Start.y - lineB.Start.y
	b := lineA.Start.x - lineB.Start.x

	num1 := ((lineB.End.x - lineB.Start.x) * a) - ((lineB.End.y - lineB.Start.y) * b)
	num2 := ((lineA.End.x - lineA.Start.x) * a) - ((lineA.End.y - lineA.Start.y) * b)

	c := num1 / denom
	d := num2 / denom

	if (c > 0 && c < 1) && (d > 0 && d < 1) {
		return Point{
			lineA.Start.x + (c * (lineA.End.x - lineA.Start.x)),
			lineA.Start.y + (c * (lineA.End.y - lineA.Start.y)),
		}, nil
	}

	return Point{}, errors.New("no intersection")
}

func getFirstVisitedPoint(lines []*Line) Point {
	for _, lineA := range lines {
		for _, lineB := range lines {
			point, err := findIntersection(lineA, lineB)
			if err != nil {
				continue
			}

			return point
		}
	}
	return Point{0, 0}
}

func main() {
	steps := tools.Split(tools.ReadInput("input.txt"))

	var cardinalPoint = "N"
	startingPoint := Point{0, 0}
	endPoint := Point{0, 0}
	lines := []*Line{}

	for _, step := range steps {
		var direction = step[0:1]
		length, _ := strconv.Atoi(step[1:])
		line := new(Line)
		line.Start = Point{endPoint.x, endPoint.y}

		nextCardinalPoint := determineNextCardinalPoint(cardinalPoint, direction)

		switch nextCardinalPoint {
		case "N":
			endPoint.y += float64(length)
		case "E":
			endPoint.x += float64(length)
		case "S":
			endPoint.y -= float64(length)
		case "W":
			endPoint.x -= float64(length)
		}

		line.End = Point{endPoint.x, endPoint.y}
		lines = append(lines, line)

		cardinalPoint = nextCardinalPoint
	}
	distance := tools.Abs(startingPoint.x-endPoint.x) + tools.Abs(startingPoint.y-endPoint.y)
	firstVisitedPoint := getFirstVisitedPoint(lines)
	firstVisitedDistance := tools.Abs(startingPoint.x-firstVisitedPoint.x) + tools.Abs(startingPoint.y-firstVisitedPoint.y)

	fmt.Println("solution part 1 :", distance)
	fmt.Println("solution part 2 :", firstVisitedDistance)
}
