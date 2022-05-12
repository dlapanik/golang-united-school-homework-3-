package homework

func reverse(input []int64) (result []int64) {

	result = make([]int64, len(input))
	j := 0

	for i := len(result) - 1; i >= 0; i-- {
		result[i] = input[j]
		j++
	}

	return
}
