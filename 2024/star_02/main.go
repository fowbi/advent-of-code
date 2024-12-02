package main

import (
	"aoc_2024/utils"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("No input provided")
	}
}

func part_01(input string) int {
	safe := 0

	for _, report := range strings.Split(input, "\n") {
		if report == "" {
			continue
		}
		levels, _ := utils.ExtractNumbers(report)
		if safeReport(levels) {
			safe++
		}
	}

	return safe
}

func part_02(input string) int {
	safe := 0

	for _, report := range strings.Split(input, "\n") {
		if report == "" {
			continue
		}
		levels, _ := utils.ExtractNumbers(report)
		if safeReport(levels) {
			safe++
			continue
		}
		if tolerate(levels) {
			safe++
		}
	}

	return safe
}

func main() {
	fmt.Println("Solution part 1:", part_01(input))
	fmt.Println("Solution part 2:", part_02(input))
}

func tolerate(levels []int) bool {
	for k, _ := range levels {
		potentialToleratedReport := utils.Remove(levels, k)
		if safeReport(potentialToleratedReport) {
			return true
		}
	}
	return false
}

func safeReport(levels []int) bool {
	direction := 0
	for k, level := range levels {
		if k == 0 {
			continue
		}

		previous := levels[k-1]
		// skip when a level is the same as the previous
		if level == previous {
			return false
		}

		newDirection := 0
		if level > previous {
			newDirection = 1
		}
		if level < previous {
			newDirection = -1
		}
		if direction != 0 && direction != newDirection {
			return false
		}
		direction = newDirection

		if utils.Diff(level, previous) > 3 {
			return false
		}
	}
	return true
}
