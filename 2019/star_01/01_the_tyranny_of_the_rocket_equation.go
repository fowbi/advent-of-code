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
	var totalFuel int

	var input = readInput("input.txt")
	var inputAsInts = extractInts(input)

	for _, mass := range inputAsInts {
		totalFuel += (mass / 3) - 2
	}

	fmt.Printf("Solution 1 : %d\n", totalFuel)
}
