package utils

import "strconv"

func FindLargestNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	largest := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > largest {
			largest = nums[i]
		}
	}

	return largest
}

func Diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func UniqueValues(list []int) map[int]int {
	dict := make(map[int]int)
	for _, num := range list {
		dict[num] = dict[num] + 1
	}
	return dict
}

func StringToNumber(s string) int {
	num, _ := strconv.Atoi(s)

	return num
}
