package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type SubmarineCommand struct {
	direction string
	units     int
}

func readInput(filename string) string {
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

func extractSubmarineCommands(input string) []SubmarineCommand {
	var commands []SubmarineCommand
	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		s := strings.Fields(scanner.Text())
		unit, _ := strconv.Atoi(s[1])
		commands = append(commands, SubmarineCommand{s[0], unit})
	}

	return commands
}

func main() {
	var input = readInput("input.txt")
	var submarineCommands = extractSubmarineCommands(input)

	var horizontal int = 0
	var depth int = 0

	for _, command := range submarineCommands {
		switch command.direction {
		case "forward":
			horizontal += command.units
		case "up":
			depth -= command.units
		case "down":
			depth += command.units
		}
	}

	fmt.Printf("Solution 1: Horizontal = %d | Depth = %d | Horizontal * Depth = %d\n", horizontal, depth, horizontal*depth)

	var aim int = 0
	horizontal = 0
	depth = 0
	for _, command := range submarineCommands {
		switch command.direction {
		case "forward":
			horizontal += command.units
			depth += (aim * command.units)
		case "up":
			aim -= command.units
		case "down":
			aim += command.units
		}
	}

	fmt.Printf("Solution 2: Horizontal = %d | Depth = %d | Aim = %d | Horizontal * Depth = %d\n", horizontal, depth, aim, horizontal*depth)
}
