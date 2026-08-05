package SingleNumber

var Test []int = []int{22, 2, 1, 5, 4, 1, 7, 9, 6, 4, 5, 7, 22, 2, 6, 1}

func BruteForce(a []int) int {

	counter := make(map[int]int)

	for _, val := range a {
		if _, ok := counter[val]; ok {
			counter[val] += 1
		} else {
			counter[val] = 1
		}
	}

	for i, val := range counter {

		if val == 1 {
			return i
		}
	}

	return -1
}

// This XOR solution works only when there is a even occurences of that number
func Optimal(a []int) int {
	var ans int

	for _, val := range a {
		//NOTE: XOR Operation
		ans ^= val
	}
	return ans
}
