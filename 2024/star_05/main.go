package main

import (
	"aoc_2024/utils"
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
	inputParts := strings.Split(input, "\n\n")
	rules := extractList(inputParts[0], "|")
	updates := extractList(inputParts[1], ",")

	count := 0
	for _, update := range updates {
		if !rightOrder(update, rules) {
			continue
		}

		count += update[(len(update)-1)/2]
	}

	return count
}

func part_02(input string) int {
	inputParts := strings.Split(input, "\n\n")
	rules := extractList(inputParts[0], "|")
	updates := extractList(inputParts[1], ",")

	count := 0
	for _, update := range updates {
		if rightOrder(update, rules) {
			continue
		}

		fixedUpdate := fixOrder(update, rules)

		count += fixedUpdate[(len(fixedUpdate)-1)/2]
	}

	return count
}

func main() {
	fmt.Println("Solution part 1:", part_01(input))
	fmt.Println("Solution part 2:", part_02(input))
}

func extractList(input string, separator string) [][]int {
	list := [][]int{}
	for _, row := range strings.Split(input, "\n") {
		if row == "" {
			continue
		}
		splitRow := strings.Split(row, separator)
		newRow := []int{}
		for _, item := range splitRow {
			newRow = append(newRow, utils.CastStringToInt(item))
		}
		list = append(list, newRow)
	}
	return list
}

func rightOrder(update []int, rules [][]int) bool {
	for _, rule := range rules {
		first := slices.Index(update, rule[0])
		if first == -1 {
			continue
		}
		second := slices.Index(update, rule[1])
		if second == -1 {
			continue
		}

		if first > second {
			return false
		}
	}
	return true
}

func fixOrder(update []int, rules [][]int) []int {
	fixedUpdate := append([]int(nil), update...)
	for _, rule := range rules {
		first := slices.Index(fixedUpdate, rule[0])
		if first == -1 {
			continue
		}
		second := slices.Index(fixedUpdate, rule[1])
		if second == -1 {
			continue
		}

		if first > second {
			fixedUpdate[first], fixedUpdate[second] = fixedUpdate[second], fixedUpdate[first]
		}
	}

	if !rightOrder(fixedUpdate, rules) {
		return fixOrder(fixedUpdate, rules)
	}

	return fixedUpdate
}
