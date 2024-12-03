package main

import (
	"aoc_2024/utils"
	_ "embed"
	"fmt"
	"regexp"
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
	return calculateSumOfMulInstructions(input)
}

func part_02(input string) int {
	input = strings.ReplaceAll(input, "\n", "")
	re := regexp.MustCompile(`don\'t\(\)(.*?)do\(\)|don't()(.*?)$`)

	newString := re.ReplaceAllString(input, "")

	return calculateSumOfMulInstructions(newString)
}

func main() {
	fmt.Println("Solution part 1:", part_01(input))
	fmt.Println("Solution part 2:", part_02(input))
}

func calculateSumOfMulInstructions(s string) int {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	instructions := re.FindAllStringSubmatch(s, -1)

	sum := 0
	for _, instruction := range instructions {
		sum += (utils.StringToNumber(instruction[1]) * utils.StringToNumber(instruction[2]))
	}

	return sum
}
