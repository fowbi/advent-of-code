package main

import (
	"aoc_2024/utils"
	"bytes"
	_ "embed"
	"fmt"
	"regexp"
	"strings"
)

//go:embed input.txt
var input string

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("No input provided")
	}
}

func part_01(input string) int {
	horizontal := utils.InputToSlice(input)
	diagonalOne, diagonalTwo := diagonals(horizontal)

	totalMatches := 0
	xmasRe := regexp.MustCompile("XMAS")
	samxRe := regexp.MustCompile("SAMX")
	directions := []map[int]string{
		utils.SliceToMap(horizontal),
		vertical(horizontal),
		diagonalOne, diagonalTwo,
	}
	for _, direction := range directions {
		for _, row := range direction {
			if len(row) <= 3 {
				continue
			}

			matches1 := xmasRe.FindAllStringIndex(row, -1)
			matches2 := samxRe.FindAllStringIndex(row, -1)
			totalMatches += len(matches1) + len(matches2)
		}
	}

	return totalMatches
}

func part_02(input string) int {
	wordSearch := utils.InputToSlice(input)

	rows := len(wordSearch)
	cols := len(wordSearch[0])
	count := 0

	for y := range rows {
		for x := range cols {
			if wordSearch[y][x] == 'A' {
				count += checkMas(wordSearch, rows, cols, y-1, y+1, x-1, x+1)
			}
		}
	}

	return count
}

func main() {
	fmt.Println("Solution part 1:", part_01(input))
	fmt.Println("Solution part 2:", part_02(input))
}

func vertical(matrix []string) map[int]string {
	vertical := make(map[int]string)

	for column, _ := range matrix[0] {
		for row, _ := range matrix {
			vertical[column] += string(matrix[row][column])
		}
	}

	return vertical
}

func diagonals(matrix []string) (map[int]string, map[int]string) {
	diagonalOne := make(map[int]string)
	diagonalTwo := make(map[int]string)
	rows := len(matrix)
	cols := len(matrix[0])

	for i := 0; i < rows+cols-1; i++ {
		for j := max(0, i-rows+1); j <= min(i, cols-1); j++ {
			diagonalOne[i] += string(matrix[i-j][j])
			diagonalTwo[i] += string(matrix[rows-1-i+j][j])
		}
	}

	return diagonalOne, diagonalTwo
}

func checkMas(input []string, rows int, cols int, y1 int, y2 int, x1 int, x2 int) int {
	sm := []byte{'S', 'M'}
	ms := []byte{'M', 'S'}

	if (0 <= y1 && y1 < rows) && (0 <= y2 && y2 < rows) && (0 <= x1 && x1 < cols) && (0 <= x2 && x2 < cols) {
		mas1 := []byte{input[y1][x1], input[y2][x2]}
		mas2 := []byte{input[y1][x2], input[y2][x1]}

		if (bytes.Compare(mas1, sm) == 0 || bytes.Compare(mas1, ms) == 0) && (bytes.Compare(mas2, sm) == 0 || bytes.Compare(mas2, ms) == 0) {
			return 1
		}
	}

	return 0
}
