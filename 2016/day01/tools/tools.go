package tools

import (
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

func Split(s string) []string {
	fs := strings.Split(s, ",")
	return fs
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
