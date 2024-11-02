package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
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
	overlappingPairs := 0

	for _, pair := range parseInput(input) {
		if pair.FullyOverlap() {
			overlappingPairs++
		}
	}
	return overlappingPairs
}

func part_02(input string) int {
	overlappingPairs := 0

	for _, pair := range parseInput(input) {
		if pair.Overlap() {
			overlappingPairs++
		}
	}
	return overlappingPairs
}

func main() {
	fmt.Println("Solution part 1:", part_01(input))
	fmt.Println("Solution part 2:", part_02(input))
}

type Pair struct {
	x, y int
}
type SectionAssignmentPairs struct {
	first  Pair
	second Pair
}

func (sap *SectionAssignmentPairs) FullyOverlap() bool {
	if sap.first.x <= sap.second.x && sap.first.y >= sap.second.y {
		return true
	}

	if sap.second.x <= sap.first.x && sap.second.y >= sap.first.y {
		return true
	}

	return false
}

func (sap *SectionAssignmentPairs) Overlap() bool {
	if sap.first.x >= sap.second.x && sap.first.x <= sap.second.y {
		return true
	}

	if sap.first.y >= sap.second.x && sap.first.y <= sap.second.y {
		return true
	}

	if sap.second.x >= sap.first.x && sap.second.x <= sap.first.y {
		return true
	}

	if sap.second.y >= sap.first.x && sap.second.y <= sap.first.y {
		return true
	}

	return false
}

func parseInput(input string) (pairs []SectionAssignmentPairs) {
	for _, row := range strings.Split(input, "\n") {
		if row == "" {
			continue
		}
		numbers, _ := extractNumbers(row)
		pairs = append(pairs, SectionAssignmentPairs{
			first:  Pair{x: numbers[0], y: numbers[1]},
			second: Pair{x: numbers[2], y: numbers[3]},
		})
	}

	return pairs
}

func extractNumbers(s string) ([]int, error) {
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(s, -1)

	var numbers []int
	for _, match := range matches {
		num, err := strconv.Atoi(match)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, num)
	}

	return numbers, nil
}
