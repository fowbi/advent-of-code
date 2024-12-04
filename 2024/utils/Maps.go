package utils

func SliceToMap(input []string) map[int]string {
	output := make(map[int]string)
	for i := 0; i < len(input); i++ {
		output[i] = input[i]
	}

	return output
}
