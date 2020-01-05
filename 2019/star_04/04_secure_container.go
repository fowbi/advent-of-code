package main

import (
	"fmt"
	"strconv"
)

func validPass(pass string) bool {
	for i := 1; i <= 5; i++ {
		if pass[i-1] > pass[i] {
			return false
		}
	}

	for i := 1; i <= 5; i++ {
		if pass[i-1] == pass[i] {
			return true
		}
	}

	return false
}

func main() {
	var minPass = 231832
	var maxPass = 767346
	var passes = 0

	for pass := minPass; pass <= maxPass; pass++ {
		if validPass(strconv.Itoa(pass)) {
			passes++
		}
	}

	fmt.Println("passes", passes)
}
