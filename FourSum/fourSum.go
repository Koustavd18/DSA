package foursum

var Test1, Test2, Test3, Test4 []int = []int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2} //output := 2

func BruteForce(nums1, nums2, nums3, nums4 []int) (ans int) {

	for _, val1 := range nums1 {
		for _, val2 := range nums2 {
			for _, val3 := range nums3 {
				for _, val4 := range nums4 {
					if val1+val2+val3+val4 == 0 {
						ans += 1
					}
				}
			}
		}
	}

	return
}

func Optimal(nums1, nums2, nums3, nums4 []int) (ans int) {

	sm := make(map[int]int)

	for _, val1 := range nums1 {
		for _, val2 := range nums2 {
			sum := val1 + val2
			sm[sum] += 1
		}
	}

	for _, val1 := range nums3 {
		for _, val2 := range nums4 {
			target := -(val1 + val2)
			if _, ok := sm[target]; ok {
				ans += 1
			}
		}
	}

	return
}
