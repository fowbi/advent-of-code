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

func main() {
	var input = readInput("input.txt")
	var inputAsInts = splitAndExtractInts(input)

	inputAsInts[1] = 12
	inputAsInts[2] = 2

	var i = 0
	for _, pos := range inputAsInts {
		if i%4 == 0 || i == 0 {
			if pos == 99 {
				break
			}

			inputAsInts[inputAsInts[i+3]] = execOpcode(pos, inputAsInts[inputAsInts[i+1]], inputAsInts[inputAsInts[i+2]])
		}

		i++
	}

	fmt.Println(inputAsInts[0])
}
