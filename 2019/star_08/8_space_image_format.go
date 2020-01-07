package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func readInput(filename string) string {
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

type Layer struct {
	lines []string
}

func extractLayers(input string, width int, height int) []Layer {
	scanner := bufio.NewScanner(strings.NewReader(input))

	layers := []Layer{}
	for scanner.Scan() {
		line := scanner.Text()

		var lines = []string{}

		for i := 0; i < len(line); i += width {
			lines = append(lines, line[i:i+width])

			if len(lines)%height == 0 && i != 0 {
				layers = append(layers, Layer{lines})
				lines = []string{}
			}
		}
	}

	return layers
}

func numOf(lines []string, digit string) int {
	var num = 0
	for _, line := range lines {
		num += strings.Count(line, digit)
	}

	return num
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func main() {
	var width = 25
	var height = 6
	var input = readInput("input.txt")
	var layers = extractLayers(input, width, height)

	var leastZeroes = width + 1
	var check = 0

	for _, layer := range layers {
		numOfZeroes := numOf(layer.lines, "0")

		var newLeastZeroes = Min(numOfZeroes, leastZeroes)

		if newLeastZeroes < leastZeroes {
			leastZeroes = newLeastZeroes

			check = numOf(layer.lines, "1") * numOf(layer.lines, "2")
		}
	}

	fmt.Println("solution part 1 :", check)

	picture := make([][]uint8, height)
	for i := range picture {
		picture[i] = make([]uint8, width)
	}

	// make everything transparent
	for x, row := range picture {
		for y := range row {
			picture[x][y] = 2
		}
	}

	for x, row := range picture {
		for y := range row {
			for _, layer := range layers {
				var line = layer.lines[x]
				l := strings.Split(line, "")

				switch l[y] {
				case "0":
					if picture[x][y] == 2 {
						picture[x][y] = 0
					}
					break
				case "1":
					if picture[x][y] == 2 {
						picture[x][y] = 1
					}
					break
				}
			}
		}
	}

	fmt.Println("solution part 2 :")
	for _, row := range picture {
		for _, pixel := range row {
			if pixel == 1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
}
