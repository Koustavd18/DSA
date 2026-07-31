package ProductExpectSelf

func Execute(nums []int) []int {
	n := len(nums)
	res := make([]int, n)

	leftProd := 1
	for i := 0; i < n; i++ {
		res[i] = leftProd
		leftProd *= nums[i]
	}

	rightProd := 1
	for i := n - 1; i >= 0; i-- {
		res[i] *= rightProd
		rightProd *= nums[i]
	}

	return res
}
