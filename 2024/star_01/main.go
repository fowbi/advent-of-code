package main

import (
	"aoc_2024/utils"
	_ "embed"
	"fmt"
	"sort"
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
	leftList := []int{}
	rightList := []int{}

	leftList, rightList = extractlists(input)

	sort.Ints(leftList)
	sort.Ints(rightList)

	distance := 0
	for i := 0; i < len(leftList); i++ {
		distance += utils.Diff(leftList[i], rightList[i])
	}

	return distance
}

func part_02(input string) int {
	similarityScore := 0
	leftList := []int{}
	rightList := []int{}

	leftList, rightList = extractlists(input)
	mappedRightList := utils.UniqueValues(rightList)

	for i := 0; i < len(leftList); i++ {
		similarityScore += leftList[i] * mappedRightList[leftList[i]]
	}

	return similarityScore
}

func main() {
	fmt.Println("Solution part 1:", part_01(input))
	fmt.Println("Solution part 2:", part_02(input))
}

func extractlists(input string) ([]int, []int) {
	leftList := []int{}
	rightList := []int{}

	for _, row := range strings.Split(input, "\n") {
		if row == "" {
			continue
		}

		splitRow := strings.Fields(row)
		leftList = append(leftList, utils.CastStringToInt(splitRow[0]))
		rightList = append(rightList, utils.CastStringToInt(splitRow[1]))
	}

	return leftList, rightList
}
