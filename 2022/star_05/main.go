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

func part_01(input string) string {
	stacks, instructions := parseInput(input)

	for _, instruction := range instructions {
		src := stacks[instruction.source]
		dest := stacks[instruction.destination]
		move := src[len(src)-instruction.amount:]
		slices.Reverse(move)

		stacks[instruction.source] = src[0 : len(src)-instruction.amount]
		stacks[instruction.destination] = append(dest, move...)
	}

	result := ""
	for _, stack := range stacks {
		if len(stack) == 0 {
			continue
		}

		result += stack[len(stack)-1]
	}

	return result
}

func part_02(input string) string {
	stacks, instructions := parseInput(input)

	for _, instruction := range instructions {
		src := stacks[instruction.source]
		dest := stacks[instruction.destination]
		move := src[len(src)-instruction.amount:]

		stacks[instruction.source] = src[0 : len(src)-instruction.amount]
		stacks[instruction.destination] = append(dest, move...)
	}

	result := ""
	for _, stack := range stacks {
		if len(stack) == 0 {
			continue
		}

		result += stack[len(stack)-1]
	}

	return result
}

func main() {
	fmt.Println("Solution part 1:", part_01(input))
	fmt.Println("Solution part 2:", part_02(input))
}

type Instruction struct {
	amount      int
	source      int
	destination int
}

func parseInput(input string) (stacks [][]string, instructions []Instruction) {
	parts := strings.Split(input, "\n\n")
	structures := strings.Split(parts[0], "\n")
	structures = structures[:len(structures)-1] // remove last numbered line

	// get length of the bottom row and divide by 4 [the length of one stack]
	stackCount := (len(structures[len(structures)-1]) + 1) / 4
	for i := 0; i < stackCount; i++ {
		stacks = append(stacks, []string{})
	}

	for i := len(structures) - 1; i >= 0; i-- {
		line := structures[i]

		for j := 0; j < stackCount; j++ {
			ptr := 1 + j*4

			if ptr < len(line) && line[ptr] != ' ' {
				stacks[j] = append(stacks[j], string(line[ptr]))
			}
		}
	}

	for _, instruction := range strings.Split(parts[1], "\n") {
		if instruction == "" {
			continue
		}
		num, _ := utils.ExtractNumbers(instruction)

		instructions = append(instructions, Instruction{
			amount:      num[0],
			source:      num[1] - 1,
			destination: num[2] - 1,
		})
	}

	return stacks, instructions
}
