package main

import (
	"fmt"
	pcre "github.com/rubrikinc/go-pcre"
	"strconv"
	"strings"
)

func validPass(pass string) bool {
	for i := 1; i <= 5; i++ {
		if pass[i-1] > pass[i] {
			return false
		}
	}

	var isRepeated = false
	for i := 1; i <= 5; i++ {
		if pass[i-1] == pass[i] {
			isRepeated = true
		}
	}

	if isRepeated {
		re := pcre.MustCompile(`([0-9])\1`, 0)
		matches, _ := re.FindAll(pass, 0)

		for _, m := range matches {
			var first = strings.Index(pass, m.Finding)
			var last = strings.LastIndex(pass, m.Finding)

			if first == last {
				return true
			}
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
