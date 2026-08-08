package majority

var Test []int = []int{1, 2, 1, 1, 3}

func BruteForce(nums []int) int {

	for i, val := range nums {

		currEl := val
		currCount := 0

		for _, j := range nums[i+1:] {
			if j == currEl {
				currCount++
			}
		}
		if currCount >= len(nums)/2 {
			return currEl
		}
	}

	return -1
}

func UseMaps(nums []int) int {

	seen := make(map[int]int)

	for _, val := range nums {
		seen[val] += 1

		if seen[val] > len(nums)/2 {
			return val
		}
	}
	return -1
}

func Optimal(nums []int) int {

	var candidate, count int = nums[0], 1

	for _, val := range nums {
		if count == 0 {
			candidate = val
		}
		if val != candidate {
			count--
		} else {
			count++
		}
	}

	return candidate
}
