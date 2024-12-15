package main

import (
	"aoc_2024/utils"
	_ "embed"
	"fmt"
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
	return countStones(input, 25)
}

func part_02(input string) int {
	return countStones(input, 75)
}

func main() {
	fmt.Println("Solution part 1:", part_01(input))
	fmt.Println("Solution part 2:", part_02(input))
}

type BlinkResult struct {
	stone          int
	numberOfBlinks int
}

var cache = make(map[BlinkResult]int)

func countStones(input string, numberOfBlinks int) (totalStones int) {
	stones := strings.Split(input, " ")

	for _, stone := range stones {
		totalStones += blinking(utils.StringToNumber(stone), numberOfBlinks)
	}

	return totalStones
}

func blinking(stone int, numberOfBlinks int) int {
	// Skip processing if we already have the result
	if cachedStones, ok := cache[BlinkResult{stone, numberOfBlinks}]; ok {
		return cachedStones
	}

	// End of the line = 1 stone
	if numberOfBlinks == 0 {
		return 1
	}

	// 0 turns into 1
	if stone == 0 {
		return blinking(1, numberOfBlinks-1)
	}

	// Stones with even number of digits are split in half, no extra leading zeroes!
	if splitNumber := strconv.Itoa(stone); len(splitNumber)%2 == 0 { // even
		leftStone, _ := strconv.Atoi(splitNumber[:len(splitNumber)/2])
		rightStone, _ := strconv.Atoi(splitNumber[len(splitNumber)/2:])
		stones := blinking(leftStone, numberOfBlinks-1) + blinking(rightStone, numberOfBlinks-1)
		cache[BlinkResult{stone, numberOfBlinks}] = stones

		return stones
	}

	// the other stones are multiplied by 2024
	stones := blinking(stone*2024, numberOfBlinks-1)
	cache[BlinkResult{stone, numberOfBlinks}] = stones

	return stones
}
