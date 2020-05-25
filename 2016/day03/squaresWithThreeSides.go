package main

import (
	"fmt"
	"strings"

	"github.com/fowbi/advent-of-code-2016/tools"
)

func isValidTriangle(sideA, sideB, sideC float64) bool {
	return (sideA+sideB) > sideC && (sideB+sideC) > sideA && (sideA+sideC) > sideB
}

func main() {
	lines := tools.ReadLines("input.txt")
	validTriangleCount := 0
	for _, line := range lines {
		values := strings.Fields(line)
		isValidTriangle := isValidTriangle(
			float64(tools.Atoi(values[0])),
			float64(tools.Atoi(values[1])),
			float64(tools.Atoi(values[2])),
		)

		if isValidTriangle {
			validTriangleCount += 1
			fmt.Println(line)
		}
	}

	fmt.Println("Solution part 1 :", validTriangleCount)
}
