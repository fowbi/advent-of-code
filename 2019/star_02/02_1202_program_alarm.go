package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func readInput(filename string) string {
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

func extractInts(input string) []int {
	var list = []int{}

	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		v, _ := strconv.Atoi(scanner.Text())
		list = append(list, v)
	}

	return list
}

func splitAndExtractInts(input string) []int {
	var list = []int{}
	s := strings.Split(input, ",")

	for _, i := range s {
		v, _ := strconv.Atoi(i)
		list = append(list, v)
	}

	return list
}

func execOpcode(opcode int, a int, b int) int {
	switch opcode {
	case 1:
		return a + b
	case 2:
		return a * b
	}

	return 1
}

func calcOutput(memory []int, noun int, verb int) int {
	memory[1] = noun
	memory[2] = verb

	for i := 1; i <= len(memory); i++ {
		if i%4 == 0 || i == 0 {
			if memory[i] == 99 {
				break
			}

			memory[memory[i+3]] = execOpcode(memory[i], memory[memory[i+1]], memory[memory[i+2]])
		}
	}

	return memory[0]
}

func findNounVerb(inputAsInts []int, match int) int {
	for noun := 1; noun <= 99; noun++ {
		for verb := 1; verb <= 99; verb++ {
			memory := make([]int, len(inputAsInts))
			copy(memory, inputAsInts)
			var output = calcOutput(memory, noun, verb)

			if output == match {
				return (100 * noun) + verb
			}
		}
	}

	return 0
}

func main() {
	var input = readInput("input.txt")
	var inputAsInts = splitAndExtractInts(input)

	memory := make([]int, len(inputAsInts))
	copy(memory, inputAsInts)
	var output = calcOutput(memory, 12, 2)

	fmt.Println("solution part 1 :", output)

	fmt.Println("solution part 2 :", findNounVerb(inputAsInts, 19690720))
}
