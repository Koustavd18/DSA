package TwoSum

var Test []int = []int{2, 55, 11, 4, 7, 88, 90, 51, 16, 15, 29, 31, 33} //2,9

func BruteForce(nums []int, target int) []int {

	i := 0

	for i < len(nums) {
		if nums[i] >= target {
			i++
			continue
		}

		req := target - nums[i]

		j := i + 1

		for j < len(nums) {

			if nums[j] > req {
				j++
				continue
			}

			if nums[j] == req {
				return []int{i, j}
			}

			j++
		}

		i++
	}

	return []int{}
}

func Optimal(nums []int, target int) []int {

	seen := make(map[int]int)

	for i, val := range nums {
		req := target - val

		if j, ok := seen[req]; ok {
			return []int{j, i}
		}

		seen[val] = i
	}

	return []int{}
}
