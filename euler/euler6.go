package euler

import "GoProjectEuler/goutils"

// Euler6 solution
func Euler6(p int) int {
	sumSquare := goutils.SumSquare(p)
	squareSum := goutils.SquareSum(p)

	squareDiff := squareSum - sumSquare
	return squareDiff
}
