package calc

import "math"

func Sum(a, b int) int {
	return a + b
}
func abs(x int) int {
	if x < 0 {
		return int(math.Abs(float64(x)))
	}
	return x
}
