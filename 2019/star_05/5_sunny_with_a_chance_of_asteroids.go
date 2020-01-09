package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Mode string

const (
	PositionMode  Mode = "position"
	ImmediateMode Mode = "immediate"
)

type OpCode struct {
	op       int
	modeArgA Mode
	modeArgB Mode
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

func PadLeft(str, pad string, lenght int) string {
	for {
		str = pad + str
		if len(str) >= lenght {
			return str[0:lenght]
		}
	}
}

func determineMode(opCode string) Mode {
	if opCode == "1" {
		return ImmediateMode
	}

	return PositionMode
}

func getOpCode(opCode int) OpCode {
	o := PadLeft(strconv.Itoa(opCode), "0", 5)

	return OpCode{
		opCode % 100,
		determineMode(string(o[2])),
		determineMode(string(o[1])),
	}
}

func loadArg(ops []int, pointer int, mode Mode) int {
	if mode == PositionMode {
		return ops[pointer]
	} else {
		return pointer
	}
}

func runOps(ops []int, inputInstruction int) []int {
	diagnostics := []int{}

	pointer := 0
	for {
		opCode := getOpCode(ops[pointer])

		switch opCode.op {
		case 1:
			location := ops[pointer+3]
			argA := loadArg(ops, pointer+1, opCode.modeArgA)
			argB := loadArg(ops, pointer+2, opCode.modeArgB)

			ops[location] = ops[argA] + ops[argB]
			pointer += 4
		case 2:
			location := ops[pointer+3]
			argA := loadArg(ops, pointer+1, opCode.modeArgA)
			argB := loadArg(ops, pointer+2, opCode.modeArgB)

			ops[location] = ops[argA] * ops[argB]
			pointer += 4
		case 3:
			location := ops[pointer+1]
			argA := inputInstruction

			ops[location] = argA
			pointer += 2
		case 4:
			location := ops[pointer+1]
			diagnostics = append(diagnostics, ops[location])

			pointer += 2
		case 5:
			argA := loadArg(ops, pointer+1, opCode.modeArgA)
			argB := loadArg(ops, pointer+2, opCode.modeArgB)

			if ops[argA] != 0 {
				pointer = ops[argB]
			} else {
				pointer += 3
			}
		case 6:
			argA := loadArg(ops, pointer+1, opCode.modeArgA)
			argB := loadArg(ops, pointer+2, opCode.modeArgB)

			if ops[argA] == 0 {
				pointer = ops[argB]
			} else {
				pointer += 3
			}
		case 7:
			location := ops[pointer+3]
			argA := loadArg(ops, pointer+1, opCode.modeArgA)
			argB := loadArg(ops, pointer+2, opCode.modeArgB)

			if ops[argA] < ops[argB] {
				ops[location] = 1
			} else {
				ops[location] = 0
			}
			pointer += 4
		case 8:
			location := ops[pointer+3]
			argA := loadArg(ops, pointer+1, opCode.modeArgA)
			argB := loadArg(ops, pointer+2, opCode.modeArgB)

			if ops[argA] == ops[argB] {
				ops[location] = 1
			} else {
				ops[location] = 0
			}
			pointer += 4
		case 99:
			return diagnostics
		}
	}
}

func main() {
	input := readInput("input.txt")
	input = strings.TrimRight(input, "\n")
	ops := splitAndExtractInts(input)

	fmt.Println("Solution part 1 :", runOps(append(ops[:0:0], ops...), 1))
	fmt.Println("Solution part 2 :", runOps(append(ops[:0:0], ops...), 5))
}
