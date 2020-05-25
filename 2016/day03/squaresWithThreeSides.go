package main

import (
	"fmt"
	"strings"

	"github.com/fowbi/advent-of-code-2016/tools"
)

func isValidTriangle(sideA, sideB, sideC int) bool {
	return (sideA+sideB) > sideC && (sideB+sideC) > sideA && (sideA+sideC) > sideB
}

func main() {
	lines := tools.ReadLines("input.txt")
	validTriangleCount := 0
	validTriangleCount2 := 0

	var column1, column2, column3 []int

	for _, line := range lines {
		values := strings.Fields(line)
		v1 := tools.Atoi(values[0])
		v2 := tools.Atoi(values[1])
		v3 := tools.Atoi(values[2])

		isValidTriangle := isValidTriangle(v1, v2, v3)

		if isValidTriangle {
			validTriangleCount += 1
			fmt.Println(line)
		}
		column1 = append(column1, v1)
		column2 = append(column2, v2)
		column3 = append(column3, v3)
	}

	columns := [][]int{column1, column2, column3}

	var possibleTriangles [][]int
	var triangle []int

	h := -1
	for _, column := range columns {
		for _, row := range column {
			hh := int(tools.Abs(float64(row / 100 % 10)))

			if h == -1 || hh == h {
				triangle = append(triangle, row)

				if len(triangle) == 3 {
					possibleTriangles = append(possibleTriangles, triangle)
					triangle = []int{}
				}
			} else {
				triangle = []int{}
			}

			h = hh
		}
	}
	fmt.Println(possibleTriangles)

	for _, possibleTriangle := range possibleTriangles {
		isValidTriangle := isValidTriangle(possibleTriangle[0], possibleTriangle[1], possibleTriangle[2])

		if isValidTriangle {
			validTriangleCount2 += 1
		}
	}

	fmt.Println("Solution part 1 :", validTriangleCount)
	fmt.Println("Solution part 2 :", validTriangleCount2)
}
