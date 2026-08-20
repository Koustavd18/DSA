package subsets

var Test []int = []int{1, 2} // 8 len

func Cascading(nums []int) [][]int {

	ans := [][]int{{}}

	for _, num := range nums {
		for _, set := range ans {

			cur := set
			cur = append(cur, num)
			ans = append(ans, cur)
		}
	}

	return ans
}

func BackTrack(nums []int) [][]int {

	ans := [][]int{}
	cur := []int{}
	solution(&nums, &ans, cur, 0)

	return ans
}

func solution(nums *[]int, ans *[][]int, cur []int, idx int) {

	if idx > len(*nums) {
		return
	}

	*ans = append(*ans, append([]int(nil), cur...))

	for i := idx; i < len(*nums); i++ {

		if i > idx && (*nums)[i] == (*nums)[i-1] {
			continue
		}
		cur = append(cur, (*nums)[i])
		solution(nums, ans, cur, i+1)
		cur = cur[:len(cur)-1]
	}
}
