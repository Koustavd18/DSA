package MoveZeros

var test []int = []int{0, 0, 1, 3, 0, 12}

func BruteForce() []int {

	result := []int{}

	for _, val := range test {
		if val != 0 {
			result = append(result, val)
		}
	}

	diff := len(test) - len(result)

	for i := 0; i < diff; i++ {
		result = append(result, 0)
	}

	return result
}

func Optimal() []int {
	j := 0

	for _, val := range test {
		if val != 0 {
			test[j] = val
			j++
		}
	}

	for j < len(test) {
		test[j] = 0
		j++
	}

	return test
}
