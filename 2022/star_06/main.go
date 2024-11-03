package main

import (
	"aoc_2022/utils"
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
	datastream := parseInput(input)
	for i := 0; i < len(datastream); i++ {
		if i+4 > len(datastream) {
			continue
		}

		if !utils.HasDuplicateChars(datastream[i : i+4]) {
			return i + 4
		}
	}

	return 0
}

func part_02(input string) int {
	datastream := parseInput(input)
	for i := 0; i < len(datastream); i++ {
		if i+14 > len(datastream) {
			continue
		}

		if !utils.HasDuplicateChars(datastream[i : i+14]) {
			return i + 14
		}
	}

	return 0
}

func main() {
	fmt.Println("Solution part 1:", part_01(input))
	fmt.Println("Solution part 2:", part_02(input))
}

func parseInput(input string) (datastream string) {
	return strings.Split(input, "\n")[0]
}
