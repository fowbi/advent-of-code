package utils

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
