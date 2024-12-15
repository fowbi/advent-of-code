package main

import (
	"aoc_2024/utils"
	_ "embed"
	"fmt"
	"image"
	"strings"
)

//go:embed input.txt
var input string

type Position struct {
	id    string
	xy    image.Point
	value int
}

func (p Position) next(coord image.Point) image.Point {
	return image.Point{p.xy.X + coord.X, p.xy.Y + coord.Y}
}

func (p Position) nextId(coord image.Point) string {
	return fmt.Sprintf("%d_%d", p.xy.X+coord.X, p.xy.Y+coord.Y)
}

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("No input provided")
	}
}

func part_01(input string) int {
	grid, zeroes := extractGridAndZeroes(input)
	result := 0
	counter := 0
	for _, zeroPosition := range zeroes {
		nines, _ := extract(grid, zeroPosition, counter, map[string]bool{})
		result += len(nines)
	}

	return result
}

func part_02(input string) int {
	grid, zeroes := extractGridAndZeroes(input)
	result := 0
	counter := 0
	for _, zeroPosition := range zeroes {
		_, c := extract(grid, zeroPosition, counter, map[string]bool{})
		result += c
	}

	return result
}

func extractGridAndZeroes(input string) (map[string]Position, []string) {
	grid := map[string]Position{}
	zeroes := []string{}
	for x, row := range strings.Split(input, "\n") {
		if row == "" {
			continue
		}

		for y, column := range strings.Split(row, "") {
			value := -1
			if column != "." {
				value = utils.CastStringToInt(column)
			}
			id := fmt.Sprintf("%d_%d", x, y)
			if value == 0 {
				zeroes = append(zeroes, id)
			}

			grid[id] = Position{id, image.Point{x, y}, value}
		}
	}

	return grid, zeroes
}

func extract(grid map[string]Position, position string, counter int, nines map[string]bool) (map[string]bool, int) {
	higherGrounds := getHigherGrounds(grid, grid[position])
	for _, higherGround := range higherGrounds {
		if grid[higherGround].value == 9 {
			nines[higherGround] = true
			counter += 1
			continue
		}
		n, c := extract(grid, higherGround, counter, nines)
		for k, v := range n {
			nines[k] = v
		}
		counter = c
	}

	return nines, counter
}

func getHigherGrounds(grid map[string]Position, position Position) (higherGrounds []string) {
	for _, direction := range utils.AllDirections() {
		nextId := position.nextId(direction)
		if next, ok := grid[nextId]; ok && next.value == (position.value+1) {
			higherGrounds = append(higherGrounds, nextId)
		}
	}

	return higherGrounds
}

func main() {
	fmt.Println("Solution part 1:", part_01(input))
	fmt.Println("Solution part 2:", part_02(input))
}
