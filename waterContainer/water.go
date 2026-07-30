package waterContainer

import "math"

var max_area float64

var container []float64 = []float64{5, 9, 2, 4, 3, 7}

func BruteForce() float64 {

	max_area = 0

	n := len(container)

	for i, val := range container {
		for p2 := i + 1; p2 < n; p2++ {
			length := math.Min(val, container[p2])
			width := p2 - i
			area := length * float64(width)
			max_area = math.Max(max_area, area)
		}
	}

	return max_area

}

func Optimal() float64 {

	max_area = 0

	l := 0
	r := len(container) - 1

	for l < r {
		length := math.Min(container[l], container[r])
		width := r - l
		area := length * float64(width)
		max_area = math.Max(max_area, area)
		if container[l] < container[r] {
			l++
		} else {
			r--
		}
	}

	return max_area
}
