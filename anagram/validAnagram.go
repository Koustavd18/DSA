package anagram

import (
	"slices"
)

var Test []string = []string{"eat", "tea", "tan", "ate", "nat", "bat"}

func BruteForce(words []string) [][]string {

	var res [][]string

	for _, val := range words {
		r := []rune(val)
		slices.Sort(r)
		a := string(r)
		var set []string

		for _, w := range words {
			c := []rune(w)
			slices.Sort(c)

			if a == string(c) {
				set = append(set, w)
			}
		}
		res = append(res, set)
	}

	return res
}

func Optimal(words []string) [][]string {

	var res [][]string
	resultMap := make(map[string][]string)

	for _, s := range words {
		sorted := []rune(s)
		slices.Sort(sorted)
		sortedS := string(sorted)

		if _, ok := resultMap[sortedS]; !ok {
			resultMap[sortedS] = make([]string, 0)
		}

		resultMap[sortedS] = append(resultMap[sortedS], s)

	}

	for _, v := range resultMap {
		res = append(res, v)
	}

	return res
}
