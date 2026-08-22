package combinationsum

var Candidate = []int{2, 3, 5, 6, 7}
var Target = 7

func CombiSum(candidates []int, target int) [][]int {

	ans := [][]int{}
	cur := []int{}
	Backtrack(&ans, candidates, target, cur, 0, 0)
	return ans
}

func Backtrack(ans *[][]int, candidates []int, target int, cur []int, sum int, idx int) {
	if sum == target {
		combi := append([]int(nil), cur...)
		*ans = append(*ans, combi)
	} else if sum < target {
		for i := idx; i < len(candidates); i++ {
			cur = append(cur, candidates[i])
			Backtrack(ans, candidates, target, cur, sum+candidates[i], i)
			cur = cur[:len(cur)-1]
		}
	}
}
