package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type OpCode struct {
	op            int
	positionMode  bool
	immediateMode bool
}

func readInput(filename string) string {
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	return string(content)
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

func getOpCode(opCode int) OpCode {
	return OpCode{opCode, true, false}
}

func runOps(ops []int) []int {
	diagnostics := []int{}

	pointer := 0
	for {
		opCode := getOpCode(ops[pointer])

		switch opCode.op {
		case 1:
			location := ops[pointer+3]
			argA := ops[pointer+1]
			argB := ops[pointer+2]
			ops[location] = ops[argA] + ops[argB]
			pointer += 4
		case 2:
			location := ops[pointer+3]
			argA := ops[pointer+1]
			argB := ops[pointer+2]
			ops[location] = ops[argA] * ops[argB]
			pointer += 4
		case 99:
			return diagnostics
		}
	}
}

func main() {
	input := readInput("input-test.txt")
	ops := splitAndExtractInts(input)

	ops[1] = 12
	ops[2] = 2

	fmt.Println("Solution part 1 :", runOps(ops))
	fmt.Println(ops[0])
}
