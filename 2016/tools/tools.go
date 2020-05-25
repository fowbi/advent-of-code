package tools

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadInput(filename string) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(bytes.TrimSpace(data))

}

func Split(s string) []string {
	fs := strings.Split(s, ",")
	return fs
}

func Abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func ReadLines(filename string) []string {
	var input = ReadInput(filename)
	scanner := bufio.NewScanner(strings.NewReader(input))

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
