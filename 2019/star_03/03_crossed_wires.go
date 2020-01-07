package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	x float64
	y float64
}

type Line struct {
	Start Point
	End   Point
}

func readInput(filename string) string {
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

func determineEndPoint(startingPoint Point, direction string, length int) Point {

	switch direction {
	case "L":
		return Point{startingPoint.x - float64(length), startingPoint.y}
	case "R":
		return Point{startingPoint.x + float64(length), startingPoint.y}
	case "U":
		return Point{startingPoint.x, startingPoint.y + float64(length)}
	case "D":
		return Point{startingPoint.x, startingPoint.y - float64(length)}
	}

	return startingPoint
}

func extractLines(input string) [][]*Line {
	scanner := bufio.NewScanner(strings.NewReader(input))

	lines := [][]*Line{}

	for scanner.Scan() {
		startingPoint := Point{0, 0}

		list := []*Line{}

		s := strings.Split(scanner.Text(), ",")
		for _, p := range s {
			var direction = p[0:1]
			length, _ := strconv.Atoi(p[1:])
			var endPoint = determineEndPoint(startingPoint, direction, length)

			line := new(Line)
			line.Start = startingPoint
			line.End = endPoint

			list = append(list, line)

			startingPoint = endPoint
		}

		lines = append(lines, list)
	}

	return lines
}

func calcDistanceFromCenter(p Point) float64 {
	return Abs(p.y) + Abs(p.x)
}

func Abs(x float64) float64 {
	if x < 0 {
		return x * -1
	}

	return x
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

func main() {
	var input = readInput("input.txt")
	var lines = extractLines(input)

	distances := []float64{}

	wireA := lines[0]
	wireB := lines[1]

	points := []Point{}

	for _, lineA := range wireA {
		for _, lineB := range wireB {
			point, err := findIntersection(lineA, lineB)
			if err != nil {
				continue
			}

			points = append(points, point)
		}
	}

	for _, p := range points {
		var d = calcDistanceFromCenter(p)
		if d != 0 {
			distances = append(distances, d)
		}
	}

	sort.Float64s(distances)
	fmt.Println("solution part 1 :", distances[0])
}
