package euler

import "GoProjectEuler/goutils"

// Euler1 solution
func Euler1(l int) int {
	sum := 0
	for i := 1; i < l; i++ {
		if goutils.IsDiv(i, 3) || goutils.IsDiv(i, 5) {
			sum += i
		}
	}
	return sum
}
