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
	rucksacks := parseInput(input)

	sum := 0
	for _, rucksack := range rucksacks {
		partA := rucksack[:(len(rucksack) / 2)]
		partB := rucksack[(len(rucksack) / 2):]
		common := utils.CommonChars(partA, partB)[0]

		if int(common) >= 97 && int(common) <= 122 {
			sum += (int(common) - 96)
		}

		if int(common) >= 65 && int(common) <= 90 {
			sum += (int(common) - 38)
		}
	}

	return sum
}

func part_02(input string) int {
	rucksacks := parseInput(input)

	sum := 0
	for i := 0; i < len(rucksacks); i += 3 {
		firstCommon := utils.CommonChars(rucksacks[i], rucksacks[i+1])
		common := utils.CommonChars(string(firstCommon), rucksacks[i+2])[0]

		if int(common) >= 97 && int(common) <= 122 {
			sum += (int(common) - 96)
		}

		if int(common) >= 65 && int(common) <= 90 {
			sum += (int(common) - 38)
		}
	}

	return sum
}

func main() {
	fmt.Println("Solution part 1:", part_01(input))
	fmt.Println("Solution part 2:", part_02(input))
}

func parseInput(input string) (rucksacks []string) {
	for _, row := range strings.Split(input, "\n") {
		if row == "" {
			continue
		}
		rucksacks = append(rucksacks, row)
	}

	return rucksacks
}
