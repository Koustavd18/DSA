package validMountain

var sequence []int = []int{0, 2, 3, 4, 5, 2, 1}

func Execute() bool {
	if len(sequence) < 3 {
		return false
	}
	var i int = 1

	for i < len(sequence) && sequence[i] > sequence[i-1] {
		i++
	}
	if i == 1 || i == len(sequence) {
		return false
	}

	for i < len(sequence) && sequence[i] < sequence[i-1] {
		i++
	}

	return i == len(sequence)
}
