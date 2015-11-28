package euler

import "GoProjectEuler/goutils"

// Euler3 solution
func Euler3(l int) []int {
	allFactors := goutils.GetFactors(l)
	var primeFactors []int

	for i := 0; i < len(allFactors); i++ {
		if goutils.IsPrime(allFactors[i]) {
			primeFactors = append(primeFactors, allFactors[i])
		}
	}

	return primeFactors
}
