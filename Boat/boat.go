package Boat

import (
	"sort"
)

var people []int = []int{3, 2, 1, 3}
var limit = 4

func Rescue() int {
	boats := 0
	sort.Ints(people)

	heavyP := len(people) - 1
	lightP := 0

	for heavyP >= lightP {
		if people[heavyP]+people[lightP] <= limit {
			boats++
			heavyP--
			lightP++
		} else {
			boats++
			heavyP--
		}
	}
	return boats
}
