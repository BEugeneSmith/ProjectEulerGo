package euler

import "GoProjectEuler/goutils"

// Euler2 solution
func Euler2(l int) int {
	FibSeq := goutils.FibonacciLimit(l)
	var SumOfEvens int

	for i := 1; i < len(FibSeq); i++ {
		FibElem := FibSeq[i]
		if FibElem < l && goutils.IsDiv(FibElem, 2) {
			SumOfEvens += FibElem
		}
	}
	return SumOfEvens
}
