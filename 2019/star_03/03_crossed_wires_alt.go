package main

import (
	"bufio"
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

func readInput(filename string) string {
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

func Abs(x float64) float64 {
	if x < 0 {
		return x * -1
	}

	return x
}

func addDirection(p1 *Point, p2 Point) *Point {
	p := new(Point)

	p.x = p1.x + p2.x
	p.y = p1.y + p2.y

	return p
}

func comparePoints(p1 *Point, p2 *Point) bool {
	return p1.x == p2.x && p1.y == p2.y
}

func determineDirection(direction string) Point {
	switch direction {
	case "L":
		return Point{-1, 0}
	case "R":
		return Point{1, 0}
	case "U":
		return Point{0, 1}
	case "D":
		return Point{0, -1}
	}

	return Point{0, 0}
}

func extractPoints(line string) []*Point {
	s := strings.Split(line, ",")
	var pointer = new(Point)
	pointer.x = 0
	pointer.y = 0

	points := []*Point{}

	var startingPoint = new(Point)
	startingPoint.x = 0
	startingPoint.y = 0
	points = append(points, startingPoint)

	for _, p := range s {
		length, _ := strconv.Atoi(p[1:])
		direction := determineDirection(p[0:1])

		for i := 0; i < length; i++ {
			point := addDirection(pointer, direction)
			points = append(points, point)
			pointer = point
		}
	}

	return points
}

func calcDistanceFromCenter(p *Point) float64 {
	return Abs(p.y) + Abs(p.x)
}

func main() {
	var input = readInput("input.txt")

	distances := []float64{}
	wires := [][]*Point{}

	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		points := extractPoints(scanner.Text())
		wires = append(wires, points)
	}

	wireA := wires[0]
	wireB := wires[1]

	for _, pointA := range wireA {
		for _, pointB := range wireB {
			if comparePoints(pointA, pointB) {
				var d = calcDistanceFromCenter(pointA)
				if d != 0 {
					distances = append(distances, d)
				}
			}
		}
	}

	sort.Float64s(distances)
	fmt.Println("solution part 1 :", distances[0])
}
