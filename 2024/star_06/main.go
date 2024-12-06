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

type Guard struct {
	position  image.Point
	direction int
}

func (m Guard) nextPosition(direction image.Point) image.Point {
	return image.Point{m.position.X + direction.X, m.position.Y + direction.Y}
}

type Area struct {
	area [][]string
	size map[string]int
}

type GuardArea struct {
	area  Area
	guard Guard
}

func generateGuardArea(input [][]string) GuardArea {
	area := Area{input, map[string]int{"rows": len(input), "columns": len(input[0])}}
	return GuardArea{area, detectGuard(area)}
}

func detectGuard(guardArea Area) Guard {
	for y, row := range guardArea.area {
		for x, cell := range row {
			i := strings.IndexRune("^>v<", []rune(cell)[0])
			if i == -1 {
				continue
			}

			return Guard{image.Point{x, y}, i}
		}
	}

	panic("No guard found")
}

func (ga GuardArea) inside(position image.Point) bool {
	return position.Y < 0 || position.Y >= ga.area.size["rows"] || position.X < 0 || position.X >= ga.area.size["columns"]
}

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("No input provided")
	}
}

func part_01(input string) int {
	area := utils.InputToMatrix(input)
	guardArea := generateGuardArea(area)
	visitedArea := patrol(guardArea)

	count := 0
	for y, row := range visitedArea {
		for x, cell := range row {
			if cell == "X" || (image.Point{x, y} == guardArea.guard.position) {
				count++
			}
		}
	}

	return count
}

func part_02(input string) int {
	area := utils.InputToMatrix(input)
	guardArea := generateGuardArea(area)
	firstRun := patrol(guardArea)

	validPositions := 0
	for row := 0; row < guardArea.area.size["rows"]; row++ {
		for column := 0; column < guardArea.area.size["columns"]; column++ {
			// We only want visited positions
			if firstRun[row][column] != "X" {
				continue
			}

			validPositions += patrolWithExtraObstacle(guardArea, image.Point{column, row})
		}
	}

	return validPositions
}

func main() {
	fmt.Println("Solution part 1:", part_01(input))
	fmt.Println("Solution part 2:", part_02(input))
}

func patrol(guardArea GuardArea) [][]string {
	directions := utils.AllDirections()
	visitedArea := append([][]string(nil), guardArea.area.area...)

	for {
		nextPosition := guardArea.guard.nextPosition(directions[guardArea.guard.direction])

		// Check if we are out of the guard area
		if guardArea.inside(nextPosition) {
			break
		}

		// Turn 90 degrees to the right when encountering an obstacle
		if guardArea.area.area[nextPosition.Y][nextPosition.X] == "#" {
			guardArea.guard.direction = (guardArea.guard.direction + 1) % 4
		} else {
			// Keep moving forward
			guardArea.guard.position = image.Point{nextPosition.X, nextPosition.Y}
			visitedArea[guardArea.guard.position.Y][guardArea.guard.position.X] = "X"
		}
	}

	return visitedArea
}

// Separate patrol because we do not need keep track of the visited area
func patrolWithExtraObstacle(guardArea GuardArea, obstacle image.Point) int {
	directions := utils.AllDirections()
	counter := 0

	for {
		nextPosition := guardArea.guard.nextPosition(directions[guardArea.guard.direction])

		// Check if we are out of the guard area
		if guardArea.inside(nextPosition) {
			break
		}

		// Yolo-ing out of here :rocket:
		if counter > 10000 {
			return 1
		}
		counter++

		// Turn 90 degrees to the right when encountering an obstacle
		if guardArea.area.area[nextPosition.Y][nextPosition.X] == "#" || (obstacle == nextPosition) {
			guardArea.guard.direction = (guardArea.guard.direction + 1) % 4
		} else { // Keep moving forward
			guardArea.guard.position = image.Point{nextPosition.X, nextPosition.Y}
		}
	}

	return 0
}
