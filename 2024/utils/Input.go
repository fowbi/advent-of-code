package utils

import "strings"

func InputToSlice(input string) []string {
	list := []string{}
	for _, row := range strings.Split(input, "\n") {
		if row == "" {
			continue
		}
		list = append(list, row)
	}

	return list
}

func InputToMap(input string) map[int]string {
	return SliceToMap(InputToSlice(input))
}

func InputToMatrix(input string) [][]string {
	matrix := [][]string{}
	for _, row := range strings.Split(input, "\n") {
		if row == "" {
			continue
		}
		matrix = append(matrix, strings.Split(row, ""))
	}

	return matrix
}
