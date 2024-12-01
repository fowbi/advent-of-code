package utils

import "strconv"

func CastStringToInt(s string) int {
	x, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return x
}
