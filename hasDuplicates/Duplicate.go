package hasduplicates

import (
	"slices"
)

var Test []int = []int{2, 1, 3, 1}

func BruteForce(nums []int) bool {

	for i, val := range nums {

		for j := i + 1; j < len(nums); j++ {
			if val == nums[j] {
				return true
			}
		}
	}

	return false
}

// Requires Sorting the input slice
func TwoPointer(nums []int) bool {

	sortedNums := slices.Sorted(slices.Values(nums))

	for i := 1; i < len(sortedNums); i++ {
		if sortedNums[i] == sortedNums[i-1] {
			return true
		}
	}

	return false
}

func Optimal(nums []int) bool {
	seen := make(map[int]int)

	for i, val := range nums {
		if _, ok := seen[val]; ok {
			return true
		}

		seen[val] = i
	}

	return false
}
