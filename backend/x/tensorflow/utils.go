package tfrt

func toIntSlice(input []int64) []int {
	output := make([]int, len(input))
	for i := 0; i < len(input); i++ {
		output[i] = int(input[i])
	}
	return output
}
