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

func main() {
	var increases int = 0

	var input = readInput("input.txt")
	var inputAsInts = extractInts(input)

	var previous int = inputAsInts[0]
	for _, number := range inputAsInts {
		if previous < number {
			increases++
		}
		previous = number
	}

	fmt.Printf("Solution 1: %d\n", increases)

	increases = -1
	previous = 0

	var sumOfThree int
	for i := 0; i < len(inputAsInts)-2; i++ {
		sumOfThree = inputAsInts[i] + inputAsInts[i+1] + inputAsInts[i+2]
		if sumOfThree > previous {
			increases++
		}
		previous = sumOfThree
	}
	fmt.Printf("Solution 2: %d\n", increases)
}
