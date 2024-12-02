package utils

func Remove(slice []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, slice[:index]...)

	return append(ret, slice[index+1:]...)
}
