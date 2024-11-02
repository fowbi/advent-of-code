package main

import (
	"aoc_2022/utils"
	_ "embed"
	"fmt"
	"slices"
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
	return utils.FindLargestNumber(parseInput(input))
}

func part_02(input string) int {
	elves := parseInput(input)
	slices.Sort(elves)
	slices.Reverse(elves)

	return elves[0] + elves[1] + elves[2]
}

func main() {
	fmt.Println("Solution part 1:", part_01(input))
	fmt.Println("Solution part 2:", part_02(input))
}

func parseInput(input string) (elves []int) {
	elf := 0

	for _, row := range strings.Split(input, "\n") {
		if row == "" {
			elves = append(elves, elf)
			elf = 0
			continue
		}
		elf += utils.CastStringToInt(row)
	}

	return elves
}
