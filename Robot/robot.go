package Robot

var Moves []string = []string{"U", "D", "L", "R"}

func Solution(moves []string) (int, int, bool) {
	x := 0
	y := 0
	for _, move := range moves {
		switch move {
		case "U":
			y++
		case "D":
			y--
		case "L":
			x--
		case "R":
			x++

		}
	}

	return x, y, x == 0 && y == 0
}
