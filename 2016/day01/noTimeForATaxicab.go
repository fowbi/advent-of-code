package main

import (
	"fmt"
	"strconv"

	"github.com/fowbi/advent-of-code-2016/tools"
)

type Point struct {
	x int
	y int
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

func main() {
	steps := tools.Split(tools.ReadInput("input.txt"))

	var cardinalPoint = "N"
	startingPoint := Point{0, 0}
	endPoint := Point{0, 0}

	for _, step := range steps {
		var direction = step[0:1]
		length, _ := strconv.Atoi(step[1:])

		nextCardinalPoint := determineNextCardinalPoint(cardinalPoint, direction)

		switch nextCardinalPoint {
		case "N":
			endPoint.y += length
		case "E":
			endPoint.x += length
		case "S":
			endPoint.y -= length
		case "W":
			endPoint.x -= length
		}

		cardinalPoint = nextCardinalPoint
	}
	distance := tools.Abs(startingPoint.x-endPoint.x) + tools.Abs(startingPoint.y-endPoint.y)

	fmt.Println(distance)
}
