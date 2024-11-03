package utils

import (
	"regexp"
	"strconv"
)

func CommonChars(str1, str2 string) []rune {
	charMap := make(map[rune]bool)
	var common []rune

	for _, char := range str1 {
		charMap[char] = true
	}

	for _, char := range str2 {
		if charMap[char] {
			common = append(common, char)
			delete(charMap, char)
		}
	}

	return common
}

func ExtractNumbers(s string) ([]int, error) {
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(s, -1)

	var numbers []int
	for _, match := range matches {
		num, err := strconv.Atoi(match)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, num)
	}

	return numbers, nil
}

func DuplicateChars(s string) map[rune]int {
	charMap := make(map[rune]int)

	for _, char := range s {
		charMap[char]++
	}

	return charMap
}

func HasDuplicateChars(s string) bool {
	for _, count := range DuplicateChars(s) {
		if count > 1 {
			return true
		}
	}

	return false
}
