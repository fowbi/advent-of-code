package tools

import (
	"bytes"
	"io/ioutil"
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
